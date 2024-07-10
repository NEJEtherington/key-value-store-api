package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"	
	"github.com/stretchr/testify/assert"
	"kvp-api/internal/db"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	// start the server
	kvdb := db.NewKeyValueDB(map[string]string{"b":"B", "c":"C"})
	router = InitRoutes(kvdb)
	

	// run the tests
	code := m.Run()

	os.Exit(code)
}

func TestGetKeys(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `["b","c"]`, w.Body.String())
}

func TestGetValue(t *testing.T) {
	notFound := httptest.NewRecorder()
	failreq, _ := http.NewRequest("GET", "/a", nil)
	router.ServeHTTP(notFound, failreq)

	assert.Equal(t, 404, notFound.Code)
	assert.Equal(t, `{"error":"Key does not exist"}`, notFound.Body.String())

	ok := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/b", nil)
	router.ServeHTTP(ok, req)

	assert.Equal(t, 200, ok.Code)
	assert.Equal(t, "B", ok.Body.String())
}

func TestUpdateValue(t *testing.T) {
	notFoundBody := PutRequestBody{
		NewValue: "a",
	}
	notFoundJSONBody, _ := json.Marshal(notFoundBody)
	notFound := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/a", strings.NewReader(string(notFoundJSONBody)))
	router.ServeHTTP(notFound, req)

	assert.Equal(t, 404, notFound.Code)
	assert.Equal(t, `{"error":"Key does not exist"}`, notFound.Body.String())

	okBody := PutRequestBody{
		NewValue: "b",
	}
	okJSONBody, _ := json.Marshal(okBody)
	ok := httptest.NewRecorder()
	req1, _ := http.NewRequest("PUT", "/b", strings.NewReader(string(okJSONBody)))
	router.ServeHTTP(ok, req1)

	assert.Equal(t, 200, ok.Code)
	assert.Equal(t, `{"b":"b"}`, ok.Body.String())
}

func TestDeleteValue(t *testing.T) {
	notFound := httptest.NewRecorder()
	notFoundReq, _ := http.NewRequest("DELETE", "/a", nil)
	router.ServeHTTP(notFound, notFoundReq)

	assert.Equal(t, 404, notFound.Code)
	assert.Equal(t, `{"error":"Key does not exist"}`, notFound.Body.String())

	ok := httptest.NewRecorder()
	okReq, _ := http.NewRequest("DELETE", "/c", nil)
	router.ServeHTTP(ok, okReq)

	assert.Equal(t, 200, ok.Code)
	assert.Equal(t, `"c"`, ok.Body.String())
}

