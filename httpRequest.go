package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	// "net/url"
)

func HandleMessenger(resp http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		// u, _ := url.Parse(request.RequestURI)
		// values, _ := url.ParseQuery(u.RawQuery)
		resp.WriteHeader(200)
		resp.Write([]byte(`Some response`))
		return
	}
}

// Initialize request
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HandleMessenger).Methods("POST", "GET")
	port := ":8000"
	log.Printf("Server started on %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
