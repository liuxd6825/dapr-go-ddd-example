package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/inventory/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/utils"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/userinterface/rest/inventory/dto"
	"github.com/kataras/iris/v12"
)

type InventoryAssembler struct {
}

//
// AssInventoryCreateAppCmd
// @Description: 创建存货档案
// @receiver a
// @param ictx
// @return *appcmd.InventoryCreateAppCmd 创建存货档案 应用层DTO对象
// @return error 错误
//
func (a *InventoryAssembler) AssInventoryCreateAppCmd(ictx iris.Context) (*appcmd.InventoryCreateAppCmd, error) {
	var request dto.InventoryCreateCommandRequest
	var appCmd appcmd.InventoryCreateAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}

//
// AssInventoryUpdateAppCmd
// @Description: 更新存货档案
// @receiver a
// @param ictx
// @return *appcmd.InventoryUpdateAppCmd 更新存货档案 应用层DTO对象
// @return error 错误
//
func (a *InventoryAssembler) AssInventoryUpdateAppCmd(ictx iris.Context) (*appcmd.InventoryUpdateAppCmd, error) {
	var request dto.InventoryUpdateCommandRequest
	var appCmd appcmd.InventoryUpdateAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}
