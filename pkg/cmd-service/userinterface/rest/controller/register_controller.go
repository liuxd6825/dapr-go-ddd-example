package controller

import "github.com/liuxd6825/dapr-go-ddd-sdk/restapp"

//
// GetRegisterController
// @Description: 注册的控制器
// @return *[]restapp.Controller
//
func GetRegisterController() *[]restapp.Controller {
	var list []restapp.Controller
	list = append(list, NewUserController())
	return &list
}
