package main

import (
    "fmt"
    "log"
    "net/http"
)
func handler(W http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(W, "Welcome to my home page\n", r.URL.Path[1:])
}
func handler2(W http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(W, "hello world\n")
}
func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}