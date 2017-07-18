package main

import (
    "log"
    "net/http"
)

// main create new HTTP handlers and start HTTP server on port 8080
func main() {

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}
