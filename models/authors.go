package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Author struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthorBookResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Gender string    `json:"gender"`
	Email  string    `json:"email"`
}

type AuthorResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (author *Author) BeforeCreate(DB *gorm.DB) error {
	author.ID = uuid.New()
	return nil
}
