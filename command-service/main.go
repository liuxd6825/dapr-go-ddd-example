package main

import (
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/userinterface/rest"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func main() {
	eventStorage, err := ddd.NewDaprEventStorage("localhost", 9011, ddd.PubsubName("pubsub"))
	if err != nil {
		panic(err)
	}
	ddd.RegisterEventStorage("", eventStorage)
	rest.Init(9010)
}
