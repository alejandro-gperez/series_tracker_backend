package handlers

import (
	"encoding/json"
	"net/http"
	"series-tracker-backend/internal/db"
	"series-tracker-backend/internal/models"
)

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

	// por si hubo error durante la iteración
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(series)
}
