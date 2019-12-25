package hmgdb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_ERROR = errors.New("99100:mongo->%s")
)

func MongoPanic(err error) {
	if err != nil {
		log.Panic(errors.New(fmt.Sprintf(MONGO_ERROR.Error(), err.Error())))
	}
}

func InsertOne(ctx context.Context, c *mongo.Collection, document interface{}, opts ...*options.InsertOneOptions) (objectId string) {
	insertOneResult, err := c.InsertOne(ctx, document, opts...)
	MongoPanic(err)
	return insertOneResult.InsertedID.(primitive.ObjectID).Hex()
}

func UpdateOne(ctx context.Context, c *mongo.Collection, filter interface{}, update interface{}, opts ...*options.UpdateOptions) *mongo.UpdateResult {
	updateResult, err := c.UpdateOne(ctx, filter, update, opts...)
	MongoPanic(err)
	return updateResult
}

func FindOneAndUpdate(ctx context.Context, collection mongo.Collection, filter interface{}, update interface{}, singleResult interface{}) {
	sr := collection.FindOneAndUpdate(ctx, filter, update)
	err := sr.Decode(singleResult)
	if err != nil {
		MongoPanic(err)
	}
}

// 检查集合是否存在，不存在自动创建，因不存在的集合无法使用事务
func CollectionCreate(collection *mongo.Collection) {
	ctx := context.TODO()
	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetLimit(1))
	if err != nil {
		MongoPanic(err)
	}
	if !cursor.Next(ctx) {
		InsertOne(nil, collection, bson.M{"createdAt": time.Now()})
	}
}

func UpdateOneOrInsertOne(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}, insert interface{}) {
	ur, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		MongoPanic(err)
	}
	if ur.ModifiedCount >= 1 {
		return
	}
	_, err = collection.InsertOne(ctx, insert)
	if err != nil {
		MongoPanic(err)
	}
}

// 使用事务
func UseSession(ctx context.Context, client *mongo.Client, fn func(context mongo.SessionContext) error) {
	err := client.UseSession(ctx, func(sessionContext mongo.SessionContext) (err error) {
		defer func() {
			if err != nil {
				sessionContext.AbortTransaction(sessionContext)
				MongoPanic(err)
			}
			err := recover().(error)
			MongoPanic(err)
		}()
		err = sessionContext.StartTransaction()
		MongoPanic(err)
		err = fn(sessionContext)
		if err != nil {
			return err
		}
		err = sessionContext.CommitTransaction(sessionContext)
		MongoPanic(err)
		return err
	})
	MongoPanic(err)
}

func Exists(ctx context.Context, c *mongo.Collection, filter interface{}, opts ...*options.FindOneOptions) (b bool) {
	var rsr map[string]interface{}
	err := c.FindOne(ctx, filter, opts...).Decode(&rsr)
	return err == nil
}

func Find(ctx context.Context, c *mongo.Collection, filter interface{}, results interface{}, opts ...*options.FindOptions) {
	cursor, err := c.Find(ctx, filter, opts...)
	MongoPanic(err)
	err = cursor.All(ctx, results)
	MongoPanic(err)
}

func FindOne(ctx context.Context, c *mongo.Collection, filter interface{}, singleResult interface{}, opts ...*options.FindOneOptions) {
	err := c.FindOne(ctx, filter, opts...).Decode(singleResult)
	MongoPanic(err)
}
