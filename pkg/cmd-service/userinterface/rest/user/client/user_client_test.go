package client

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserCmdFeign_FindAggregateById(t *testing.T) {
	err := test.InitCommand()
	assert.NoError(t, err)
	if err != nil {
		return
	}

	ctx := context.Background()

	req := &FindAggregateByIdConfig{
		TenantId: "test",
		Id:       "fiUBX28y2BU4rAbytgXE8Gu",
	}

	userFeign, err := NewUserCmdFeign()
	assert.NoError(t, err)
	if err != nil {
		return
	}

	resp, err := userFeign.FindAggregateById(ctx, req)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	t.Log(resp.Name)

}
