package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/hello", helloHandler).Methods("GET")
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}