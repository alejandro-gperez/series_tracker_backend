package main

import (
	"log"
	"net/http"
	"series-tracker-backend/internal/db"
)

func main() {

	db.Connect()

	http.HandleFunc("/marco", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("polo"))
	})

	log.Println("Server running on :3000")
	http.ListenAndServe(":3000", nil)
}
