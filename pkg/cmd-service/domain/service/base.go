package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type BaseDomainService struct {
}

func (s *BaseDomainService) ValidateCommand(cmd interface{}) error {
	if cmd == nil {
		return errors.New("command is nil")
	}
	if err := validate.Struct(cmd); err != nil {
		return err
	}
	return nil
}
