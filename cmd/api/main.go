package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Handler() {
	http.HandleFunc("/ooo", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "ooo, %q", html.EscapeString(r.URL.Path))

		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		fmt.Printf("dbHost:%v", dbHost)

		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPassword, dbName)
		db, err := sql.Open("postgres", connStr)

		if err != nil {
			log.Printf("failed to open to postgres: %v", err)
			http.Error(w, "failed to open to postgres", http.StatusInternalServerError)
			return
		}

		defer db.Close()

		if err := db.Ping(); err != nil {
			log.Printf("failed to connect to postgres: %v", err)
			http.Error(w, "ailed to connect to postgres", http.StatusInternalServerError)
			return
		}
		log.Println("connected to postgres")

	})
}
