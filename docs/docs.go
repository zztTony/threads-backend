// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "发表意见平台技术条款",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/articles": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取多个文章",
                "parameters": [
                    {
                        "maxLength": 30,
                        "type": "string",
                        "description": "文章标题",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "文章状态",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.ArticleSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Err"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Err"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "创建文章",
                "parameters": [
                    {
                        "maxLength": 30,
                        "minLength": 1,
                        "description": "文章标题",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 50,
                        "description": "描述",
                        "name": "desc",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 5000,
                        "minLength": 10,
                        "description": "内容",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "封面图片地址",
                        "name": "cover_image_urlbody",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 10,
                        "description": "发表人",
                        "name": "created_by",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.ArticleSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Err"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Err"
                        }
                    }
                }
            }
        },
        "/api/v1/articles/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取单个文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.ArticleSwagger"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标签ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/articles{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章ID",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "maxLength": 30,
                        "minLength": 1,
                        "description": "文章标题",
                        "name": "title",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 50,
                        "description": "描述",
                        "name": "desc",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 5000,
                        "minLength": 10,
                        "description": "正文",
                        "name": "content",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "封面图片地址",
                        "name": "cover_image_urlbody",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 10,
                        "description": "发表人",
                        "name": "modified_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.ArticleSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Err"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Err"
                        }
                    }
                }
            }
        },
        "/api/v1/tags": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取多个标签",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "标签名称",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "type": "boject"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Err"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增标签",
                "parameters": [
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "标签名称",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "创建者",
                        "name": "created_by",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Err"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Err"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标签ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "app.Paper": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total_rows": {
                    "type": "integer"
                }
            }
        },
        "errcode.Err": {
            "type": "object"
        },
        "model.Article": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "cover_image_url": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted": {
                    "type": "integer"
                },
                "deleted_on": {
                    "type": "integer"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "modified_on": {
                    "type": "integer"
                },
                "state": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ArticleSwagger": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Article"
                    }
                },
                "pager": {
                    "$ref": "#/definitions/app.Paper"
                }
            }
        },
        "model.Tag": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted": {
                    "type": "integer"
                },
                "deleted_on": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "modified_on": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                }
            }
        },
        "model.TagSwagger": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Tag"
                    }
                },
                "pager": {
                    "$ref": "#/definitions/app.Paper"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "战舰世界论坛",
	Description:      "战舰世界论坛, 畅言好玩战舰",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
