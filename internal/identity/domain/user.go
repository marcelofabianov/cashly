package identity

import "github.com/marcelofabianov/cashly/internal/core"

type User struct {
	ID               core.ID
	PublicID         core.PublicID
	OwershipID       core.OwnershipID
	IdentityDocument core.IdentityDocument
	Name             string
	Email            core.Email
	Password         string
	Enabled          core.Enabled
	CreatedAt        core.CreatedAt
	UpdatedAt        core.UpdatedAt
	DeletedAt        core.DeletedAt
	Version          core.Version
}

func NewUser(ownershipID core.OwnershipID, identityDocument core.IdentityDocument, name string, email core.Email, password string) *User {
	return &User{
		ID:               core.NewID(),
		PublicID:         core.NewPublicID(),
		OwershipID:       ownershipID,
		IdentityDocument: identityDocument,
		Name:             name,
		Email:            email,
		Password:         password,
		Enabled:          core.NewEnabled(),
		CreatedAt:        core.NewCreatedAt(),
		UpdatedAt:        core.NewUpdatedAt(),
		DeletedAt:        nil,
		Version:          core.NewVersion(),
	}
}

func NewFromUser(
	id core.ID,
	publicID core.PublicID,
	ownershipID core.OwnershipID,
	identityDocument core.IdentityDocument,
	name string,
	email core.Email,
	password string,
	enabled core.Enabled,
	createdAt core.CreatedAt,
	updatedAt core.UpdatedAt,
	deletedAt core.DeletedAt,
	version core.Version,
) *User {
	return &User{
		ID:               id,
		PublicID:         publicID,
		OwershipID:       ownershipID,
		IdentityDocument: identityDocument,
		Name:             name,
		Email:            email,
		Password:         password,
		Enabled:          enabled,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		DeletedAt:        deletedAt,
		Version:          version,
	}
}
