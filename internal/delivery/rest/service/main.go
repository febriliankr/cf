package service

import (
	"github.com/febriliankr/go-cfstore-api/internal/app"
)

type Services struct {
	*KantinService
}

func GetServices(app *app.StoreApp) *Services {
	return &Services{
		KantinService: NewKantinService(app),
	}
}
