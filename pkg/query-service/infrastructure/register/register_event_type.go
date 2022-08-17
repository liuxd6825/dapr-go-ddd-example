package register

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/register"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterEventType
// @Description: 获取所有注册事件类型
// @return []restapp.RegisterEventType
//
func GetRegisterEventType() []restapp.RegisterEventType {
	return register.GetRegisterEventType()
}
