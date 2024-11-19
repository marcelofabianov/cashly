package identity

import (
	"database/sql"

	"go.uber.org/dig"

	"github.com/marcelofabianov/cashly/config"
	usecase "github.com/marcelofabianov/cashly/internal/identity/domain/usecase"
	repository "github.com/marcelofabianov/cashly/internal/identity/infra/repository"
	inbound "github.com/marcelofabianov/cashly/internal/identity/port/inbound"

	"github.com/marcelofabianov/cashly/pkg/hasher"
)

type IdentityContainer struct {
	*dig.Container
	cfg *config.Config
	db  *sql.DB
}

func NewIdentityContainer(cfg *config.Config, db *sql.DB) *IdentityContainer {
	return &IdentityContainer{
		Container: dig.New(),
		cfg:       cfg,
		db:        db,
	}
}

func (c *IdentityContainer) Build() {
	c.registerPackages()
	c.registerRepository()
}

func (c *IdentityContainer) registerPackages() {
	c.Provide(func() inbound.Hasher {
		return hasher.NewHasher()
	})
}

func (c *IdentityContainer) registerRepository() {
	c.Provide(func() inbound.CreateUserRepository {
		return repository.NewUserRepository(c.db)
	})
}

func (c *IdentityContainer) registerUseCase() {
	c.Provide(func(r inbound.CreateUserRepository, h inbound.Hasher) inbound.CreateUserUseCase {
		return usecase.NewCreateUserUseCase(r, h)
	})
}
