package mongodb

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
)

type RepositoryOptions struct {
	mongoDB *ddd_mongodb.MongoDB
}

func NewRepositoryOptions() *RepositoryOptions {
	return &RepositoryOptions{}
}

func (o *RepositoryOptions) SetMongoDB(mongoDB *ddd_mongodb.MongoDB) *RepositoryOptions {
	o.mongoDB = mongoDB
	return o
}

func (o *RepositoryOptions) Merge(opts ...*RepositoryOptions) {
	if opts != nil {
		for _, item := range opts {
			if item.mongoDB != nil {
				o.mongoDB = item.mongoDB
			}
		}
	}
	if o.mongoDB == nil {
		o.mongoDB = GetMongoDB()
	}
}
