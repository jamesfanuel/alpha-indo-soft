package helper

import (
	"bytes"

	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"

	"net/http"

	"net/http/httptest"
)

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/articles", RootEndpoint).Methods("GET")
	router.HandleFunc("/articles/{author}", RootEndpoint).Methods("GET")
	router.HandleFunc("/articles/query?title={title}&body={body}", RootEndpoint).Methods("GET")
	router.HandleFunc("/articles", RootEndpoint).Methods("POST")

	return router
}

func TestFindArticles(t *testing.T) {
	request, _ := http.NewRequest("GET", "/articles", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok Response")
}

func TestFindArticlesByAuthor(t *testing.T) {
	request, _ := http.NewRequest("GET", "/articles/abc", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok Response")
}

func TestFindArticlesByQuery(t *testing.T) {
	request, _ := http.NewRequest("GET", "/articles/query?title=a&body=b", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok Response")
}

func TestCreateArticle(t *testing.T) {
	testBody := `{
		"title": "TEST",
		"author": "TEST",
		"body": "TEST"
	  }`
	request, _ := http.NewRequest("POST", "/articles", bytes.NewBufferString(testBody))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok Response")
}