package main

import (
	"log"
	"net/http"
	"series-tracker-backend/internal/db"
	"series-tracker-backend/internal/handlers"
)

func main() {

	db.Connect()

	http.HandleFunc("/series/", handlers.GetSeriesByID)

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
	http.ListenAndServe(":3000", nil)
}
