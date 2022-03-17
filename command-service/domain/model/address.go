package model

import "fmt"

type AddressValue struct {
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Location string `json:"location"`
}

func NewAddressValue() *AddressValue {
	return &AddressValue{}
}

func (a *AddressValue) ToString() string {
	return fmt.Sprintf("%s省 %s市 %s区 详细：%s", a.Province, a.City, a.Area, a.Location)
}
