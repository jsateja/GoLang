package main

import (
        "fmt"
        "log"
        "net/http"

        "github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintf(w, "Welcome home!")
}

func main() {
     router := mux.NewRouter().StrictSlash(true)
     router.HandleFunc("/api/{module}/{env}", homeLink)
     log.Fatal(http.ListenAndServe(":8082", router))
}
