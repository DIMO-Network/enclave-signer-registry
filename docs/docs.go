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
        "/": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/Developer-license": {
            "get": {
                "description": "Get the developer license of the controller",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "keys"
                ],
                "summary": "Get developer license",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_app.DeveloperLicenseResponse"
                        }
                    }
                }
            }
        },
        "/add-signer": {
            "post": {
                "description": "Add a new signer by verifying their NSM attestation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "signer"
                ],
                "summary": "Add a new signer",
                "parameters": [
                    {
                        "description": "NSM attestation document",
                        "name": "attestation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_app.AddSignerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_app.AddSignerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_app.codeResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_app.codeResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_DIMO-Network_enclave-signer-registry_internal_config.PCRValues": {
            "type": "object",
            "properties": {
                "pcr0": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "pcr1": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "pcr2": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "internal_app.AddSignerRequest": {
            "type": "object",
            "properties": {
                "attestationDocument": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "signerAddress": {
                    "type": "string"
                }
            }
        },
        "internal_app.AddSignerResponse": {
            "type": "object",
            "properties": {
                "alreadyAdded": {
                    "type": "boolean"
                },
                "txHash": {
                    "type": "string"
                }
            }
        },
        "internal_app.DeveloperLicenseResponse": {
            "type": "object",
            "properties": {
                "clientId": {
                    "type": "string"
                },
                "tokenId": {
                    "type": "string"
                },
                "validPCRs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_DIMO-Network_enclave-signer-registry_internal_config.PCRValues"
                    }
                }
            }
        },
        "internal_app.codeResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Enclave Signer Registry",
	Description:      "This is the API documentation for the Enclave Signer Registry service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
