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
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),	
	)
	
	//Login
	service.Server.POST("/login", service.Login)
	
	service.Server.GET("/login", service.GetLoginView)
	service.Server.GET("/admin", service.GetAdminView)
	service.Server.GET("/escuderia", service.GetSquadronView)
	service.Server.GET("/piloto", service.GetPilotView)


	//Relatorios
	service.Server.GET("/admin/resultados-por-status", service.GetResultsByEachStatus)

	//Utils
	service.Server.POST("/sql", service.RawSQL)

	//Test
	service.Server.GET("/report/all-users", service.ReportAllUsers)

	service.Server.Logger.Fatal(service.Server.Start(":8080"))
}
