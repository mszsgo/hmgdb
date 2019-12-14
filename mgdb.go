// Mongodb 通用操作代码库
package hmgdb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_ERROR = errors.New("99100:mongo->%s")
)

func mongoPanic(err error) {
	if err != nil {
		panic(errors.New(fmt.Sprintf(MONGO_ERROR.Error(), err.Error())))
	}
}

func UseSession(ctx context.Context, client *mongo.Client, fn func(sessionContext mongo.SessionContext) error) {
	err := client.UseSession(ctx, fn)
	mongoPanic(err)
}

func InsertOne(ctx context.Context, c *mongo.Collection, document interface{}, opts ...*options.InsertOneOptions) (objectId string) {
	insertOneResult, err := c.InsertOne(ctx, document, opts...)
	mongoPanic(err)
	return insertOneResult.InsertedID.(primitive.ObjectID).Hex()
}

func UpdateOne(ctx context.Context, c *mongo.Collection, filter interface{}, update interface{}, opts ...*options.UpdateOptions) *mongo.UpdateResult {
	updateResult, err := c.UpdateOne(ctx, filter, update, opts...)
	mongoPanic(err)
	return updateResult
}

func Exists(ctx context.Context, c *mongo.Collection, filter interface{}, opts ...*options.FindOneOptions) (b bool) {
	var rsr map[string]interface{}
	err := c.FindOne(ctx, filter, opts...).Decode(&rsr)
	return err == nil
}

func FindOne(ctx context.Context, c *mongo.Collection, filter interface{}, returnSingleResult interface{}, opts ...*options.FindOneOptions) {
	err := c.FindOne(ctx, filter, opts...).Decode(returnSingleResult)
	mongoPanic(err)
}

func Find(ctx context.Context, c *mongo.Collection, filter interface{}, returnSingleResult interface{}, opts ...*options.FindOptions) {
	/*err := c.Find(ctx, filter, opts...).Decode(returnSingleResult)
	if err != nil {
		Panic(MONGO_ERROR, err.Error())
	}*/
}
