package main

import (
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/userinterface/rest"
	_ "github.com/liuxd6825/dapr-go-ddd-example/common"
	"github.com/liuxd6825/dapr-go-ddd-sdk/applog"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/httpclient"
)

func main() {
	hc, err := httpclient.NewHttpClient("localhost", 9011)
	if err != nil {
		panic(err)
	}
	applog.Init(hc, "command-example", applog.DEBUG)
	eventStorage, err := ddd.NewDaprEventStorage(hc, ddd.PubsubName("pubsub"))
	if err != nil {
		panic(err)
	}
	ddd.RegisterEventStorage("", eventStorage)
	rest.Init(9010)
}
