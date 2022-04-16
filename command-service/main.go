package main

import (
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/userinterface/rest"
	_ "github.com/liuxd6825/dapr-go-ddd-example/common"
	"github.com/liuxd6825/dapr-go-ddd-sdk/applog"
	"github.com/liuxd6825/dapr-go-ddd-sdk/daprclient"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"

	// _ "github.com/liuxd6825/dapr-go-ddd-example/command-service/infrastructure/config"
	_ "github.com/liuxd6825/dapr-go-ddd-example/command-service/infrastructure/idapr"
)

func main() {
	hc, err := daprclient.NewClient("localhost", 9011, 9012)
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
