// Package query GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package query

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tenants/{tenantId}/users": {
            "get": {
                "description": "分页查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "分页查询",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tenant ID",
                        "name": "tenantId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserFindPagingResponse"
                        }
                    },
                    "500": {
                        "description": "应用错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tenants/{tenantId}/users/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "按ID获取用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tenant ID",
                        "name": "tenantId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserFindByIdResponse"
                        }
                    },
                    "404": {
                        "description": "按ID找到数据",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "应用错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tenants/{tenantId}/users:all": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "获取所有用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tenant ID",
                        "name": "tenantId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.UserFindAllResponseItem"
                            }
                        }
                    },
                    "500": {
                        "description": "应用错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.UserFindAllResponseItem": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "telephone": {
                    "type": "string"
                },
                "tenantId": {
                    "type": "string"
                },
                "userCode": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "dto.UserFindByIdResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "telephone": {
                    "type": "string"
                },
                "tenantId": {
                    "type": "string"
                },
                "userCode": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "dto.UserFindPagingResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.UserFindPagingResponseItem"
                    }
                },
                "filter": {
                    "type": "string"
                },
                "pageNum": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "sort": {
                    "type": "string"
                },
                "totalPages": {
                    "type": "integer"
                },
                "totalRows": {
                    "type": "integer"
                }
            }
        },
        "dto.UserFindPagingResponseItem": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "telephone": {
                    "type": "string"
                },
                "tenantId": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "string"
                },
                "userCode": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9020",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "go-ddd-example query-service API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
