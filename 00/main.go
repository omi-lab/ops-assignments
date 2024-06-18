package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Instantiating a new mux server...")

	router := http.NewServeMux()

	fmt.Println("Attaching the handlers...")

	router.HandleFunc("GET /v1/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("POST /v1/files", func(w http.ResponseWriter, r *http.Request) {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		dst, err := os.Create(filepath.Join("/tmp", handler.Filename))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error creating the file on server", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Error copying the file on server", http.StatusInternalServerError)
			return
		}

		r.MultipartForm.RemoveAll()

		w.WriteHeader(http.StatusCreated)
	})

	fmt.Println("Starting the server...")

	if err := http.ListenAndServe(":3000", router); !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("An error occured while running the server: %v\n", err)
	}

	fmt.Println("The server was successfully closed.")
}
