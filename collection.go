package hmgdb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _cmap = make(map[string]bool)

// 创建集合
func CreateCollection(db *mongo.Database, collectionName string) error {
	if _cmap[collectionName] {
		return nil
	}
	cursor, err := db.Collection(collectionName).Find(nil, bson.M{}, options.Find().SetLimit(1))
	if err != nil {
		return err
	}
	if !cursor.Next(nil) {
		_, err := db.Collection(collectionName).InsertOne(nil, bson.M{"__temp": "createCollection"})
		if err != nil {
			return err
		}
		_, err = db.Collection(collectionName).DeleteOne(nil, bson.M{"__temp": "createCollection"})
		if err != nil {
			return err
		}
	}
	_cmap[collectionName] = true
	return nil
}
