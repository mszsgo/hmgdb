package hmgdb

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UlUser struct {
	Uid    string `bson:"uid"`
	Email  string `bson:"email"`
	Mobile string `bson:"mobile"`
}

// 测试列表查询
func TestFind(t *testing.T) {

	connectionString := os.Getenv("MS_MONGODB_CONNECT")
	SetConnectString(DEFAULT, connectionString)
	collection := GetDatabase(DEFAULT).Collection("ul_user")

	var user []UlUser
	type args struct {
		ctx          context.Context
		filter       interface{}
		returnResult interface{}
		opts         []*options.FindOptions
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: DEFAULT, args: args{ctx: nil, filter: bson.M{}, returnResult: &user}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args
			Find(args.ctx, collection, args.filter, args.returnResult, args.opts...)
		})
	}
	b, _ := json.Marshal(user)
	fmt.Println(string(b))
}
