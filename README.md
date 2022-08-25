# DDD示例

#### 介绍
DDD架构示例，采用六边形架构，清洁架构，CQRS模式,EventSourcing模式，一套完整的DDD示例。
与Dapr集成，实现微服务治理功能，
EventSourcing事件溯源功能，集成在gitee.com/liuxd6825/dapr项目中。

#### 项目结构说明
[](项目目录说明)

#### 调试

1. dapr -app-port 9010 -dapr-http-port 9011 -app-id cmd-example   -dapr-grpc-port 9012 --enable-metrics=false -config /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/config.yaml -components-path /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/components
2. dapr -app-port 9020 -dapr-http-port 9021 -app-id query-example -dapr-grpc-port 9022 --enable-metrics=false -config /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/config.yaml -components-path /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/components
3. go run ./cmd/cmd-service
4. go run ./cmd/query-service

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx


#### 部署
1. 生成 swagger doc \
   swag init -d ./pkg/query-service/userinterface/rest -o ./swagger/query --parseDependency --parseInternal \
   swag init -d ./pkg/cmd-service/userinterface/rest -o ./swagger/cmd  --parseDependency   --parseInternal

2. 部署Dapr component 文件到k8s上，注册dapr应用与component需要在同一个namespace上。\
   kubectl apply ./config/components/pubsub.yaml\
   kubectl apply ./config/components/applogger-mongo.yaml \
   kubectl apply ./config/components/eventstorage-mongo.yaml \

3. 编译二进制文件\
   make build-linux

4. 生成docker images\
   sudo make docker-build APP_REGISTRY=192.168.64.12 APP_TAG=dapr TARGET_ARCH=arm64

5. 推送images到私仓中\
   make docker-push APP_REGISTRY=192.168.64.12 APP_TAG=dapr TARGET_ARCH=arm64

6. 如果已安装过，卸载已有的k8s安装\
   make uninstall-k8s

7. 安装项目到 k8s\
   make install-k8s


#### 特技



