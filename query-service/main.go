package main

import (
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/userinterface/rest"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func main() {
	eventStorage := ddd.NewDaprEventStorage("localhost", 3501, "pubsub")

	ddd.Init(eventStorage)
	rest.Init(5001)
}
