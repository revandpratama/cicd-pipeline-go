package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	router := setupRouter()

	http.ListenAndServe(":8080", router)

}

func setupRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"ping": "pong"})
	})

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"message": "Hello From CICD Pipeline!"})
	})
	return router
}
