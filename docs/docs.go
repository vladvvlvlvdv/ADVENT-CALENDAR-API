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
        "/api/days": {
            "get": {
                "tags": [
                    "Days"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Например Europe/Samara",
                        "name": "timeZone",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/repository.Day"
                            }
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Days"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "minLength": 5,
                        "type": "string",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 5,
                        "type": "string",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "csv",
                        "description": " ",
                        "name": "attachments",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    }
                }
            }
        },
        "/api/days/admin": {
            "get": {
                "tags": [
                    "Days"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/repository.Day"
                            }
                        }
                    }
                }
            }
        },
        "/api/days/{id}": {
            "put": {
                "tags": [
                    "Days"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "name": "attachmentIds",
                        "in": "formData"
                    },
                    {
                        "minLength": 5,
                        "type": "string",
                        "name": "description",
                        "in": "formData"
                    },
                    {
                        "minLength": 5,
                        "type": "string",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "csv",
                        "description": " ",
                        "name": "attachments",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    }
                }
            }
        },
        "/api/settings": {
            "get": {
                "tags": [
                    "Settings"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/repository.Setting"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "Settings"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "maximum": 12,
                        "minimum": 1,
                        "type": "integer",
                        "name": "month",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "showAllDays",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    }
                }
            }
        },
        "/api/users/check": {
            "get": {
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    }
                }
            }
        },
        "/api/users/confirm": {
            "patch": {
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 5,
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Tokens"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    }
                }
            }
        },
        "/api/users/login": {
            "post": {
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "minLength": 5,
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 5,
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Tokens"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    }
                }
            }
        },
        "/api/users/refresh": {
            "patch": {
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "RefreshToken",
                        "name": "RefreshToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Tokens"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    }
                }
            }
        },
        "/api/users/register": {
            "post": {
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "Email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/validators.GlobalHandlerResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Tokens": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "exp": {
                    "type": "integer"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "repository.Attachment": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "label": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "repository.Day": {
            "type": "object",
            "properties": {
                "attachments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/repository.Attachment"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isLongRead": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "repository.Setting": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "month": {
                    "type": "integer"
                },
                "showAllDays": {
                    "type": "boolean"
                }
            }
        },
        "repository.User": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "refreshToken": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "validators.GlobalHandlerResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "172.23.116.163:9000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Advent Calendar API docs",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
