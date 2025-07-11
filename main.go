package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello From CICD Pipeline!"})
	})

	http.ListenAndServe(":8080", nil)

	// new comment
}
