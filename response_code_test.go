package webpzngo

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(400) // BadRequest
		fmt.Fprint(writer, "name is empty")
	} else {
		writer.WriteHeader(http.StatusOK) // atau nulis response codenya bisa kayak gini
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)

	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	assert.Equal(t, "Hello Luthfi", bodyString, "Result must be same as 'Hello Luthfi'")
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/?name=Luthfi", nil)

	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	assert.Equal(t, "Hello Luthfi", bodyString, "Result must be same as 'Hello Luthfi'")
}
