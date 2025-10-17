package main

import (
	"architecture/internal/handler"
	"architecture/internal/repository"
	"architecture/internal/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

func main() {
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
			return
		}

		defer db.Close()

		if err := db.Ping(); err != nil {
			log.Printf("failed to connect to postgres: %v", err)
			return
		}
		log.Println("connected to postgres")

		diaryRepo := repository.NewDiaryRepository(db)
		diarySvc := service.NewDiaryService(diaryRepo)
		diaryHandler := handler.NewDiaryHandler(diarySvc)

		mux := http.NewServeMux()
		mux.HandleFunc("GET /users/{id}/posts",diaryHandler.GetUserWithPosts)

		fmt.Println("Server startiong on port 8080")

		if err := http.ListenAndServe(":8080",mux);err != nil{
			log.Fatalf("Faild to start server:%v",err)
		}

	
}
