package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

var validate = validator.New()

type BaseDomainService struct {
}

func (s *BaseDomainService) ValidateCommand(cmd ddd.Command) error {
	if cmd == nil {
		return errors.New("command is nil")
	}
	return cmd.Validate()
}
