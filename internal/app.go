package app

import (
	"database/sql"

	"github.com/marcelofabianov/cashly/config"
	"github.com/marcelofabianov/cashly/internal/identity"
)

type App struct {
	Identity *identity.IdentityContainer
}

func NewApp(cfg *config.Config, db *sql.DB) *App {
	return &App{
		Identity: identity.NewIdentityContainer(cfg, db),
	}
}

func (a *App) Build() {
	a.Identity.Build()
}
