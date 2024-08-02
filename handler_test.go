package webpzngo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) { // buat nerima request dari klien, tapi cuman bisa bikin 1 handler/gabisa multiple
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "Hello world") // write = response yang mau dibalikin klien
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) { // mux bisa buat bikin multiple handler
	mux := http.NewServeMux() // mux juga bisa disebut sebagai router

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) { // kalo dikasih slash (/), misal akses /images/test itu bakal munculin Images, tapi kalo gaada slash (/) bakal muncul 404 not found
		fmt.Fprint(w, "Images")
	})

	mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Thumbnails")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
