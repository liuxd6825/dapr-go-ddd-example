package register

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/infrastructure/register"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterEventTypes
// @Description: 注册的控制器
// @return *[]restapp.RegisterEventType
//
func GetRegisterEventTypes() *[]restapp.RegisterEventType {
	return register.GetRegisterEventTypes()
}
