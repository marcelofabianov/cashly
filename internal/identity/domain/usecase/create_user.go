package identity

import (
	"context"

	domain "github.com/marcelofabianov/cashly/internal/identity/domain"
	domainError "github.com/marcelofabianov/cashly/internal/identity/domain/error"
	inbound "github.com/marcelofabianov/cashly/internal/identity/port/inbound"
)

type CreateUserUseCase struct {
	repository inbound.CreateUserRepository
	hasher     inbound.Hasher
}

func NewCreateUserUseCase(
	repository inbound.CreateUserRepository,
	hasher inbound.Hasher,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		repository: repository,
		hasher:     hasher,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input inbound.CreateUserUseCaseInput) (*inbound.CreateUserUseCaseOutput, error) {
	exists, err := uc.repository.Exists(ctx, inbound.UserExistsRepositoryInput{
		Email:            input.Email,
		IdentityDocument: input.IdentityDocument,
	})
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, domainError.NewErrUserAlreadyExists()
	}

	hashedPassword, err := uc.hasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user := domain.NewUser(input.IdentityDocument, input.Name, input.Email, hashedPassword)

	output, err := uc.repository.Create(ctx, inbound.CreateUserRepositoryInput{
		User: user,
	})
	if err != nil {
		return nil, err
	}

	return &inbound.CreateUserUseCaseOutput{
		User: output,
	}, nil
}
