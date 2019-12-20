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

var (
	MONGO_ERROR = errors.New("99100:mongo->%s")
)

func MongoPanic(err error) {
	if err != nil {
		log.Panic(errors.New(fmt.Sprintf(MONGO_ERROR.Error(), err.Error())))
	}
}

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

func Exists(ctx context.Context, c *mongo.Collection, filter interface{}, opts ...*options.FindOneOptions) (b bool) {
	var rsr map[string]interface{}
	err := c.FindOne(ctx, filter, opts...).Decode(&rsr)
	return err == nil
}

func FindOne(ctx context.Context, c *mongo.Collection, filter interface{}, returnSingleResult interface{}, opts ...*options.FindOneOptions) {
	err := c.FindOne(ctx, filter, opts...).Decode(returnSingleResult)
	MongoPanic(err)
}

func Find(ctx context.Context, c *mongo.Collection, filter interface{}, cursorFunc func(cursor *mongo.Cursor), opts ...*options.FindOptions) {
	cursor, err := c.Find(ctx, filter, opts...)
	for {
		if !cursor.Next(ctx) {
			break
		}
		cursorFunc(cursor)
	}
	MongoPanic(err)
	return
}

func Cursor(cursor *mongo.Cursor, result interface{}) {
	err := cursor.Decode(result)
	MongoPanic(err)
}

func FindOneAndUpdate(ctx context.Context, collection mongo.Collection, filter interface{}, update interface{}, singleResult interface{}) {
	sr := collection.FindOneAndUpdate(ctx, filter, update)
	err := sr.Decode(singleResult)
	if err != nil {
		MongoPanic(err)
	}
}
