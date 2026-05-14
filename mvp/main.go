package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// type Msg struct {
// 	URL string `json:"url"`
// }

func main() {
	var sites []string = make([]string, 0)
	sites = append(sites, "https://www.google.com/")

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Post("/block", func(w http.ResponseWriter, r *http.Request) {
		var msg struct {
			URL string `json:"url"`
		}

		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		log.Printf("URL: %s", msg.URL)

		if slices.Contains(sites, msg.URL) {
			w.Write([]byte("URL is in blocklist"))
		} else {
			w.Write([]byte("URL is not in blocklist"))
		}

	})

	http.ListenAndServe(":3598", r)
}
