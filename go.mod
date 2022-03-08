module github.com/liuxd6825/dapr-go-ddd-example

go 1.17

require (
	github.com/liuxd6825/dapr-go-ddd-sdk v0.0.0
	github.com/kataras/iris/v12 v12.2.0-alpha7
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
)
replace github.com/liuxd6825/dapr-go-ddd-sdk => ../dapr-go-ddd-sdk