package hmgdb

import (
	"os"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestDb(t *testing.T) {
	tests := []struct {
		name string
		want *mongo.Database
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Db(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDatabase(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *mongo.Database
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDatabase(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetConnectString(t *testing.T) {
	type args struct {
		name             string
		connectionString string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSetDatabase(t *testing.T) {
	type args struct {
		name string
		v    *mongo.Database
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_connect(t *testing.T) {
	connectionString := os.Getenv("MS_MONGODB_CONNECT")

	type args struct {
		connectionString string
	}
	tests := []struct {
		name string
		args args
		want *mongo.Database
	}{
		// TODO: Add test cases.
		{name: DEFAULT, args: args{connectionString: connectionString}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.connectionString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
