package webpzngo

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query().Get("name")
	if query == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", query)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Luthfi", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	assert.Equal(t, "Hello Luthfi", bodyString, "Result must be same as the query data")
}

func SayHelloMultiple(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Luthfi&last_name=Izzuddin", nil)
	recorder := httptest.NewRecorder()

	SayHelloMultiple(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	assert.Equal(t, "Hello Luthfi Izzuddin", bodyString, "Result must be same as the query data")
}

func SayHelloMultipleValue(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleValueQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Luthfi&name=Izzuddin&name=Hanif", nil)
	recorder := httptest.NewRecorder()

	SayHelloMultipleValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	assert.Equal(t, "Luthfi Izzuddin Hanif", bodyString, "Result must be same as the query data")
}
