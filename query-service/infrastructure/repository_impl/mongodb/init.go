package mongodb

import "github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"

var mongoDB *ddd_mongodb.MongoDB

func Init(config *ddd_mongodb.Config) error {
	mongoDB = ddd_mongodb.NewMongoDB()
	return mongoDB.Init(config)
}

func GetMongoDB() *ddd_mongodb.MongoDB {
	return mongoDB
}
