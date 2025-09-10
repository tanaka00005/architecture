package main

import (
	"architecture/handler"
	"database/sql"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Post struct {
	Title   string
	Content string
	ID      uint
}

func getPosts(db *sql.DB, title string) ([]Post, error) {
	var posts []Post

	rows, err := db.Query("SELECT * FROM posts")

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", title, err)
	}

	defer rows.Close()

	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", title, err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", title, err)
	}
	return posts, nil
}

func main() {

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
		log.Panicf("failed to open database: %v", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Panicf("failed to connect to postgres: %v", err)
	}
	log.Println("connected to postgres")

	posts, err := getPosts(db, "hello")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Posts found: %v\n", posts)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
	})

	handler.Handler()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
