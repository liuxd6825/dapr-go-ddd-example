package queryappservice

import (
	dapr "github.com/dapr/go-sdk/client"
)

var client dapr.Client
var appId = "query-example"
var apiVersion = "v1.0"

func init() {
	var err error
	client, err = dapr.NewClientWithPort("9012")
	if err != nil {
		panic(err)
	}
}
