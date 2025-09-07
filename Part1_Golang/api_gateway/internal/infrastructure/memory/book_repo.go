package memory

import (
	"context"
	"errors"
	"sync"

	domain "github.com/anho58/Assignment/part1_Golang/api_gateway/internal/domain/book"

	"github.com/google/uuid"
)

var ErrBookNotFound = errors.New("book not found")

type BookRepo struct {
	mu    sync.RWMutex
	items map[uuid.UUID]domain.Book
}

func NewBookRepo() *BookRepo {
	return &BookRepo{items: make(map[uuid.UUID]domain.Book)}
}

func (r *BookRepo) List(ctx context.Context) ([]domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	books := make([]domain.Book, 0, len(r.items))
	for _, v := range r.items {
		books = append(books, v)
	}
	return books, nil
}

func (r *BookRepo) Get(ctx context.Context, id uuid.UUID) (domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	book, success := r.items[id]
	if !success {
		return domain.Book{}, ErrBookNotFound
	}
	return book, nil
}

func (r *BookRepo) Add(ctx context.Context, book domain.Book) (domain.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	book.ID = uuid.New()
	r.items[book.ID] = book

	return book, nil
}

func (r *BookRepo) Delete(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.items[id]; !ok {
		return ErrBookNotFound
	}
	delete(r.items, id)

	return nil
}
