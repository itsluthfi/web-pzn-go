package webpzngo

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie) // bentuknya key:value
	cookie.Name = "X-LIH-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success create cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-LIH-Name")
	if err != nil {
		panic(err)
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Luthfi%20Izzuddin%20Hanif", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	response := recorder.Result().Cookies()

	for _, cookie := range response {
		fmt.Printf("Hello %s from cookie %s", cookie.Value, cookie.Name)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)

	cookie := new(http.Cookie)
	cookie.Name = "X-LIH-Name"
	cookie.Value = "Luthfi Izzuddin Hanif"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	fmt.Println(string(bodyString))
}
