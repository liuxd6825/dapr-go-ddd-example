# dapr-go-ddd-example

#### 目标
框架目标是简化DDD开发难度，使开发人员可直接进行业务开发，不需关心技术细节与实现。实现技术与业务分离，提升开发效率与质量。

#### 介绍
dapr-go-ddd 示例程序，基于liuxd6825/dapr的ddd架构开发工具包。总体架构设计方面部分参考了 Java Axon Framework框架。
本sdk对DDD各层进行了完整的封装，可以进行快速DDD业务开发。

#### 功能
1. 采用ddd六边形架构，CQRS、Aggregate、EventSourcing等架构模式
2. 框架采用多租户模式设计，数据和方法中预留TenantId属性或参数。
3. 框架采用接口、链式、函数式编程，可支持多种数据库扩展，目前仅支持MongoDB。
4. 采用iris实现Http UserInterface层封装。
5. 采用RestSQL查询语言，更好的支持复杂查询。
6. 对事件定义、事件注册、事件存储、事件发送、事件溯源等进行封装。
7. 事件溯源设计采用契约方式定义事件处理器EventHandler
8. 支持事件的多版本管理，解决事件变更问题。
9. sdk支持CQRS，可分别部署Command服务与Query服务。提高计算资源使用率。
10. 对Repository进行了封装，采用接口与链式方式编程，可支持多种数据库，
11. 优化了前端调用Command服务，异步交互的问题。前端调用更友好。
12. 优化事件处理机制，解决领域事件容错问题。
     1. 事件处理错误
     2. 事件顺序执行问题
12.采用actor模型进行事件溯源等并发问题。


#### 调试

1. dapr -app-port 9010 -dapr-http-port 9011 -app-id cmd-example -dapr-grpc-port 9012 --enable-metrics=false -config /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/config.yaml -components-path /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/components
2. dapr -app-port 9020 -dapr-http-port 9021 -dapr-grpc-port 9022 -app-id query-example --enable-metrics=false -config /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/config.yaml -components-path /Users/lxd/go/src/github.com/liuxd6825/dapr-go-ddd-example/config/dapr/components
3. go run ./cmd/cmd-service
4. go run ./cmd/query-service
 

#### 使用说明

1. 生成 swagger doc \
   swag init -d ./pkg/query-service/userinterface/rest -o ./swagger/query --parseDependency --parseInternal \
   swag init -d ./pkg/cmd-service/userinterface/rest -o ./swagger/cmd  --parseDependency   --parseInternal




#### 部署
1. 部署Dapr component 文件到k8s上，注册dapr应用与component需要在同一个namespace上。\
   kubectl apply ./config/components/pubsub.yaml\
   kubectl apply ./config/components/applogger-mongo.yaml \
   kubectl apply ./config/components/eventstorage-mongo.yaml \


3. 编译二进制文件\
   make build-linux


4. 生成docker images\
   sudo make docker-build APP_REGISTRY=192.168.64.12 APP_TAG=dapr TARGET_ARCH=arm64


5. 推送images到私仓中\
   make docker-push APP_REGISTRY=192.168.64.12 APP_TAG=dapr TARGET_ARCH=arm64

   
7. 如果已安装过，卸载已有的k8s安装\
   make uninstall-k8s


8. 安装项目到 k8s\
   make install-k8s


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

