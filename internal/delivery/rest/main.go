package rest

import (
	"github.com/febriliankr/go-cfstore-api/internal/app"
	"github.com/febriliankr/go-cfstore-api/internal/delivery/rest/server"
	"github.com/febriliankr/go-cfstore-api/internal/delivery/rest/service"
)

// Start server
func Start(app *app.SeminarApp) {
	srv := server.New()
	svc := service.GetServices(app)

	srv.POST("/product", svc.CreateProductHandler)
	srv.GET("/product/:product_slug", svc.GetProductHandler)
	srv.GET("/product", svc.GetProductListHandler)
	srv.DELETE("/product/:product_slug", svc.DeleteProductHandler)
	srv.PATCH("/balance", svc.UpdateCanteenBalanceHandler)

	srv.POST("/user", svc.CreateUserHandler)
	srv.POST("/user/login", svc.LoginUserHandler)

	// serve /public as static files
	srv.Static("/", "./public")

	server.Start(srv, &app.Cfg.Server)
}
