package main

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestGinHelloWorld(t *testing.T) {
	r := gofight.New()

	r.GET("/welcome?firstname=apple&lastname=boy").
		SetDebug(true).
		Run(ginEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "Hello apple boy", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestPing(t *testing.T) {
	r := gofight.New()

	r.GET("/ping").
		SetDebug(true).
		Run(ginEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"message\":{\"data1\":\"data2\"}}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
