// main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	startServer()
}

func startServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/*", ServeFrontendHandler)
	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"name": "Roshan",
		}
		jData, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	})

	fmt.Println("Starting server on port 9080...")
	http.ListenAndServe(":9080", r)
}
