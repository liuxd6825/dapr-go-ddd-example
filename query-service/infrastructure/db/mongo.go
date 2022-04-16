package db

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
)

var mongoDB *ddd_mongodb.MongoDB

func init() {
	config := &ddd_mongodb.Config{
		Host:         "192.168.64.8:27019",
		DatabaseName: "query-example",
		UserName:     "query-example",
		Password:     "123456",
	}

	mongoDB = ddd_mongodb.NewMongoDB()
	if err := mongoDB.Init(config); err != nil {
		panic(err)
	}
}

func GetMongoDB() *ddd_mongodb.MongoDB {
	return mongoDB
}

func NewSession() ddd_repository.Session {
	return ddd_mongodb.NewSession(mongoDB)
}
