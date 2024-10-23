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
        "/cpemployee/cpprofile/primresource": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cpprofile-controller"
                ],
                "summary": "affiliateorchestration",
                "parameters": [
                    {
                        "type": "string",
                        "description": "X-Channel",
                        "name": "X-Channel",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-GatewayType",
                        "name": "X-GatewayType",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-LegacyUsername",
                        "name": "X-LegacyUsername",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-Username",
                        "name": "X-Username",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "requestInfo",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.RequestCpProfile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/apimodel.ResponseCpProfile"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimodel.ResponseCpProfile"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apimodel.ResponseCpProfile"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apimodel.ResponseCpProfile"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/apimodel.ResponseCpProfile"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apimodel.BackendResponseInfoArray": {
            "type": "object",
            "properties": {
                "apiName": {
                    "type": "string"
                },
                "errorCode": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "system": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "apimodel.BackendResponseList": {
            "type": "object",
            "properties": {
                "backendResponseInfoArray": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimodel.BackendResponseInfoArray"
                    }
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "apimodel.CpProfile": {
            "type": "object",
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "certificateNumber": {
                    "type": "string"
                },
                "companyGroup": {
                    "type": "string"
                },
                "companyName": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "employeeId": {
                    "type": "string"
                },
                "engName": {
                    "$ref": "#/definitions/apimodel.NameInfo"
                },
                "payroll": {
                    "type": "string"
                },
                "thaiName": {
                    "$ref": "#/definitions/apimodel.NameInfo"
                }
            }
        },
        "apimodel.NameInfo": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "apimodel.RequestCpProfile": {
            "type": "object",
            "required": [
                "businessLine",
                "correlatedId",
                "primResourceValue"
            ],
            "properties": {
                "businessLine": {
                    "type": "string"
                },
                "correlatedId": {
                    "type": "string"
                },
                "primResourceValue": {
                    "type": "string"
                }
            }
        },
        "apimodel.ResponseCpProfile": {
            "type": "object",
            "properties": {
                "backendResponseList": {
                    "$ref": "#/definitions/apimodel.BackendResponseList"
                },
                "cpProfile": {
                    "$ref": "#/definitions/apimodel.CpProfile"
                },
                "errorCode": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "uuid": {
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
	Version:     "00.00",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{"http"},
	Title:       "[INTX] affiliateorchestration Swagger APIs",
	Description: "[INTX] affiliateorchestration",
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
