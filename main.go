package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
    port := os.Getenv("PORT") // Taking PORT value from the environemt
	
    mux := http.NewServeMux() // Initialise the Mux Engine

    // If the value of the PORT is empty, assign port with default value (:1234)
    if port == "" {
        port = "1234"
    }

    // Default index URL, this is for testing only
    mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("/"))
    })

    //  Run the engine
    fmt.Printf("Running on port :%s\n", port)
    err := http.ListenAndServe(":"+port, mux)
    if err != nil {
        panic(err)
    }
}
