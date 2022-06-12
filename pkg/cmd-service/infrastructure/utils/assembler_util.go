package utils

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_errors"
	"github.com/liuxd6825/dapr-go-ddd-sdk/mapper"
)

//
// AssemblerRequestBody
// @Description: 读取iris body到dto中，并映射dto属性到command命令对象上
// @param ctx iris上下文
// @param request DTO对象
// @param command
// @return err
//
func AssemblerRequestBody(ctx iris.Context, requestDto interface{}, command interface{}) (err error) {
	return AssemblerRequestBodyOption(ctx, requestDto, command, nil, nil)
}

//
// AssemblerRequestBodyOption
// @Description: 读取iris body到dto中，并映射dto属性到command命令对象上
// @param ctx iris上下文
// @param request DTO对象
// @param dto
// @param mapperFunc request对象映射到command。如果为nil, 由mapper.AutoMapper自动映射
// @param afterFunc 后期处理方法，可以进行数据检查，对象映射等工作， 可以为nil
// @return err
//
func AssemblerRequestBodyOption(ctx iris.Context, requestDto interface{}, dto interface{}, mapperFunc func() error, afterFunc func() error) (err error) {
	defer func() {
		if e := ddd_errors.GetRecoverError(recover()); e != nil {
			err = e
		}
	}()
	if err = ctx.ReadBody(requestDto); err != nil {
		return err
	}

	if dto != nil {
		if mapperFunc == nil {
			if err = mapper.Mapper(requestDto, dto); err != nil {
				return err
			}
		} else {
			if err = mapperFunc(); err != nil {
				return err
			}
		}
	}

	if afterFunc != nil {
		err = afterFunc()
	}
	return err
}
