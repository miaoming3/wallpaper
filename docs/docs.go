// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/login": {
            "post": {
                "description": "管理员用户登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员管理"
                ],
                "summary": "管理员登录",
                "parameters": [
                    {
                        "description": "管理员登录字段",
                        "name": "AdminLoginData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AdminLoginData"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/category/del": {
            "delete": {
                "description": "删除分类",
                "tags": [
                    "分类管理"
                ],
                "summary": "删除分类",
                "parameters": [
                    {
                        "description": "Category data to delete",
                        "name": "DeleteCategory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteCategory"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/category/index": {
            "get": {
                "description": "获取分类列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类管理"
                ],
                "summary": "分类列表",
                "responses": {}
            }
        },
        "/category/save": {
            "post": {
                "description": "创建分类",
                "tags": [
                    "分类管理"
                ],
                "summary": "创建分类",
                "parameters": [
                    {
                        "description": "Category data to save",
                        "name": "saveCategory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SaveCategory"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/category/update": {
            "put": {
                "description": "修改分类",
                "tags": [
                    "分类管理"
                ],
                "summary": "修改分类",
                "parameters": [
                    {
                        "description": "Category data to update",
                        "name": "UpdateCategory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/image/index": {
            "get": {
                "description": "获取图片列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "图片列表",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "cid",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "is_recommend",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "tags_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/menu/del": {
            "delete": {
                "description": "删除菜单",
                "tags": [
                    "菜单管理"
                ],
                "summary": "删除菜单",
                "parameters": [
                    {
                        "description": "删除菜单参数",
                        "name": "DeleteMenu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteMenu"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/menu/index": {
            "get": {
                "description": "获取菜单列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "菜单列表",
                "responses": {}
            }
        },
        "/menu/save": {
            "post": {
                "description": "创建菜单",
                "tags": [
                    "菜单管理"
                ],
                "summary": "创建菜单",
                "parameters": [
                    {
                        "description": "创建菜单",
                        "name": "SaveMenu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SaveMenu"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/menu/update": {
            "put": {
                "description": "修改菜单",
                "tags": [
                    "菜单管理"
                ],
                "summary": "修改菜单",
                "parameters": [
                    {
                        "description": "修改菜单",
                        "name": "UpdateMenu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AdminLoginData": {
            "type": "object",
            "required": [
                "captcha_id",
                "password",
                "username"
            ],
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "captcha_id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.DeleteCategory": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "dto.DeleteMenu": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "dto.SaveCategory": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "pid": {
                    "type": "integer",
                    "default": 0
                },
                "sort": {
                    "type": "integer",
                    "default": 50
                },
                "status": {
                    "default": 1,
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.CategoryStatus"
                        }
                    ]
                }
            }
        },
        "dto.SaveMenu": {
            "type": "object"
        },
        "dto.UpdateCategory": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "pid": {
                    "type": "integer",
                    "default": 0
                },
                "sort": {
                    "type": "integer",
                    "default": 50
                },
                "status": {
                    "default": 1,
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.CategoryStatus"
                        }
                    ]
                }
            }
        },
        "dto.UpdateMenu": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.CategoryStatus": {
            "type": "integer",
            "enum": [
                1,
                2
            ],
            "x-enum-varnames": [
                "StatusOn",
                "StatusOff"
            ]
        },
        "response.ApiResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "error": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "page": {
                    "$ref": "#/definitions/response.PageData"
                }
            }
        },
        "response.PageData": {
            "type": "object",
            "properties": {
                "next": {
                    "type": "boolean"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
