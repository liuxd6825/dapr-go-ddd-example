package register

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/userinterface/rest/user/facade"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterController
// @Description: 注册的控制器
// @return *[]restapp.Controller
//
func GetRegisterController() *[]restapp.Controller {
	var list []restapp.Controller
	list = append(list, facade.NewUserApi())
	return &list
}
