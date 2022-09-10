// Package repository : file contains operations with all DBs
package repository

import (
	"context"

	"github.com/Egor-Tihonov/Book-network/internal/model"
)

// Repository middleware
type Repository interface {
	Create(ctx context.Context, person *model.Person) (string, error)
	UpdateAuth(ctx context.Context, id string, refreshToken string) error
	Update(ctx context.Context, id string, person *model.Person) error
	SelectByID(ctx context.Context, id string) (model.Person, error)
	Delete(ctx context.Context, id string) error
	SelectByIDAuth(ctx context.Context, id string) (model.Person, error)
}
