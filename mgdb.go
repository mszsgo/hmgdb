package hmgdb

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 此文件方法不再使用

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

func FindOneAndUpdate(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}, singleResult interface{}) {
	sr := collection.FindOneAndUpdate(ctx, filter, update)
	err := sr.Decode(singleResult)
	if err != nil {
		MongoPanic(err)
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
func UseSession(ctx context.Context, client *mongo.Client, fn func(context mongo.SessionContext) error) error {
	return client.UseSession(ctx, func(sessionContext mongo.SessionContext) (err error) {
		e := sessionContext.StartTransaction()
		if e != nil {
			MongoPanic(e)
		}
		err = fn(sessionContext)
		if err == nil {
			e = sessionContext.CommitTransaction(sessionContext)
			if e != nil {
				MongoPanic(e)
			}
			return
		}
		e = sessionContext.AbortTransaction(sessionContext)
		if e != nil {
			MongoPanic(e)
		}
		return err
	})
}

func CountDocuments(ctx context.Context, c *mongo.Collection, filter interface{}, opts ...*options.CountOptions) int64 {
	count, err := c.CountDocuments(ctx, filter, opts...)
	if err != nil {
		MongoPanic(err)
	}
	return count
}

func Exists(ctx context.Context, c *mongo.Collection, filter interface{}) (b bool) {
	cursor, err := c.Find(ctx, filter, options.Find().SetLimit(1))
	if err != nil {
		MongoPanic(err)
	}
	return cursor.Next(ctx)
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
