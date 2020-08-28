package main

import (
    "log"
    "net/http"
    "fmt"

    "github.com/gorilla/mux"
)

func main() {
    fmt.Println("Init server")
    r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "hello\n")
    }).Methods(http.MethodGet)
    log.Fatal(http.ListenAndServe(":8003", r))
}
