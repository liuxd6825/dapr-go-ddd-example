package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/appcmd"
	user_event "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/field"
	"github.com/liuxd6825/dapr-go-ddd-sdk/test"
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/idutils"
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/randomutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserCommandAppService_UserCreate(t *testing.T) {
	ctx := context.Background()

	err := test.InitCommand(test.NewOption().SetEventTypes(user_event.GetRegisterEventTypes()))
	assert.NoError(t, err)

	cmd := &appcmd.UserCreateAppCmd{
		CommandId:   idutils.NewId(),
		IsValidOnly: false,
		Data: field.UserCreateFields{
			Id:       idutils.NewId(),
			TenantId: "test",
			Email:    "example@example.com",
			Name:     randomutils.NameCN(),
		},
	}

	service := NewUserCommandAppService()
	err = service.UserCreate(ctx, cmd)

	assert.NoError(t, err)
}

func init() {
	test.EnvName = "dev"
}
