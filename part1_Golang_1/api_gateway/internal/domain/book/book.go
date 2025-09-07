package book

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	IsDelete    bool      `json:"isDelete"`
	IsSale      bool      `json:"isSale"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func New(id uuid.UUID, title, author, description string, price float64) Book {
	now := time.Now().UTC()

	if id == uuid.Nil {
		id = uuid.New()
	}

	return Book{
		ID:          id,
		Title:       title,
		Author:      author,
		Description: description,
		Price:       price,
		IsDelete:    false,
		IsSale:      false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
