package main

import (
	"encoding/json"
	"fmt"
	"github.com/AlperKocaman/server-with-aws/cmd/config"
	"github.com/AlperKocaman/server-with-aws/core/app"
	"github.com/AlperKocaman/server-with-aws/core/aws"
	"github.com/AlperKocaman/server-with-aws/core/server"
	"github.com/appleboy/gofight/v2"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"testing"
)

func TestController_SaveObject(t *testing.T) {

	tests := []struct {
		name         string
		code         int
		requestBody  app.SaveObjectParam
		responseData interface{}
	}{
		{
			name: "Server should return 200 and saves given object ",
			code: http.StatusOK,
			requestBody: app.SaveObjectParam{
				Key:  "test",
				Data: "This is the test data.",
			},
			responseData: app.SaveObjectResponse{
				Key: "test",
			},
		},
	}

	err := config.InitializeConfigForTest()
	if err != nil {
		log.Fatal("error while reading config, exiting.")
	}

	handler := server.InitializeServer()
	var response gofight.HTTPResponse

	for _, tt := range tests {
		var body gjson.Result

		t.Run(tt.name, func(t *testing.T) {
			gofight.New().
				GET("/picus/list").
				Run(handler, func(r gofight.HTTPResponse, request gofight.HTTPRequest) {
					response = r
				})

			Convey("When client sends a request to get object", t, func() {
				Convey("Then server should return expected code", func() {
					So(response.Code, ShouldEqual, tt.code)
					body = gjson.GetBytes(response.Body.Bytes(), "data")
				})

				if tt.code == http.StatusOK {
					Convey("Then server should return object", func() {
						expectedResponseBytes, err := json.Marshal(tt.responseData)
						So(err, ShouldBeNil)

						So(body.String(), ShouldEqual, string(expectedResponseBytes))
					})
				}
			})
		})
	}
}

func Test_GetObject(t *testing.T) {

	tests := []struct {
		name         string
		key          string
		code         int
		responseData interface{}
	}{
		{
			name:         "Server should return 400 and return error when no key is given",
			key:          "",
			code:         http.StatusBadRequest,
			responseData: app.GetObjectResponse{},
		},
		{
			name: "Server should return 200 and return object when key is test",
			key:  "test",
			code: http.StatusOK,
			responseData: app.GetObjectResponse{
				Content: aws.Content{
					Data:          "This is the test data.",
					ContentLength: 22,
					ContentType:   "binary/octet-stream",
				},
			},
		},
		{
			name:         "Server should return 404 and return nothing when key is not_existing_key",
			key:          "not_existing_key",
			code:         http.StatusNotFound,
			responseData: app.GetObjectResponse{},
		},
	}

	err := config.InitializeConfigForTest()
	if err != nil {
		log.Fatal("error while reading config, exiting.")
	}

	handler := server.InitializeServer()
	var response gofight.HTTPResponse

	for _, tt := range tests {
		var body gjson.Result

		t.Run(tt.name, func(t *testing.T) {
			gofight.New().
				GET(fmt.Sprintf("/picus/get?key=%s", tt.key)).
				Run(handler, func(r gofight.HTTPResponse, request gofight.HTTPRequest) {
					response = r
					body = gjson.GetBytes(response.Body.Bytes(), "data")
				})

			Convey("When client sends a request to get object", t, func() {
				Convey("Then server should return expected code", func() {
					So(response.Code, ShouldEqual, tt.code)
				})

				if tt.code == http.StatusOK {
					Convey("Then server should return object", func() {
						expectedResponseBytes, err := json.Marshal(tt.responseData)
						So(err, ShouldBeNil)

						So(body.String(), ShouldEqual, string(expectedResponseBytes))
					})
				}
			})
		})
	}
}

func TestController_ListObjects(t *testing.T) {

	tests := []struct {
		name string
		code int
	}{
		{
			name: "Server should return 200 and return list of objects",
			code: http.StatusOK,
		},
	}

	err := config.InitializeConfigForTest()
	if err != nil {
		log.Fatal("error while reading config, exiting.")
	}

	handler := server.InitializeServer()
	var response gofight.HTTPResponse

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			gofight.New().
				GET("/picus/list").
				Run(handler, func(r gofight.HTTPResponse, request gofight.HTTPRequest) {
					response = r
				})

			Convey("When client sends a request to get object", t, func() {
				Convey("Then server should return expected code", func() {
					So(response.Code, ShouldEqual, tt.code)
				})
			})
		})
	}

}
