package server

import (
	"fmt"
	"net/http"
)

func StartREST() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "User Service REST API is running")
	})
	http.ListenAndServe(":8081", nil)
}
