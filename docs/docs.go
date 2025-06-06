// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import (
	_ "embed"
	"encoding/json"

	"github.com/swaggo/swag"
)

//go:embed swagger_spec.json
var specJSON []byte

var doc string

func init() {
	// 讀取並解析 swagger_spec.json
	var spec map[string]interface{}
	if err := json.Unmarshal(specJSON, &spec); err != nil {
		panic(err)
	}

	// 轉換為字符串
	jsonBytes, err := json.MarshalIndent(spec, "", "    ")
	if err != nil {
		panic(err)
	}
	doc = string(jsonBytes)

	swaggerInfo := &swag.Spec{
		Version:          "1.0",
		Host:            "localhost:8080",
		BasePath:        "/api/v1",
		Schemes:         []string{"http"},
		Title:           "Go API with Gin and MongoDB",
		Description:     "這是一個使用 Gin 和 MongoDB 的 RESTful API 服務",
		InfoInstanceName: "swagger",
		SwaggerTemplate: doc,
		LeftDelim:       "{{",
		RightDelim:      "}}",
	}

	swag.Register(swaggerInfo.InstanceName(), swaggerInfo)
}
