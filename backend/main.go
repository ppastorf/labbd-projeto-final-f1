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
	)

	// Login
	service.Server.POST("/login", service.Login)

	authenticated := service.Server.Group("", AuthMiddleware())

	// Relatorios
	authenticated.GET("/status-report", service.GetStatusReport)
	// service.Server.GET("/custom-report", service.GetCustomReport)
	//Relatorios
	service.Server.GET("/status-report", service.GetStatusReport)
	service.Server.GET("/custom-report", service.GetAdminReport2)
	service.Server.GET("/report3", service.GetAdminReport3)
	service.Server.GET("/report5", service.GetAdminReport5)

	//Admin
	service.Server.GET("/report5", service.GetAdminReport5)
	service.Server.POST("/create-user", service.SearchPilot)

	// service.Server.GET("/overviewInfo", service.GetOverviewInfo)

	service.Server.Logger.Fatal(service.Server.Start(":8080"))
}
