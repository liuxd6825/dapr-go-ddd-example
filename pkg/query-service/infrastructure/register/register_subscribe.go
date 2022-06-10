package register

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/application/internals/handler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterSubscribe
// @Description: 注册领域事件监听器
// @return *[]restapp.RegisterSubscribe
//
func GetRegisterSubscribe() *[]restapp.RegisterSubscribe {
	var list []restapp.RegisterSubscribe
	list = append(list, handler.NewUserSubscribes())
	return &list
}
