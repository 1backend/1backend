// Package spec Code generated by swaggo/swag. DO NOT EDIT
package spec

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
            "url": "https://1backend.com/",
            "email": "sales@singulatron.com"
        },
        "license": {
            "name": "Proprietary"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/basic-svc/pet": {
            "put": {
                "description": "Save a pet for a user and an organization.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basic Svc"
                ],
                "summary": "Save Pet",
                "operationId": "savePet",
                "parameters": [
                    {
                        "description": "Registration Tracking Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/basic_svc.SavePetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{}",
                        "schema": {
                            "$ref": "#/definitions/basic_svc.SavePetResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/basic-svc/pets": {
            "post": {
                "description": "List pets.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basic Svc"
                ],
                "summary": "List Pets",
                "operationId": "listPets",
                "parameters": [
                    {
                        "description": "List Pets Request",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/basic_svc.ListPetsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{}",
                        "schema": {
                            "$ref": "#/definitions/basic_svc.ListPetsResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error Listing Pets",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "basic_svc.ListPetsRequest": {
            "type": "object"
        },
        "basic_svc.ListPetsResponse": {
            "type": "object",
            "properties": {
                "pets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/basic_svc.Pet"
                    }
                }
            }
        },
        "basic_svc.Pet": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "basic_svc.SavePetRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "basic_svc.SavePetResponse": {
            "type": "object"
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and token acquired from the User Svc Login endpoint.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.3.0-rc.8",
	Host:             "localhost:58231",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Basic Svc",
	Description:      "An example service for bootstrapping.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
