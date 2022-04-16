package mongodb

import (
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/infrastructure/db"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
)

func GetMongoDB() *ddd_mongodb.MongoDB {
	return db.GetMongoDB()
}
