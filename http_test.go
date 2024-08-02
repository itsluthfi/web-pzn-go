package webpzngo

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello world")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hi", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	assert.Equal(t, "Hello world", bodyString, "Result must be 'Hello world'")
}
