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

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // diparsing dulu buat ngecek apakah body yg dikirim formatnya valid atau ga
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")

	// tapi ada cara langsung tanpa parse manual
	lastName := request.PostFormValue("last_name")

	fmt.Fprintf(writer, "%s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Luthfi&last_name=Izzuddin") // tanpa ? karena ga disimpen di url
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded") // wajib dikasih header kalo ngirim form body

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	assert.Equal(t, "Luthfi Izzuddin", bodyString, "Result must be same as 'Luthfi Izzuddin'")
}
