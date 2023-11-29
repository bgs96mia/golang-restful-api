package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	AuthorID    uuid.UUID `gorm:"size:191" json:"author_id" `
	Author      Author    `gorm:"foreignKey:AuthorID" json:"author"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookResponse struct {
	ID          uuid.UUID          `gorm:"primaryKey" json:"id"`
	Title       string             `json:"title"`
	AuthorID    uuid.UUID          `gorm:"size:191" json:"-"`
	Author      AuthorBookResponse `gorm:"foreignKey:AuthorID" json:"author"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

func (book *Book) BeforeCreate(BD *gorm.DB) error {
	book.ID = uuid.New()
	return nil

}
