package app

import (
	"github.com/febriliankr/go-cfstore-api/internal"
	"github.com/febriliankr/go-cfstore-api/internal/app/config"
	"github.com/febriliankr/go-cfstore-api/internal/repo"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repos struct {
	KantinRepo internal.KantinRepo
}

// Inject dependency for repository layer
func newRepos(cfg config.Config) (*Repos, error) {
	db, err := initDB(cfg)
	if err != nil {
		return nil, err
	}
	return &Repos{
		KantinRepo: repo.NewKantinDB(db),
	}, nil
}

func initDB(cfg config.Config) (*sqlx.DB, error) {

	db, err := sqlx.Connect(cfg.DB.Driver, cfg.DB.Address)
	if err != nil {
		return nil, err
	}

	// Set db params
	db.SetMaxIdleConns(cfg.DB.MaxConns)
	db.SetMaxOpenConns(cfg.DB.MaxIdleConns)

	return db, nil
}

func (r *Repos) Close() []error {
	var errs []error

	return errs
}
