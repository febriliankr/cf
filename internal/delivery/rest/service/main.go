package service

import (
	"github.com/febriliankr/go-cfstore-api/internal/app"
)

type Services struct {
	*KantinService
}

func GetServices(app *app.SeminarApp) *Services {
	return &Services{
		KantinService: NewKantinService(app),
	}
}
