package identity

import (
	"context"
	"database/sql"

	"github.com/marcelofabianov/cashly/internal/core"
	domain "github.com/marcelofabianov/cashly/internal/identity/domain"
	inbound "github.com/marcelofabianov/cashly/internal/identity/port/inbound"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, input inbound.CreateUserRepositoryInput) (*domain.User, error) {
	query := `
		INSERT INTO identity_users (public_id, identity_document, name, email, password, enabled, created_at, updated_at, version)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id;
	`

	data := inbound.UserMapping{
		PublicID:         input.User.PublicID.String(),
		IdentityDocument: input.User.IdentityDocument.String(),
		Name:             input.User.Name,
		Email:            input.User.Email.String(),
		Password:         input.User.Password,
		Enabled:          input.User.Enabled.Bool(),
		CreatedAt:        input.User.CreatedAt.String(),
		UpdatedAt:        input.User.UpdatedAt.String(),
		Version:          input.User.Version.Int(),
	}

	var id int64
	err := r.db.QueryRowContext(ctx, query, data.PublicID, data.IdentityDocument, data.Name, data.Email, data.Password, data.Enabled, data.CreatedAt, data.UpdatedAt, data.Version).Scan(&id)
	if err != nil {
		return nil, err
	}

	return domain.NewFromUser(
		core.ID(id),
		input.User.PublicID,
		input.User.IdentityDocument,
		input.User.Name,
		input.User.Email,
		input.User.Password,
		input.User.Enabled,
		input.User.CreatedAt,
		input.User.UpdatedAt,
		nil,
		input.User.Version,
	), nil
}

func (r *UserRepository) Exists(ctx context.Context, input inbound.UserExistsRepositoryInput) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM identity_users
			WHERE email = $1 OR identity_document = $2
		);
	`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, input.Email.String(), input.IdentityDocument.String()).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
