package service
//アプリケーション層

import (
	"architecture/internal/model"
	"architecture/internal/repository"
	"fmt"
)

type UserWithPosts struct {
	Email string       `json:"email"`
	Name  string       `json:"name"`
	ID    uint         `json:"id"`
	Posts []model.Post `json:"posts"`
}

type DiaryService interface {
	GetUserWithPosts(userID uint) (*UserWithPosts, error)
}

type diaryService struct {
	db repository.DiaryRepository
}

func NewDiaryService(db repository.DiaryRepository) DiaryService {
	return &diaryService{db: db}
}

func (s *diaryService) GetUserWithPosts(userID uint) (*UserWithPosts, error) {
	posts, err := s.db.FindUserAll(userID)
	if err != nil {
		return nil, fmt.Errorf("postsの取得に失敗しました: %w", err)
	}

	// 本来はユーザー情報も取得する
	// user, err := s.repo.FindUserByID(userID)
	// ...

	// 取得した情報を使ってレスポンス用の構造体を組み立てて返す
	response := &UserWithPosts{
		ID:    userID,             // 本来はuser.ID
		Name:  "Test User",        // 本来はuser.Name
		Email: "test@example.com", // 本来はuser.Email
		Posts: posts,
	}

	return response, nil
}
