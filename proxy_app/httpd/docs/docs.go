// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/trustmeshes": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "get all trustmeshes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trustmeshes"
                ],
                "summary": "Get all trustmeshes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.trustmeshDto"
                            }
                        }
                    }
                }
            }
        },
        "/trustmeshes/{id}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "get single trustmesh",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trustmesh"
                ],
                "summary": "Get single trustmesh",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.trustmeshDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.trustmeshDto": {
            "type": "object",
            "properties": {
                "businessObjectTypes": {
                    "type": "string"
                },
                "containsRejections": {
                    "type": "boolean"
                },
                "createdAt": {
                    "type": "string"
                },
                "endTime": {
                    "type": "string"
                },
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.trustmeshEntryDto"
                    }
                },
                "finalized": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "participants": {
                    "type": "string"
                },
                "startTime": {
                    "type": "string"
                }
            }
        },
        "handler.trustmeshEntryDto": {
            "type": "object",
            "properties": {
                "baseledgerBusinessObjectId": {
                    "type": "string"
                },
                "baseledgerTransactionId": {
                    "type": "string"
                },
                "baseledgerTransactionType": {
                    "type": "string"
                },
                "businessObjectType": {
                    "type": "string"
                },
                "commitmentState": {
                    "type": "string"
                },
                "entryType": {
                    "type": "string"
                },
                "offchainProcessMessageId": {
                    "type": "string"
                },
                "receiverOrgId": {
                    "type": "string"
                },
                "receiverOrgName": {
                    "type": "string"
                },
                "referencedBaseledgerBusinessObjectId": {
                    "type": "string"
                },
                "referencedBaseledgerTransactionId": {
                    "type": "string"
                },
                "referencedProcessMessageId": {
                    "type": "string"
                },
                "senderOrgId": {
                    "type": "string"
                },
                "senderOrgName": {
                    "type": "string"
                },
                "tendermintBlockId": {
                    "type": "string"
                },
                "tendermintTransactionId": {
                    "type": "string"
                },
                "tendermintTransactionTimestamp": {
                    "type": "string"
                },
                "transactionHash": {
                    "type": "string"
                },
                "trustmeshId": {
                    "type": "string"
                },
                "workgroupId": {
                    "type": "string"
                },
                "workgroupName": {
                    "type": "string"
                },
                "workstepType": {
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

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0.0",
	Host:        "localhost:8081",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Baseledger Proxy API documentation",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
