package model

import (
	"time"
)
type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	ID      uint   `json:"id"`
	UserID uint `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	ID    uint   `json:"id"`
}
