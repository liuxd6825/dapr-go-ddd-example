# DDD示例

#### 示例介绍
DDD架构示例，采用六边形架构，清洁架构，CQRS模式,EventSourcing模式，一套完整的DDD示例。
与Dapr集成，实现微服务治理功能，
EventSourcing事件溯源功能，集成在gitee.com/liuxd6825/dapr项目中。

#### 框架目标
框架目标是简化DDD开发难度，使开发人员可直接进行业务开发，不需关心技术细节与实现。实现技术与业务分离，提升开发效率与质量。

#### 框架功能
1. 采用ddd六边形架构，CQRS、Aggregate、EventSourcing等架构模式
2. 框架采用多租户模式设计，数据和方法中预留TenantId属性或参数。
3. 可使用yaml文件对ddd进行建模，并通过dapr-ddd-cli工具，快速生成DDD项目代码，
4. 框架采用接口、链式、函数式编程。
5. 采用iris实现Http UserInterface层封装。
6. 采用RestSQL查询语言，更好的支持复杂查询。
7. 对事件定义、事件注册、事件存储、事件发送、事件溯源等进行封装。
8. 事件溯源设计采用契约方式定义事件处理器EventHandler
9. 支持事件的多版本管理，解决事件变更问题。
10. sdk支持CQRS，可分别部署Command服务与Query服务。提高计算资源使用率。
11. 优化了前端调用Command服务，异步交互的问题。前端调用更友好。
12. 优化事件处理机制，解决领域事件容错问题。
13. 事件处理错误
14. 事件顺序执行问题
    12.采用actor模型进行事件溯源等并发问题。
#### 项目结构
|-- cmd              项目入口  
|   |-- cmd-service                    命令服务 main() 函数定义   
|   |-- query-service                  查询服务 main() 函数定义    
|-- config                    配置目录     
|   |-- dapr                  dapr配置目录     
|   |   |-- components                  dapr 组件配置目录        
|-- dddml                  DDD建模文件，可通过dapr-ddd-cli工具快速生成项目代码。  
|-- dist                       项目可执行文件输出目录     
|-- docker                 Dockerfile    
|   |-- cmd                     写服务dockerfile     
|   |-- query                  读服务dockerfile     
|-- k8s                  k8s yaml文件     
|-- pkg                  项目包目录     
|   |-- cmd-service                    写服务    
|   |   |-- application                  应用服务层    
|   |   |   |-- internals                   内部服务层   
|   |   |   |   +-- inventory                       存货聚合层    
|   |   |   |   +-- user                              用户聚合层    
|   |   |   |   |-- sale_bill                         销售单聚合层      
|   |   |   |   |   |-- appcmd                         应用命令层      
|   |   |   |   |   |-- assembler                     转化层      
|   |   |   |   |   |-- executor                        命令执行层     
|   |   |   |   |   |-- service                          应用服务层       
|   |   |-- domain                  领域层     
|   |   |   +-- inventory                  存货聚合层      
|   |   |   +-- user                          用户聚合层     
|   |   |   |-- sale_bill                     销售单聚合层    
|   |   |   |   |-- command                          领域命令    
|   |   |   |   |-- event                                  领域事件  
|   |   |   |   |-- factory                               领域工厂   
|   |   |   |   |-- field                                   领域层Dto  
|   |   |   |   |-- model                                领域模式   
|   |   |   |   |-- service                              领域服务  
|   |   |-- infrastructure                   基础架构层   
|   |   |-- userinterface                   用户接口层   
|   |   |   |-- rest                   HTTP API 接口层   
|   |   |   |   +-- inventory                  存货聚合层  
|   |   |   |   +-- user                          用户聚合层  
|   |   |   |   |-- sale_bill                      销售单聚合层  
|   |   |   |   |   |-- assembler                    转化层    
|   |   |   |   |   |-- dto                               数据传输层  
|   |   |   |   |   |-- facade                         API层   
|   |-- query-service           读服务   
|   |   |-- application                    应用服务层   
|   |   |   |-- internals                          内部服务层  
|   |   |   |   +-- inventory                        存货聚合层  
|   |   |   |   +-- user                                用户聚合层   
|   |   |   |   |-- sale_bill                           销售单聚合层     
|   |   |   |   |   |-- appquery                          查询命令层   
|   |   |   |   |   |-- assembler                         转化层    
|   |   |   |   |   |-- executor                             命令执行层   
|   |   |   |   |   |   |-- sale_bill_impl                      销售单命令接口实现  
|   |   |   |   |   |   |-- sale_item_impl                    销售明细命令接口实现  
|   |   |   |   |   |-- handler                  内部领域事件处理层   
|   |   |   |   |   |-- service                   应用服务层    
|   |   |-- domain                   领域层  
|   |   |   +-- inventory                 存货聚合层  
|   |   |   +-- user                        用户聚合层  
|   |   |   |-- sale_bill                    聚合根层    
|   |   |   |   |-- factory                       构造工厂   
|   |   |   |   |-- query                         查询命令   
|   |   |   |   |-- repository                  仓储定义   
|   |   |   |   |-- service                       服务定义   
|   |   |   |   |-- view                           视图投射类  
|   |   |-- infrastructur             基础架构层   
|   |   |   |-- base                           基类   
|   |   |   |-- db                               数据库   
|   |   |   |-- domain_impl                 领域接口实现     
|   |   |   |   +-- inventory                    存货聚合层    
|   |   |   |   +-- user                            用户聚合层    
|   |   |   |   |-- sale_bill                        销售接口实现     
|   |   |   |   |   |-- repository_impl     
|   |   |   |   |   |-- service_impl     
|   |   |   |-- register                       ddd 功能注册      
|   |   |   |-- types                           扩展数据类型      
|   |   |   |-- utils                            工具包       
|   |   |-- userinterface           用户接口层     
|   |   |   |-- rest                          接口层    
|   |   |   |   +-- inventory                  存货聚合层    
|   |   |   |   +-- user                          用户聚合层  
|   |   |   |   |-- sale_bill                     聚合根层    
|   |   |   |   |   |-- assembler                  转化层   
|   |   |   |   |   |-- dto                              数据传输层    
|   |   |   |   |   |-- facade                        API层    
|-- swagger          Swagger     
|   |-- cmd                        命令服务swagger文件  
|   |-- query                      查询服务swagger文件  
|-- test                    测试

