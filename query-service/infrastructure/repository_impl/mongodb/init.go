package mongodb

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
)

var mongoDB *ddd_mongodb.MongoDB

func Init(config *ddd_mongodb.Config) error {
	mongoDB = ddd_mongodb.NewMongoDB()
	return mongoDB.Init(config)
}

func GetMongoDB() *ddd_mongodb.MongoDB {
	return mongoDB
}

func UseSession(ctx context.Context, fun ddd_repository.SessionFunc) error {
	return ddd_mongodb.UseSession(ctx, mongoDB, fun)
}
