// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "@yusufocaliskan",
            "email": "yusufocaliskan@gmail.com"
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
        "/api/v1/auth/login": {
            "post": {
                "description": "Sing-in With Access Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "operationId": "access-token-login",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authmodel.AuthRefreshTokenModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Language preference",
                        "name": "Accept-Language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usermodel.UserWithToken"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/logout": {
            "post": {
                "description": "Sing out",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout",
                "operationId": "sing-out",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authmodel.AuthRefreshTokenModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Language preference",
                        "name": "Accept-Language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/translator.TranslationSwaggerResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/refreshToken": {
            "post": {
                "description": "Generating new accessToken using refreshToken",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh Token",
                "operationId": "refresh-token",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authmodel.AuthRefreshTokenModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Language preference",
                        "name": "Accept-Language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usermodel.UserWithToken"
                        }
                    }
                }
            }
        },
        "/api/v1/user/createByEmail": {
            "post": {
                "description": "Creates new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "New user",
                "operationId": "create-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "query params",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Language preference",
                        "name": "Accept-Language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usermodel.UserWithToken"
                        }
                    }
                }
            }
        },
        "/api/v1/user/deleteById": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deletes a user by given user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user",
                "operationId": "Delete-User",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usermodel.UserDeleteModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Language preference",
                        "name": "Accept-Language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/translator.TranslationSwaggerResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/getUserById": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get user details by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get User",
                "operationId": "get-user-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "format": "ObjectID",
                        "description": "user id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Language preference",
                        "name": "Accept-Language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usermodel.UserWithoutPasswordModel"
                        }
                    }
                }
            }
        },
        "/api/v1/user/updateUserInformationsById": {
            "put": {
                "description": "Updates user informations by giving the Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update User",
                "operationId": "update-user-informations",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usermodel.UserUpdateModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Language preference",
                        "name": "Accept-Language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usermodel.UserSwaggerParams"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authmodel.AuthRefreshTokenModel": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string",
                    "example": "the_refresh_token"
                }
            }
        },
        "time.Duration": {
            "type": "integer",
            "enum": [
                -9223372036854775808,
                9223372036854775807,
                1,
                1000,
                1000000,
                1000000000,
                60000000000,
                3600000000000
            ],
            "x-enum-varnames": [
                "minDuration",
                "maxDuration",
                "Nanosecond",
                "Microsecond",
                "Millisecond",
                "Second",
                "Minute",
                "Hour"
            ]
        },
        "tinytoken.SingleToken": {
            "type": "object",
            "properties": {
                "expiry_time": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/time.Duration"
                        }
                    ],
                    "example": 86400000000000
                },
                "key": {
                    "type": "string"
                }
            }
        },
        "tinytoken.TinyTokenData": {
            "type": "object",
            "properties": {
                "access_token": {
                    "$ref": "#/definitions/tinytoken.SingleToken"
                },
                "refresh_token": {
                    "$ref": "#/definitions/tinytoken.SingleToken"
                }
            }
        },
        "translator.TranslationEntry": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "translator.TranslationSwaggerResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "$ref": "#/definitions/translator.TranslationEntry"
                }
            }
        },
        "usermodel.UserDeleteModel": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "usermodel.UserModel": {
            "type": "object",
            "required": [
                "email",
                "password",
                "role",
                "username"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "hashed_password": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "admin",
                        "moderator",
                        "user"
                    ]
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "usermodel.UserSwaggerParams": {
            "type": "object",
            "required": [
                "email",
                "password",
                "role"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "fullname": {
                    "type": "string",
                    "example": "John Doe"
                },
                "name": {
                    "type": "string",
                    "example": "johndoe"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "admin",
                        "moderator",
                        "user"
                    ]
                },
                "username": {
                    "type": "string",
                    "example": "johndoe"
                }
            }
        },
        "usermodel.UserUpdateModel": {
            "type": "object",
            "required": [
                "email",
                "role",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "admin",
                        "moderator",
                        "user"
                    ]
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "usermodel.UserWithToken": {
            "type": "object",
            "properties": {
                "tokens": {
                    "$ref": "#/definitions/tinytoken.TinyTokenData"
                },
                "user": {
                    "$ref": "#/definitions/usermodel.UserModel"
                }
            }
        },
        "usermodel.UserWithoutPasswordModel": {
            "type": "object",
            "required": [
                "email",
                "role",
                "username"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "admin",
                        "moderator",
                        "user"
                    ]
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "GPTVerse Admin Backend",
	Description:      "To manage the whole gptv.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
