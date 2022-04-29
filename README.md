# dapr-go-ddd-example

#### 介绍
DDD示例程序。

#### 调试

1. dapr -app-port 9010 -dapr-http-port 9011 -app-id cmd-example -dapr-grpc-port 9012 --enable-metrics=false -config /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/config.yaml -components-path /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/components
2. dapr -app-port 9020 -dapr-http-port 9021 -dapr-grpc-port 9022 -app-id query-example --enable-metrics=false -config /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/config.yaml -components-path /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/components
3. go run ./cmd/cmd-service
4. go run ./cmd/query-service

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx


#### 部署

1. make build-linux
2. sudo make docker-build APP_REGISTRY=192.168.64.12 APP_TAG=dapr TARGET_ARCH=arm64
3. make docker-push APP_REGISTRY=192.168.64.12 APP_TAG=dapr TARGET_ARCH=arm64
4. docker push 192.168.64.12/cmd-service:dapr-linux-arm64
5. make build-k8s

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1. 使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2. Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3. 你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4. [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5. Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6. Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)

