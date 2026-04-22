package main

import (
	"log"
	"net/http"
	"os"
	"series-tracker-backend/internal/db"
	"series-tracker-backend/internal/handlers"
)

// CORS So that any frontend can reach
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	db.Connect()

	// Handling by ID
	http.HandleFunc("/series/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetSeriesByID(w, r)
		case http.MethodPut:
			handlers.UpdateSeries(w, r)
		case http.MethodDelete:
			handlers.DeleteSeries(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Listing and creating
	http.HandleFunc("/series", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetSeries(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateSeries(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/marco", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("polo"))
	})

	log.Println("Server running on :3000")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server running on :" + port)
	http.ListenAndServe(":"+port, enableCORS(http.DefaultServeMux))
}
