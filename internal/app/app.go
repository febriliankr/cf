package app

import (
	"github.com/caarlos0/env/v6"
	"github.com/febriliankr/go-cfstore-api/internal/app/config"

	"github.com/pkg/errors"
)

type StoreApp struct {
	Repos    *Repos
	Usecases *Usecases
	Cfg      config.Config
}

// NewStoreApp Initialize app and do dependency injection for all layers
func NewStoreApp() (*StoreApp, error) {
	cfg, err := readConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "error getting config")
	}

	app := new(StoreApp)

	app.Cfg = cfg

	app.Repos, err = newRepos(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "errors invoking newRepos")
	}

	app.Usecases = newUsecases(cfg, app.Repos)
	if err != nil {
		return nil, errors.Wrapf(err, "errors invoking newUsecases")
	}

	return app, nil
}

func (a *StoreApp) Close() []error {
	var errs []error
	errs = append(errs, a.Repos.Close()...)
	errs = append(errs, a.Usecases.Close()...)
	return errs
}

func readConfig() (config.Config, error) {
	var cfg config.Config

	// Read vars that exist in ENV
	if err := env.Parse(&cfg); err != nil {
		return config.Config{}, errors.Wrapf(err, "error reading config from ENV")
	}

	return cfg, nil
}
