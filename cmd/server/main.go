package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/eroub/journal-api-go/internal/app" // Import the app package
    _ "github.com/joho/godotenv/autoload" // automatically load .env file
)

func main() {
    // Set up the database connection
    app.SetupDatabaseConnection() // Prefix with app
    defer app.CloseDatabaseConnection() // Prefix with app

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    })

    // Start the server
    log.Println("Server is starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
