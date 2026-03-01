package main

import (
	"fmt"
	"net/http"
)

func main() {


	http.HandleFunc("/api/user/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("USER BFF OK"))
	})

	http.HandleFunc("/api/user/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Knockzip!")
	})

	http.ListenAndServe(":8080", nil)
}
