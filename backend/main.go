package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newService() *Service {
	service := &Service{
		Server: echo.New(),
		Store:  &StoreImpl{DB: createDB()},
	}

	return service
}

func main() {

	service := newService()
	defer service.Store.Close()

	//Middleware
	service.Server.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowMethods: []string{echo.GET, echo.POST},
			AllowOrigins: []string{"*"},
		}),
		middleware.Static("/app"),
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),
		AuthMiddleware(),
	)
	//Login
	service.Server.POST("/login", service.Login)

	//Relatorios
	service.Server.GET("/status-report", service.GetStatusReport)
	// service.Server.GET("/custom-report", service.GetCustomReport)
	// service.Server.GET("/overviewInfo", service.GetOverviewInfo)

	service.Server.Logger.Fatal(service.Server.Start(":8080"))
}