#### DDDML定义
模型示例，详细请查看dddml目录


       aggregates:
         # 定义聚合根-存货档案
         Inventory:
           description: "存货档案"
           # 定义存货档案聚合根的属性
           properties:
           Id:
             type: string
             description: "Id"
           Name:
             type: string
             description: "名称"
           Spec:
             type: string
             description: "规格"
           Brand:
             type: string
             description: "品牌"
           Keywords:
             type: string
             description: "搜索关键字"
         # 定义实体类
         entities:
         
         # 定义命令
         commands:
           InventoryCreateCommand:
             description: "创建存货档案"
             fields:
               properties:
                 Id:
                   type: string
                   description: "Id"
                 Name:
                   type: string
                   description: "名称"
                 Spec:
                   type: string
                   description: "规格"
                 Brand:
                   type: string
                   description: "品牌"
                 Keywords:
                   type: string
                   description: "搜索关键字"
             InventoryUpdateCommand:
               description: "更新存货档案"
               fields:
                 properties:
                   Id:
                     type: string
                     description: "Id"
                   Name:
                     type: string
                     description: "名称"
                   Spec:
                     type: string
                     description: "规格"
                   Brand:
                     type: string
                     description: "品牌"
                   Keywords:
                     type: string
                     description: "搜索关键字"

dddml文件说明:
- aggregates_inventory.yaml            存货聚合根定义
- aggregates_sale_bill.yaml             销售单聚合根定义
- aggregates_sale_item.yaml           销售明细实体定义
- aggregates_user.yaml                    用户聚合根定义
- configuration.yaml                          系统环境配置
- types.yaml                                      系统数据类型定义

#### 示例调试

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



