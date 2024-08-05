package webpzngo

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) { // buat static file, folder resources harus ikut dipindah waktu binary juga pindah
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	// waktu kita akses /static/index.js bakal 404 not found
	// karena file server carinya di folder resources gini -> /resources/static/index.js

	// cara biar fileserver cari tanpa ada prefix static bisa pake http.StripPrefix

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) { // ga perlu mikirin folder resources waktu udah jadi binary, karena udah embed
	directory, _ := fs.Sub(resources, "resources") // masuk ke dalem folder resources, karena kalo pake embed nama folder resource masuk ke string direktorinya
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
