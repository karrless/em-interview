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
        "/ping": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Healthcheck"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get songs",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "Song group",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "Song title",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Song release date",
                        "name": "release_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Song release date before",
                        "name": "before",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Song release date after",
                        "name": "after",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Song"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Create song",
                "parameters": [
                    {
                        "description": "Song request",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateSongRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created song",
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {}
                    }
                }
            }
        },
        "/songs/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get song by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "Update song by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Song",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Delete song by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CreateSongRequest": {
            "type": "object",
            "properties": {
                "group": {
                    "description": "Название группы",
                    "type": "string"
                },
                "song": {
                    "description": "Название песни",
                    "type": "string"
                }
            }
        },
        "models.Song": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
