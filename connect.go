package hmgdb

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DEFAULT = "default"

	// Mongodb Connect, map[name]*mongo.Database
	mgdb sync.Map
)

func Db() *mongo.Database {
	return GetDatabase(DEFAULT)
}

func GetDatabase(name string) *mongo.Database {
	db, ok := mgdb.Load(name)
	if !ok {
		panic(fmt.Sprintf("No Mongodb connection was obtained '%s'", name))
	}
	return db.(*mongo.Database)
}

func SetDatabase(name string, v *mongo.Database) {
	if name == "" {
		name = DEFAULT
	}
	mgdb.Store(name, v)
}

func SetConnectString(name string, connectionString string) {
	SetDatabase(name, connect(connectionString))
}

// 注意处理连接panic
func connect(connectionString string) *mongo.Database {
	if connectionString == "" {
		log.Panic("Mongodb 连接字符串不能为空")
	}
	dbName := (strings.Split((strings.Split(connectionString, "/"))[3], "?"))[0]
	if dbName == "" {
		panic(fmt.Sprintf("Errror Mongodb connectionString %s", connectionString))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		panic(fmt.Sprintf("Errror Connect mongodb exception %s", err))
	}
	database := client.Database(dbName)
	names, err := client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		panic(fmt.Sprintf("ListDatabaseNames Connect exception: %s", err))
	}
	log.Printf("Mongodb connect success -> %s  DatabaseNames->%s", connectionString, names)
	return database
}
