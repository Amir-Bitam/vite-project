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
            "name": "API Support",
            "email": "support@farmmanagement.io"
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
        "/auth/login": {
            "post": {
                "description": "Login with email and password to get JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "$ref": "#/definitions/main.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a new user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/farms": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create a new farm for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "farms"
                ],
                "summary": "Create a new farm",
                "parameters": [
                    {
                        "description": "Farm object",
                        "name": "farm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Farm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Farm created successfully",
                        "schema": {
                            "$ref": "#/definitions/main.Farm"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/farms/{farm_id}/crops": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get all crops for a specific farm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "crops"
                ],
                "summary": "Get farm crops",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Farm ID",
                        "name": "farm_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of crops retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Crop"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid farm ID",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Farm not found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Add a new crop to a specific farm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "crops"
                ],
                "summary": "Add a crop to a farm",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Farm ID",
                        "name": "farm_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Crop object",
                        "name": "crop",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Crop"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Crop added successfully",
                        "schema": {
                            "$ref": "#/definitions/main.Crop"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Farm not found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/farms/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get detailed information about a specific farm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "farms"
                ],
                "summary": "Get farm details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Farm ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Farm details retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/main.Farm"
                        }
                    },
                    "404": {
                        "description": "Farm not found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update information about a specific farm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "farms"
                ],
                "summary": "Update farm details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Farm ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Farm object",
                        "name": "farm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Farm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Farm updated successfully",
                        "schema": {
                            "$ref": "#/definitions/main.Farm"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Farm not found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a specific farm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "farms"
                ],
                "summary": "Delete a farm",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Farm ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Farm deleted successfully"
                    },
                    "404": {
                        "description": "Farm not found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get user details including their farms",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get a user by ID",
                "parameters": [
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
                        "description": "User details retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}/farms": {
            "get": {
                "description": "Get all farms belonging to a specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "farms"
                ],
                "summary": "Get user's farms",
                "parameters": [
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Farm"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Crop": {
            "type": "object",
            "properties": {
                "area": {
                    "type": "number",
                    "example": 50.5
                },
                "created_at": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                },
                "deleted_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "farm_id": {
                    "type": "integer",
                    "example": 1
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Corn"
                },
                "plant_date": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                },
                "yield_rate": {
                    "type": "number",
                    "example": 8.5
                }
            }
        },
        "main.ErrorResponse": {
            "description": "Generic error response",
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Error message details"
                }
            }
        },
        "main.Farm": {
            "description": "Farm information",
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                },
                "crops": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Crop"
                    }
                },
                "deleted_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "last_harvest": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                },
                "location": {
                    "type": "string",
                    "example": "Kansas"
                },
                "name": {
                    "type": "string",
                    "example": "Green Acres"
                },
                "size": {
                    "type": "number",
                    "example": 100.5
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "active",
                        "inactive"
                    ],
                    "example": "active"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "main.LoginRequest": {
            "description": "Login request payload",
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "secretpassword"
                }
            }
        },
        "main.LoginResponse": {
            "description": "Login response with JWT token",
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                },
                "user": {
                    "$ref": "#/definitions/main.User"
                }
            }
        },
        "main.User": {
            "description": "User account information",
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                },
                "deleted_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "password": {
                    "type": "string",
                    "example": "secretpassword"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Farm Management API",
	Description:      "Farm information",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
