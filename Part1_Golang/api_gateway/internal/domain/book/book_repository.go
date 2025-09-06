package book

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	List(ctx context.Context) ([]Book, error)
	Get(ctx context.Context, id uuid.UUID) (Book, error)
	Add(ctx context.Context, book Book) (Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
