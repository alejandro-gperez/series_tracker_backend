package handlers

import (
	"encoding/json"
	"net/http"
	"series-tracker-backend/internal/db"
	"series-tracker-backend/internal/models"
)
//Endpoint GET para todas las series en la DB
func GetSeries(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, description, image FROM series")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	series := []models.Series{}

	for rows.Next() {
		var s models.Series
		err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.Image)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		series = append(series, s)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(series)
}

//GET de serie por id.
func GetSeriesByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/series/"):]

	row := db.DB.QueryRow("SELECT id, name, description, image FROM series WHERE id=$1", id)

	var s models.Series
	err := row.Scan(&s.ID, &s.Name, &s.Description, &s.Image)

	if err != nil {
		http.Error(w, "Series not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}
