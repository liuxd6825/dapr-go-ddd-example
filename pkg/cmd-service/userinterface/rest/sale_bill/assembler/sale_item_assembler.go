package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/utils"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/userinterface/rest/sale_bill/dto"
	"github.com/kataras/iris/v12"
)

//
// AssSaleItemDeleteCommandDto
// @Description: 删除扫描单
// @receiver a
// @param ictx
// @return *appcmd.SaleItemDeleteAppCmd 删除扫描单 应用层DTO对象
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

//
// AssSaleItemCreateCommandDto
// @Description: 创建扫描文件
// @receiver a
// @param ictx
// @return *appcmd.SaleItemCreateAppCmd 创建扫描文件 应用层DTO对象
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
// AssSaleItemUpdateCommandDto
// @Description: 更新扫描文件
// @receiver a
// @param ictx
// @return *appcmd.SaleItemUpdateAppCmd 更新扫描文件 应用层DTO对象
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
