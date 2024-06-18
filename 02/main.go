package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Instantiating a new mux server...")

	router := http.NewServeMux()

	fmt.Println("Attaching the handlers...")

	router.HandleFunc("GET /v1/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("GET /v1/load", func(w http.ResponseWriter, r *http.Request) {
		// Implement the logic you need to artificially increase CPU usage

		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Starting the server...")

	if err := http.ListenAndServe(":3000", router); !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("An error occured while running the server: %v\n", err)
	}

	fmt.Println("The server was successfully closed.")
}
