package model

import (
	"fmt"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/fields"
)

type AddressEntity struct {
	Id       string `json:"id"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Location string `json:"location"`
}

func NewAddressEntity(fields *fields.AddressFields) *AddressEntity {
	return &AddressEntity{
		Id:       fields.Id,
		Province: fields.Province,
		City:     fields.City,
		Area:     fields.Area,
		Location: fields.Location,
	}
}

func (a *AddressEntity) ToString() string {
	return fmt.Sprintf("%s省 %s市 %s区 详细：%s", a.Province, a.City, a.Area, a.Location)
}
