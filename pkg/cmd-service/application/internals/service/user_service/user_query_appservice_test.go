package user_service

import (
	"context"
	"testing"
)

func TestGetUserByUserId(t *testing.T) {
	userView, isFound, err := GetUserQueryAppService().GetById(context.Background(), "001", "040501-1")
	if err != nil {
		t.Error(err)
	} else {
		if isFound {
			println("is found userId=" + userView.Id)
		} else {
			println("not found ")
		}
	}
}
