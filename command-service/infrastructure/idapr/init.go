package idapr

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/daprclient"
)

var client *daprclient.DaprClient

func init() {
	hc, err := daprclient.NewClient("localhost", 9011, 9012)
	if err != nil {
		panic(err)
	}
	client = hc
}

func GetClient() *daprclient.DaprClient {
	return client
}
