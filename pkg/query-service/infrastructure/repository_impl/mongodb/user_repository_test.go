package mongodb

import (
	"context"
	"fmt"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
	"golang.org/x/exp/constraints"
	"testing"
)

func Test_UserRepository_Search(t *testing.T) {
	initTestDB()
	query := &ddd_repository.ListQuery{
		TenantId: "001",
		Filter:   "id=='001'",
	}
	userRepos := NewUserRepository()
	res, ok, err := userRepos.GetList(context.Background(), query)
	t.Error(err)
	if ok {
		println(res)
	}
}

func initTestDB() {
	err := Init(&ddd_mongodb.Config{
		Host:         "192.168.64.4",
		DatabaseName: "example-query-service",
		UserName:     "dapr",
		Password:     "123456",
	})
	if err != nil {
		panic(err)
	}
}

func Sum[V constraints.Float | constraints.Integer](m ...V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func Test_Sum(t *testing.T) {
	fmt.Println(Sum([]int64{1, 2, 3, 4, 5, 6}...))

}
