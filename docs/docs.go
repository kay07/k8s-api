// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/alldep": {
            "get": {
                "description": "详情",
                "tags": [
                    "k8s-deployment"
                ],
                "summary": "获取所有deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/svc": {
            "post": {
                "description": "详情",
                "tags": [
                    "k8s-svc"
                ],
                "summary": "创建svc",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "port",
                        "name": "port",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "replicas",
                        "name": "replicas",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "image",
                        "name": "image",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "env",
                        "name": "env",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "volumes",
                        "name": "volumes",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/svctest": {
            "post": {
                "description": "详情",
                "tags": [
                    "k8s-svc-test"
                ],
                "summary": "创建svctest",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "env可以为空",
                        "name": "ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "kubernetes API",
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
