package application

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	memory "github.com/anho58/Assignment/part1_Golang/api_gateway/internal/infrastructure/memory"
)

func setupService() BookServiceInterface {
	repo := memory.NewBookRepo()
	return NewBookService(repo)
}

func TestBookService_List(t *testing.T) {
	service := setupService()
	ctx := context.Background()

	dto := CreateBookDTO{
		Title:       "Book 1",
		Author:      "Book 1 Author",
		Description: "",
		Price:       9.9,
	}

	book, err := service.Add(ctx, dto)
	assert.NoError(t, err)

	books, err := service.List(ctx)
	assert.NoError(t, err)
	assert.Contains(t, books, book)
}

func TestBookService_Add(t *testing.T) {
	service := setupService()
	ctx := context.Background()

	dto := CreateBookDTO{
		Title:       "Book 1",
		Author:      "Book 1 Author",
		Description: "",
		Price:       9.9,
	}

	b, err := service.Add(ctx, dto)
	assert.NoError(t, err)
	assert.Equal(t, "Book 1", b.Title)
	assert.NotEqual(t, uuid.Nil, b.ID)
}

func TestBookService_Get(t *testing.T) {
	service := setupService()
	ctx := context.Background()

	dto := CreateBookDTO{
		Title:  "Book 1",
		Author: "Book 1 Author",
		Price:  9.9,
	}

	book, _ := service.Add(ctx, dto)

	fetched, err := service.Get(ctx, book.ID)
	assert.NoError(t, err)
	assert.Equal(t, book.ID, fetched.ID)

	_, err = service.Get(ctx, uuid.New())
	assert.Error(t, err)
}

func TestBookService_Delete(t *testing.T) {
	service := setupService()
	ctx := context.Background()

	dto := CreateBookDTO{Title: "Delete Me", Author: "C", Price: 12}
	b, _ := service.Add(ctx, dto)

	err := service.Delete(ctx, b.ID)
	assert.NoError(t, err)

	_, err = service.Get(ctx, b.ID)
	assert.Error(t, err)
}
