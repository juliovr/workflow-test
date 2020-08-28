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
    log.Fatal(http.ListenAndServe(":8003", r))
}
