package repository
//データアクセス層

import (
	"architecture/internal/model"
	"database/sql"
	"fmt"
	"log"
)

type UserWithPosts struct {
	ID    uint         `json:"id"`
	Name  string       `json:"name"`
	Email string       `json:"email"`
	Posts []model.Post `json:"posts"`
}

type diaryRepository struct {
	db *sql.DB
}

type DiaryRepository interface {
	FindUserAll(userId uint) ([]model.Post, error)
}

func NewDiaryRepository(db *sql.DB) DiaryRepository {
	return &diaryRepository{db: db}
}

func getPosts(db *sql.DB) ([]model.Post, error) {

	var posts []model.Post

	rows, err := db.Query("SELECT id, title, content, user_id, created_at, updated_at FROM posts")

	if err != nil {
		log.Printf("failed to get posts: %v", err)
		return nil, fmt.Errorf("err : %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			return nil, fmt.Errorf("err : %v", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("err : %v", err)
	}
	return posts, nil
}

func (r *diaryRepository) FindUserAll(userId uint) ([]model.Post, error) {
	posts, err := getPosts(r.db)
	if err != nil {
		log.Printf("failed to get posts: %v", err)
		return nil, fmt.Errorf("err : %v", err)
	}
	fmt.Printf("Posts found: %v\n", posts)
	return posts, nil
}
