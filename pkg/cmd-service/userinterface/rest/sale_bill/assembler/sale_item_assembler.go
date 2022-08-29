package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/utils"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/userinterface/rest/sale_bill/dto"
	"github.com/kataras/iris/v12"
)

//
// AssSaleItemCreateAppCmd
// @Description: 添加明细
// @receiver a
// @param ictx
// @return *appcmd.SaleItemCreateAppCmd 添加明细 应用层DTO对象
// @return error 错误
//
func (a *SaleBillAssembler) AssSaleItemCreateAppCmd(ictx iris.Context) (*appcmd.SaleItemCreateAppCmd, error) {
	var request dto.SaleItemCreateCommandRequest
	var appCmd appcmd.SaleItemCreateAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}

//
// AssSaleItemUpdateAppCmd
// @Description: 更新明细
// @receiver a
// @param ictx
// @return *appcmd.SaleItemUpdateAppCmd 更新明细 应用层DTO对象
// @return error 错误
//
func (a *SaleBillAssembler) AssSaleItemUpdateAppCmd(ictx iris.Context) (*appcmd.SaleItemUpdateAppCmd, error) {
	var request dto.SaleItemUpdateCommandRequest
	var appCmd appcmd.SaleItemUpdateAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}

//
// AssSaleItemDeleteAppCmd
// @Description: 删除销售明细项
// @receiver a
// @param ictx
// @return *appcmd.SaleItemDeleteAppCmd 删除销售明细项 应用层DTO对象
// @return error 错误
//
func (a *SaleBillAssembler) AssSaleItemDeleteAppCmd(ictx iris.Context) (*appcmd.SaleItemDeleteAppCmd, error) {
	var request dto.SaleItemDeleteCommandRequest
	var appCmd appcmd.SaleItemDeleteAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}
