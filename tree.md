.  
|-- cmd      项目main()入口目录  
|   |-- cmd-service    命令服务main()  
|   |-- query-service    查询服务main()  
|-- config      配置目录  
|   |-- dapr      dapr配置目录  
|       |-- components       dapr 组件配置目录     
|-- dddml      DDD建模文件，可通过dapr-ddd-cli工具快速生成项目代码。  
|-- dist       项目可执行文件输出目录   
|-- docker      Dockerfile  
|   |-- cmd       写服务dockerfile   
|   |-- query      读服务dockerfile
|-- k8s       k8s yaml文件   
|-- pkg       项目包目录   
|   |-- cmd-service        写服务  
|   |   |-- application       应用服务层  
|   |   |   |-- internals        内部服务层  
|   |   |       |-- sale_bill      聚合根层    
|   |   |       |   |-- appcmd      应用命令层  
|   |   |       |   |-- assembler       转化层  
|   |   |       |   |-- executor      命令执行层   
|   |   |       |   |-- service      应用服务层     
|   |   |-- domain      领域层  
|   |   |   |-- sale_bill       聚合根层  
|   |   |   |   |-- command      领域命令  
|   |   |   |   |-- event       领域事件  
|   |   |   |   |-- factory     领域工厂   
|   |   |   |   |-- field      领域层Dto  
|   |   |   |   |-- model     领域模式   
|   |   |   |   |-- service      领域服务  
|   |   |-- infrastructure      基础架构层   
|   |   |   |-- base      基础类     
|   |   |   |   |-- application       应用层基类      
|   |   |   |   |   |-- service      应用服务基类   
|   |   |   |   |-- domain         领域层基类     
|   |   |   |       |-- event     领域事件    
|   |   |   |       |-- service       领域服务    
|   |   |   |-- logs        日志  
|   |   |   |-- register       DDD信息注册  
|   |   |   |-- utils       工具  
|   |   |-- userinterface       用户接口层   
|   |       |-- rest       HTTP API 接口层   
|   |           |-- sale_bill        销售单聚合层  
|   |           |   |-- assembler         转化层    
|   |           |   |-- dto        数据传输层  
|   |           |   |-- facade        API层   
|   |-- query-service        读服务   
|       |-- application        应用服务层   
|       |   |-- internals        内部服务层   
|       |       |-- sale_bill  聚合根层   
|       |       |   |-- appquery          查询命令层   
|       |       |   |-- assembler         转化层    
|       |       |   |-- executor          命令执行层   
|       |       |   |   |-- sale_bill_impl    
|       |       |   |   |-- sale_item_impl    
|       |       |   |-- handler         内部领域事件处理层   
|       |       |   |-- service         应用服务层    
|       |-- domain         领域层    
|       |   |-- sale_bill         聚合根层   
|       |   |   |-- factory          构造工厂   
|       |   |   |-- query          查询命令   
|       |   |   |-- repository          仓储定义   
|       |   |   |-- service          服务定义   
|       |   |   |-- view         视图投射类  
|       |-- infrastructure          基础架构层   
|       |   |-- base          基类   
|       |   |   |-- application    
|       |   |   |   |-- handler    
|       |   |   |-- domain    
|       |   |   |   |-- view   
|       |   |   |-- userinterface   
|       |   |       |-- rest  
|       |   |           |-- assembler   
|       |   |           |-- dto   
|       |   |           |-- facade   
|       |   |-- db  数据库   
|       |   |   |-- dao          数据访求对象   
|       |   |   |   |-- mongo_dao     
|       |   |   |-- session         本地事务    
|       |   |-- domain_impl          领域接口实现   
|       |   |   |-- sale_bill   销售接口实现   
|       |   |   |   |-- repository_impl     
|       |   |   |   |   |-- mongodb    
|       |   |   |   |-- service_impl    
|       |   |-- register        ddd 功能注册    
|       |   |-- types        扩展数据类型    
|       |   |-- utils        工具包     
|       |-- userinterface        用户接口层   
|           |-- rest        接口层    
|               |-- sale_bill       聚合根层    
|               |   |-- assembler       转化层   
|               |   |-- dto        数据传输层    
|               |   |-- facade       API层    
|-- swagger         Swagger文件       
|   |-- cmd    
|   |-- query    
|-- test        测试   

176 directories
