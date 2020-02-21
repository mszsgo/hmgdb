package hmgdb

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrConnectionString = errors.New("Mongodb connection string Error")
)

// 获取Mongodb数据库连接
func ConnectDatabase(connectionString string) (*mongo.Database, error) {
	if connectionString == "" {
		return nil, ErrConnectionString
	}
	// 从连接字符串中截取数据库名称
	// 示例连接字符串
	// mongodb://user:pass@11.168.200.112:27017/dbName?authSource=authDb&authMechanism=SCRAM-SHA-1
	dbName := (strings.Split((strings.Split(connectionString, "/"))[3], "?"))[0]
	if dbName == "" {
		return nil, ErrConnectionString
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}
	database := client.Database(dbName)
	return database, nil
}
