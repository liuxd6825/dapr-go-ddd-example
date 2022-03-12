package queryservice_impl

import (
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/infrastructure/repository_impl/mongodb"
)

func NewUserQueryService() *queryservice.UserQueryService {
	return queryservice.NewUserQueryService(mongodb.NewUserRepository())
}
