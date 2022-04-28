package mongodb

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

func GetMongoDB() *ddd_mongodb.MongoDB {
	return restapp.GetMongoDB()
}
