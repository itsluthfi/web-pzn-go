package webpzngo

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type") // key header ga case sensitive kayak query param

	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json") // kirim dari client ke server

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	assert.Equal(t, "application/json", bodyString, "Result must be same as 'application/json'")
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Luthfi Izzuddin Hanif") // kirim dari server ke client

	fmt.Fprint(writer, "OK!")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	header := recorder.Header().Get("x-powered-by")

	assert.Equal(t, "OK!", bodyString, "Result must be same as 'OK!'")
	assert.Equal(t, "Luthfi Izzuddin Hanif", header, "Result must be same as 'Luthfi Izzuddin Hanif'")
}
