package main

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Utilize all CPU cores

	fmt.Println("Instantiating a new mux server...")

	router := http.NewServeMux()

	fmt.Println("Attaching the handlers...")

	router.HandleFunc("GET /v1/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("GET /v1/load", func(w http.ResponseWriter, r *http.Request) {
		n := 1000000 // Number of iterations
		fmt.Println("Calculating Pi with", n, "iterations")

		start := time.Now()
		pi := calculatePi(n)
		duration := time.Since(start)

		fmt.Printf("Estimated value of Pi: %.15f\n", pi)
		fmt.Printf("Calculation took %s\n", duration)

		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Starting the server...")

	if err := http.ListenAndServe(":3000", router); !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("An error occured while running the server: %v\n", err)
	}

	fmt.Println("The server was successfully closed.")
}

func calculatePi(n int) float64 {
	pi := 0.0
	for k := 0; k < n; k++ {
		pi += (4.0 * (1 - float64(k%2)*2)) / float64(2*k+1)
	}
	return pi
}
