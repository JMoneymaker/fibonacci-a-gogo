package main

import (
	"testing"
	"net/http"
	"httptest" //wasn't able to install this package
	"github.com/julienschmidt/httprouter"
)

func TestGetNumbers(t *testing.T) {
    router := httprouter.New()
    router.GET("/api/fibonacci/:number", GetNumbers)

    req, _ := http.NewRequest("GET", "/api/fibonacci/25", nil)
    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)
    //add expectations here
}