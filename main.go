package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Gokul!")
}

func main() {
    // Set the handler for the root URL path
    http.HandleFunc("/", helloWorld)

    // Start the HTTP server on port 8081
    fmt.Println("Server starting on port 8081...")
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        fmt.Println("Error starting the server:", err)
    }
}