package mysql_dao

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

func GetMySqlDB() *gorm.DB {
	return nil
}

func NewSession(isWrite bool) ddd_repository.Session {
	return nil
}

func IsErrRecordNotFound(err error) bool {
	if err == gorm.ErrRecordNotFound {
		return true
	}
	return false
}

func newFindOptions(opt *Options) *options.FindOptions {
	findOptions := &options.FindOptions{}
	findOptions.MaxTime = opt.MaxTime
	findOptions.Sort = opt.Sort
	return findOptions
}

func getSqlWhere(tenantId, filter string) (string, error) {
	p := NewSqlProcess()
	if err := ParseProcess(filter, p); err != nil {
		return "", err
	}
	return p.GetSqlWhere(tenantId).(string), nil
}
