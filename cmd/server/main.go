package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.request) {
        fmt.Fprintln(w, "Hello, World!")
    })

    // Start the server
    fmt.Println("Server is starting on port 8080...")
    http.ListenAndServe(":8080", nil)
}
