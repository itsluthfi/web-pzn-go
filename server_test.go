package webpzngo

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) { // bikin web server yang membuat web bisa di akses di host & port tertentu
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
