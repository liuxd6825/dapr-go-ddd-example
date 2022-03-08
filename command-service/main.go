package main

import (
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/userinterface/rest"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func main() {
	eventStorage := ddd.NewDaprEventStorage("localhost", 3500, "pubsub")

	ddd.Init(eventStorage)
	rest.Init(3500)
}
