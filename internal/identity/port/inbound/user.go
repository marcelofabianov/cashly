package identity

import (
	"context"

	"github.com/marcelofabianov/cashly/internal/core"
	domain "github.com/marcelofabianov/cashly/internal/identity/domain"
)

// PKG

type Hasher interface {
	Hash(data string) (string, error)
	Compare(data, hash string) (bool, error)
}

// DB Mapping

type UserMapping struct {
	ID               int64  `db:"id"`
	PublicID         string `db:"public_id"`
	IdentityDocument string `db:"identity_document"`
	Name             string `db:"name"`
	Email            string `db:"email"`
	Password         string `db:"password"`
	Enabled          bool   `db:"enabled"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
	DeletedAt        string `db:"deleted_at"`
	Version          int64  `db:"version"`
}

// Repository

type CreateUserRepositoryInput struct {
	User *domain.User
}

type CreateUserRepositoryOutput struct {
	User *domain.User
}

type UserExistsRepositoryInput struct {
	Email            core.Email
	IdentityDocument core.IdentityDocument
}

type CreateUserRepository interface {
	Exists(ctx context.Context, input UserExistsRepositoryInput) (bool, error)
	Create(ctx context.Context, input CreateUserRepositoryInput) (*domain.User, error)
}

// UseCase

type CreateUserUseCaseInput struct {
	IdentityDocument core.IdentityDocument
	Name             string
	Email            core.Email
	Password         string
}

type CreateUserUseCaseOutput struct {
	User *domain.User
}

type CreateUserUseCase interface {
	Execute(ctx context.Context, input CreateUserUseCaseInput) (*CreateUserUseCaseOutput, error)
}
