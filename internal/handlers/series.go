package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"series-tracker-backend/internal/db"
	"series-tracker-backend/internal/models"
	"strconv"
)

// Endpoint GET for every serie in the DB. Search and filter by name or id with query validation (anti sql injection).
func GetSeries(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("q")
	sortParam := r.URL.Query().Get("sort")
	orderParam := r.URL.Query().Get("order")

	//SORT Validation
	if sortParam != "name" && sortParam != "id" {
		sortParam = "id"
	}

	//Validate order
	if orderParam != "asc" && orderParam != "desc" {
		orderParam = "asc"
	}

	//BASE QUERY
	baseQuery := "SELECT id, name, description, image FROM series"

	var rows *sql.Rows
	var err error

	//IF SEARCHING
	if queryParam != "" {
		baseQuery += " WHERE LOWER(name) LIKE LOWER($1)"
		baseQuery += " ORDER BY " + sortParam + " " + orderParam

		rows, err = db.DB.Query(baseQuery, "%"+queryParam+"%")
	} else {
		//NOT SEARCHING
		baseQuery += " ORDER BY " + sortParam + " " + orderParam

		rows, err = db.DB.Query(baseQuery)
	}

	//QUERY ERROR
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	//RESULTS
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

	//ITERATING ERROR
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//RESPONSE
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(series)
}

// GET series by id
func GetSeriesByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/series/"):]

	row := db.DB.QueryRow("SELECT id, name, description, image FROM series WHERE id=$1", id)

	var s models.Series
	err := row.Scan(&s.ID, &s.Name, &s.Description, &s.Image)

	if err == sql.ErrNoRows {
		http.Error(w, "Series not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

// POST create series in the DB
func CreateSeries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var s models.Series

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if s.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	err = db.DB.QueryRow(
		"INSERT INTO series (name, description, image) VALUES ($1, $2, $3) RETURNING id",
		s.Name, s.Description, s.Image,
	).Scan(&s.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

// PUT Update existing series in DB
func UpdateSeries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse ID correctly
	idStr := r.URL.Path[len("/series/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var s models.Series
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if s.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec(
		"UPDATE series SET name=$1, description=$2, image=$3 WHERE id=$4",
		s.Name, s.Description, s.Image, id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Series not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DELETE Delete an existing series from DB
func DeleteSeries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse ID correctly
	idStr := r.URL.Path[len("/series/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("DELETE FROM series WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Series not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
