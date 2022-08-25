.
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   ├── cmd-service
│   │   └── main.go
│   └── query-service
│       └── main.go
├── config
│   ├── cmd-config.yaml
│   ├── dapr
│   │   ├── components
│   │   │   ├── applogger-mongo.yaml
│   │   │   ├── eventstorage-mongo.yaml_
│   │   │   ├── eventstorage-mysql.yaml
│   │   │   ├── pubsub-rabbitmq.yaml
│   │   │   ├── pubsub.yaml
│   │   │   └── statestore.yaml
│   │   └── config.yaml
│   └── query-config.yaml
├── dddml
│   ├── aggregates_inventory.yaml
│   ├── aggregates_sale_bill.yaml
│   ├── aggregates_sale_item.yaml
│   ├── aggregates_user.yaml
│   ├── configuration.yaml
│   ├── types.yaml
│   └── valueObjects.yaml
├── dist
│   └── linux_arm64
│       └── release
│           ├── cmd-service
│           ├── config
│           │   ├── cmd-config.yaml
│           │   ├── dapr
│           │   │   ├── components
│           │   │   │   ├── applogger-mongo.yaml
│           │   │   │   ├── eventstorage-mongo.yaml
│           │   │   │   ├── pubsub-rabbitmq.yaml
│           │   │   │   ├── pubsub.yaml
│           │   │   │   └── statestore.yaml
│           │   │   └── config.yaml
│           │   └── query-config.yaml
│           └── query-service
├── docker
│   ├── cmd
│   │   └── Dockerfile
│   ├── docker.mk
│   └── query
│       └── Dockerfile
├── gen
├── go.mod
├── go.sum
├── k8s
│   ├── cmd-service.yaml
│   └── query-service.yaml
├── pkg
│   ├── cmd-service
│   │   ├── application
│   │   │   └── internals
│   │   │       ├── inventory
│   │   │       │   ├── appcmd
│   │   │       │   │   └── inventory_appcmd.go
│   │   │       │   ├── assembler
│   │   │       │   │   └── inventory_assembler.go
│   │   │       │   ├── executor
│   │   │       │   │   ├── find_aggreage_id_executor.go
│   │   │       │   │   ├── inventory_create_command_executor.go
│   │   │       │   │   ├── inventory_update_command_executor.go
│   │   │       │   │   └── x_init.go
│   │   │       │   └── service
│   │   │       │       └── inventory_command_app_service.go
│   │   │       ├── sale_bill
│   │   │       │   ├── appcmd
│   │   │       │   │   ├── sale_bill_appcmd.go
│   │   │       │   │   └── sale_item_appcmd.go
│   │   │       │   ├── assembler
│   │   │       │   │   ├── sale_bill_assembler.go
│   │   │       │   │   └── sale_item_assembler.go
│   │   │       │   ├── executor
│   │   │       │   │   ├── find_aggreage_id_executor.go
│   │   │       │   │   ├── sale_bill_confirm_command_executor.go
│   │   │       │   │   ├── sale_bill_create_command_executor.go
│   │   │       │   │   ├── sale_bill_delete_command_executor.go
│   │   │       │   │   ├── sale_bill_update_command_executor.go
│   │   │       │   │   ├── sale_item_create_command_executor.go
│   │   │       │   │   ├── sale_item_delete_command_executor.go
│   │   │       │   │   ├── sale_item_update_command_executor.go
│   │   │       │   │   └── x_init.go
│   │   │       │   └── service
│   │   │       │       └── sale_bill_command_app_service.go
│   │   │       └── user
│   │   │           ├── appcmd
│   │   │           │   └── user_appcmd.go
│   │   │           ├── assembler
│   │   │           │   └── user_assembler.go
│   │   │           ├── executor
│   │   │           │   ├── find_aggreage_id_executor.go
│   │   │           │   ├── user_create_command_executor.go
│   │   │           │   ├── user_delete_command_executor.go
│   │   │           │   ├── user_update_command_executor.go
│   │   │           │   └── x_init.go
│   │   │           └── service
│   │   │               └── user_command_app_service.go
│   │   ├── domain
│   │   │   ├── inventory
│   │   │   │   ├── command
│   │   │   │   │   ├── inventory_create_command.go
│   │   │   │   │   └── inventory_update_command.go
│   │   │   │   ├── event
│   │   │   │   │   ├── event_type.go
│   │   │   │   │   ├── inventory_create_event.go
│   │   │   │   │   └── inventory_update_event.go
│   │   │   │   ├── factory
│   │   │   │   │   └── event_factory.go
│   │   │   │   ├── field
│   │   │   │   │   ├── inventory_create_fields.go
│   │   │   │   │   └── inventory_update_fields.go
│   │   │   │   ├── model
│   │   │   │   │   ├── inventory_aggregate.go
│   │   │   │   │   ├── inventory_aggregate_command.go
│   │   │   │   │   └── inventory_aggregate_event.go
│   │   │   │   └── service
│   │   │   │       └── inventory_domain_service.go
│   │   │   ├── sale_bill
│   │   │   │   ├── command
│   │   │   │   │   ├── sale_bill_confirm_command.go
│   │   │   │   │   ├── sale_bill_create_command.go
│   │   │   │   │   ├── sale_bill_delete_command.go
│   │   │   │   │   ├── sale_bill_update_command.go
│   │   │   │   │   ├── sale_item_create_command.go
│   │   │   │   │   ├── sale_item_delete_command.go
│   │   │   │   │   └── sale_item_update_command.go
│   │   │   │   ├── event
│   │   │   │   │   ├── event_type.go
│   │   │   │   │   ├── sale_bill_confirm_event.go
│   │   │   │   │   ├── sale_bill_create_event.go
│   │   │   │   │   ├── sale_bill_delete_event.go
│   │   │   │   │   ├── sale_bill_update_event.go
│   │   │   │   │   ├── sale_item_create_event.go
│   │   │   │   │   ├── sale_item_delete_event.go
│   │   │   │   │   └── sale_item_update_event.go
│   │   │   │   ├── factory
│   │   │   │   │   └── event_factory.go
│   │   │   │   ├── field
│   │   │   │   │   ├── sale_bill_confirm_fields.go
│   │   │   │   │   ├── sale_bill_create_fields.go
│   │   │   │   │   ├── sale_bill_delete_fields.go
│   │   │   │   │   ├── sale_bill_statue_enum.go
│   │   │   │   │   ├── sale_bill_update_fields.go
│   │   │   │   │   ├── sale_item_create_fields.go
│   │   │   │   │   ├── sale_item_create_item.go
│   │   │   │   │   ├── sale_item_delete_fields.go
│   │   │   │   │   ├── sale_item_delete_item.go
│   │   │   │   │   ├── sale_item_update_fields.go
│   │   │   │   │   └── sale_item_update_item.go
│   │   │   │   ├── model
│   │   │   │   │   ├── sale_bill_aggregate.go
│   │   │   │   │   ├── sale_bill_aggregate_command.go
│   │   │   │   │   ├── sale_bill_aggregate_event.go
│   │   │   │   │   ├── sale_item_entity.go
│   │   │   │   │   └── sale_item_items.go
│   │   │   │   └── service
│   │   │   │       └── sale_bill_domain_service.go
│   │   │   └── user
│   │   │       ├── command
│   │   │       │   ├── user_create_command.go
│   │   │       │   ├── user_delete_command.go
│   │   │       │   └── user_update_command.go
│   │   │       ├── event
│   │   │       │   ├── event_type.go
│   │   │       │   ├── user_create_event.go
│   │   │       │   ├── user_delete_event.go
│   │   │       │   └── user_update_event.go
│   │   │       ├── factory
│   │   │       │   └── event_factory.go
│   │   │       ├── field
│   │   │       │   ├── user_create_fields.go
│   │   │       │   ├── user_delete_fields.go
│   │   │       │   └── user_update_fields.go
│   │   │       ├── model
│   │   │       │   ├── user_aggregate.go
│   │   │       │   ├── user_aggregate_command.go
│   │   │       │   └── user_aggregate_event.go
│   │   │       └── service
│   │   │           └── user_domain_service.go
│   │   ├── infrastructure
│   │   │   ├── base
│   │   │   │   ├── application
│   │   │   │   │   └── service
│   │   │   │   │       └── base_query_service.go
│   │   │   │   └── domain
│   │   │   │       ├── event
│   │   │   │       │   └── base_event.go
│   │   │   │       └── service
│   │   │   │           └── base_command_service.go
│   │   │   ├── logs
│   │   │   │   └── logs.go
│   │   │   ├── register
│   │   │   │   ├── register_aggregate_type.go
│   │   │   │   ├── register_event_type.go
│   │   │   │   └── register_res_api.go
│   │   │   └── utils
│   │   │       └── assembler_util.go
│   │   └── userinterface
│   │       └── rest
│   │           ├── inventory
│   │           │   ├── assembler
│   │           │   │   └── inventory_assembler.go
│   │           │   ├── dto
│   │           │   │   └── inventory_dto.go
│   │           │   └── facade
│   │           │       └── inventory_api.go
│   │           ├── sale_bill
│   │           │   ├── assembler
│   │           │   │   ├── sale_bill_assembler.go
│   │           │   │   └── sale_item_assembler.go
│   │           │   ├── dto
│   │           │   │   ├── sale_bill_dto.go
│   │           │   │   └── sale_item_dto.go
│   │           │   └── facade
│   │           │       ├── sale_bill_api.go
│   │           │       └── sale_item_api.go
│   │           ├── swagger.go
│   │           └── user
│   │               ├── assembler
│   │               │   └── user_assembler.go
│   │               ├── dto
│   │               │   └── user_dto.go
│   │               └── facade
│   │                   └── user_api.go
│   └── query-service
│       ├── application
│       │   └── internals
│       │       ├── inventory
│       │       │   ├── appquery
│       │       │   │   └── inventory_appquery.go
│       │       │   ├── assembler
│       │       │   │   └── inventory_assembler.go
│       │       │   ├── executor
│       │       │   │   ├── inventory_executor.go
│       │       │   │   └── inventory_impl
│       │       │   │       ├── create_executor.go
│       │       │   │       ├── create_many_executor.go
│       │       │   │       ├── delete_all_executor.go
│       │       │   │       ├── delete_by_id_executor.go
│       │       │   │       ├── delete_many_executor.go
│       │       │   │       ├── find_all_executor.go
│       │       │   │       ├── find_by_id_executor.go
│       │       │   │       ├── find_by_ids_executor.go
│       │       │   │       ├── find_paging_executor.go
│       │       │   │       ├── update_all_executor.go
│       │       │   │       ├── update_executor.go
│       │       │   │       └── x_init.go
│       │       │   ├── handler
│       │       │   │   └── inventory_query_handler.go
│       │       │   └── service
│       │       │       └── inventory_query_appservice.go
│       │       ├── sale_bill
│       │       │   ├── appquery
│       │       │   │   ├── sale_bill_appquery.go
│       │       │   │   └── sale_item_appquery.go
│       │       │   ├── assembler
│       │       │   │   ├── sale_bill_assembler.go
│       │       │   │   └── sale_item_assembler.go
│       │       │   ├── executor
│       │       │   │   ├── sale_bill_executor.go
│       │       │   │   ├── sale_bill_impl
│       │       │   │   │   ├── create_executor.go
│       │       │   │   │   ├── create_many_executor.go
│       │       │   │   │   ├── delete_all_executor.go
│       │       │   │   │   ├── delete_by_id_executor.go
│       │       │   │   │   ├── delete_many_executor.go
│       │       │   │   │   ├── find_all_executor.go
│       │       │   │   │   ├── find_by_id_executor.go
│       │       │   │   │   ├── find_by_ids_executor.go
│       │       │   │   │   ├── find_paging_executor.go
│       │       │   │   │   ├── update_all_executor.go
│       │       │   │   │   ├── update_executor.go
│       │       │   │   │   └── x_init.go
│       │       │   │   ├── sale_item_executor.go
│       │       │   │   └── sale_item_impl
│       │       │   │       ├── create_executor.go
│       │       │   │       ├── create_many_executor.go
│       │       │   │       ├── delete_all_executor.go
│       │       │   │       ├── delete_by_id_executor.go
│       │       │   │       ├── delete_by_sale_bill_id_executor.go
│       │       │   │       ├── delete_many_executor.go
│       │       │   │       ├── find_all_executor.go
│       │       │   │       ├── find_by_id_executor.go
│       │       │   │       ├── find_by_ids_executor.go
│       │       │   │       ├── find_by_sale_bill_id_executor.go
│       │       │   │       ├── find_paging_executor.go
│       │       │   │       ├── update_all_executor.go
│       │       │   │       ├── update_executor.go
│       │       │   │       └── x_init.go
│       │       │   ├── handler
│       │       │   │   ├── sale_bill_query_handler.go
│       │       │   │   └── sale_item_query_handler.go
│       │       │   └── service
│       │       │       ├── sale_bill_query_appservice.go
│       │       │       └── sale_item_query_appservice.go
│       │       └── user
│       │           ├── appquery
│       │           │   └── user_appquery.go
│       │           ├── assembler
│       │           │   └── user_assembler.go
│       │           ├── executor
│       │           │   ├── user_executor.go
│       │           │   └── user_impl
│       │           │       ├── create_executor.go
│       │           │       ├── create_many_executor.go
│       │           │       ├── delete_all_executor.go
│       │           │       ├── delete_by_id_executor.go
│       │           │       ├── delete_many_executor.go
│       │           │       ├── find_all_executor.go
│       │           │       ├── find_by_id_executor.go
│       │           │       ├── find_by_ids_executor.go
│       │           │       ├── find_paging_executor.go
│       │           │       ├── update_all_executor.go
│       │           │       ├── update_executor.go
│       │           │       └── x_init.go
│       │           ├── handler
│       │           │   └── user_query_handler.go
│       │           └── service
│       │               └── user_query_appservice.go
│       ├── domain
│       │   ├── inventory
│       │   │   ├── factory
│       │   │   │   └── inventory_factory.go
│       │   │   ├── query
│       │   │   │   └── inventory_query.go
│       │   │   ├── repository
│       │   │   │   └── inventory_view_repository.go
│       │   │   ├── service
│       │   │   │   └── inventory_query_service.go
│       │   │   └── view
│       │   │       └── inventory_view.go
│       │   ├── sale_bill
│       │   │   ├── factory
│       │   │   │   ├── sale_bill_factory.go
│       │   │   │   └── sale_item_factory.go
│       │   │   ├── query
│       │   │   │   ├── sale_bill_query.go
│       │   │   │   └── sale_item_query.go
│       │   │   ├── repository
│       │   │   │   ├── sale_bill_view_repository.go
│       │   │   │   └── sale_item_view_repository.go
│       │   │   ├── service
│       │   │   │   ├── sale_bill_query_service.go
│       │   │   │   └── sale_item_query_service.go
│       │   │   └── view
│       │   │       ├── sale_bill_statue_enum.go
│       │   │       ├── sale_bill_view.go
│       │   │       └── sale_item_view.go
│       │   └── user
│       │       ├── factory
│       │       │   └── user_factory.go
│       │       ├── query
│       │       │   └── user_query.go
│       │       ├── repository
│       │       │   └── user_view_repository.go
│       │       ├── service
│       │       │   └── user_query_service.go
│       │       └── view
│       │           └── user_view.go
│       ├── infrastructure
│       │   ├── base
│       │   │   ├── application
│       │   │   │   └── handler
│       │   │   │       └── base_query_handler.go
│       │   │   ├── domain
│       │   │   │   └── view
│       │   │   │       └── base_view.go
│       │   │   └── userinterface
│       │   │       └── rest
│       │   │           ├── assembler
│       │   │           │   └── base_assembler.go
│       │   │           ├── dto
│       │   │           │   └── base_dto.go
│       │   │           └── facade
│       │   │               └── base_api.go
│       │   ├── db
│       │   │   ├── dao
│       │   │   │   └── mongo_dao
│       │   │   │       └── dao.go
│       │   │   └── session
│       │   │       └── session.go
│       │   ├── domain_impl
│       │   │   ├── inventory
│       │   │   │   ├── repository_impl
│       │   │   │   │   └── mongodb
│       │   │   │   │       └── inventory_view_repository_impl.go
│       │   │   │   └── service_impl
│       │   │   │       ├── inventory_query_service_impl.go
│       │   │   │       └── x_options.go
│       │   │   ├── sale_bill
│       │   │   │   ├── repository_impl
│       │   │   │   │   └── mongodb
│       │   │   │   │       ├── sale_bill_view_repository_impl.go
│       │   │   │   │       └── sale_item_view_repository_impl.go
│       │   │   │   └── service_impl
│       │   │   │       ├── sale_bill_query_service_impl.go
│       │   │   │       ├── sale_item_query_service_impl.go
│       │   │   │       └── x_options.go
│       │   │   └── user
│       │   │       ├── repository_impl
│       │   │       │   └── mongodb
│       │   │       │       └── user_view_repository_impl.go
│       │   │       └── service_impl
│       │   │           ├── user_query_service_impl.go
│       │   │           └── x_options.go
│       │   ├── register
│       │   │   ├── register_aggregate_type.go
│       │   │   ├── register_event_type.go
│       │   │   ├── register_rest_api.go
│       │   │   └── register_subscribe.go
│       │   ├── types
│       │   │   └── types.go
│       │   └── utils
│       │       └── utils.go
│       └── userinterface
│           └── rest
│               ├── inventory
│               │   ├── assembler
│               │   │   └── inventory_assembler.go
│               │   ├── dto
│               │   │   └── inventory_dto.go
│               │   └── facade
│               │       └── inventory_api.go
│               ├── sale_bill
│               │   ├── assembler
│               │   │   ├── sale_bill_assembler.go
│               │   │   └── sale_item_assembler.go
│               │   ├── dto
│               │   │   ├── sale_bill_dto.go
│               │   │   └── sale_item_dto.go
│               │   └── facade
│               │       ├── sale_bill_api.go
│               │       └── sale_item_api.go
│               ├── swagger.go
│               └── user
│                   ├── assembler
│                   │   └── user_assembler.go
│                   ├── dto
│                   │   └── user_dto.go
│                   └── facade
│                       └── user_api.go
├── swagger
│   ├── cmd
│   │   ├── docs.go
│   │   ├── swagger.json
│   │   └── swagger.yaml
│   └── query
│       ├── docs.go
│       ├── swagger.json
│       └── swagger.yaml
├── test
├── tree.,md
└── tree.md

176 directories, 284 files
