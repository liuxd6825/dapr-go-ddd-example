package mongodb

import (
	"context"
	"github.com/google/uuid"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/base/domain/repository/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
	"testing"
)

func Test_UserRepository_Inserts(t *testing.T) {
	mogodb := newMongoDb()
	userRepos := NewUserRepository(mongodb.NewRepositoryOptions().SetMongoDB(mogodb))
	id := uuid.New().String()
	userView := &projection.UserView{
		Id:        id,
		TenantId:  "001",
		UserName:  "test_user",
		UserCode:  id,
		Email:     "test@163.com",
		Telephone: "1222",
		Address:   "address",
	}
	res, err := userRepos.CreateById(context.Background(), userView)
	if err != nil {
		t.Error(err)
	} else {
		println(res)
	}
}

func Test_UserRepository_Search(t *testing.T) {
	mogodb := newMongoDb()

	queryOpt := ddd_repository.NewFindPagingQueryOptions().SetFilter("id=='001'")
	query := ddd_repository.NewFindPagingQuery("001", queryOpt)

	userRepos := NewUserRepository(mongodb.NewRepositoryOptions().SetMongoDB(mogodb))
	res, ok, err := userRepos.FindPaging(context.Background(), query).Result()
	t.Error(err)
	if ok {
		println(res)
	}
}

func newMongoDb() *ddd_mongodb.MongoDB {
	mongoDb, err := ddd_mongodb.NewMongoDB(&ddd_mongodb.Config{
		Host:         "192.168.64.8:27018, 192.168.64.8:27019, 192.168.64.8:27020",
		ReplicaSet:   "mongors",
		DatabaseName: "query-example",
		UserName:     "query-example",
		Password:     "123456",
		MaxPoolSize:  20,
	})
	if err != nil {
		panic(err)
	}
	return mongoDb
}
