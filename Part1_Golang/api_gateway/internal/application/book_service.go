package application

import (
	"context"
	"errors"
	"part1_golang/api_gateway/internal/domain/book"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
)

type CreateBookDTO struct {
	Title       string  `json:"title" validate:"required,min=1,max=200"`
	Author      string  `json:"author" validate:"required,min=1,max=120"`
	Description string  `json:"description" validate:"min=1,max=2000"`
	Price       float64 `json:"price" validate:"gte=0,lte=100000"`
}

type BookServiceInterface interface {
	List(ctx context.Context) ([]book.Book, error)
	Get(ctx context.Context, id uuid.UUID) (book.Book, error)
	Add(ctx context.Context, dto CreateBookDTO) (book.Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type bookService struct {
	repo     book.Repository
	validate *validator.Validate
}

func NewBookService(r book.Repository) BookServiceInterface {
	return &bookService{repo: r, validate: validator.New()}
}

func (s *bookService) List(ctx context.Context) ([]book.Book, error) {
	return s.repo.List(ctx)
}

func (s *bookService) Get(ctx context.Context, id uuid.UUID) (book.Book, error) {
	b, err := s.repo.Get(ctx, id)
	if err != nil {
		return book.Book{}, ErrNotFound
	}
	return b, nil
}

func (s *bookService) Add(ctx context.Context, dto CreateBookDTO) (book.Book, error) {
	if err := s.validate.Struct(dto); err != nil {
		return book.Book{}, ErrInvalidInput
	}
	b := book.New(uuid.Nil, dto.Title, dto.Author, dto.Description, dto.Price)
	return s.repo.Add(ctx, b)
}

func (s *bookService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
