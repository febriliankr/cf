package app

import (
	"github.com/febriliankr/go-cfstore-api/internal"
	"github.com/febriliankr/go-cfstore-api/internal/app/config"
	"github.com/febriliankr/go-cfstore-api/internal/usecases"
)

type Usecases struct {
	KantinUC internal.KantinUC
}

// Inject dependency for usecase layer
func newUsecases(cfg config.Config, repos *Repos) *Usecases {

	r := repos.KantinRepo

	kantinUC := usecases.NewKantinUsecase(r)
	uc := &Usecases{
		KantinUC: kantinUC,
	}
	return uc
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
