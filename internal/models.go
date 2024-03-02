package internal

import "github.com/google/uuid"

type User struct {
	ID string `json:"id"`
}

type Meme struct {
	UserID   string `json:"user_id"`
	ImageURL string `json:"image_url"`
}

func GenerateUserID() string {
	id := uuid.New()
	return id.String()
}
