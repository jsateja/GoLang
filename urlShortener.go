package main

import (
	"fmt"
	"log"
	"net/http"
	"urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/http-godoc": "https://golang.org/pkg/net/http",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlData := `
    - path: '/google'
      url: 'https://google.com/'
    - path: '/stackoverflow'
      url: 'https://stackoverflow.com/'
      `

	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlData), mapHandler)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
