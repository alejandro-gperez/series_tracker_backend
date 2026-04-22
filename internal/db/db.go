package db

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	// 🔥 FIX CLAVE: asegurar SSL para Render
	if !strings.Contains(connStr, "sslmode=") {
		if strings.Contains(connStr, "?") {
			connStr += "&sslmode=require"
		} else {
			connStr += "?sslmode=require"
		}
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
}
