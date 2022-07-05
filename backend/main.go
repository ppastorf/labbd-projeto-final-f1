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
	
	service.Server.POST("/login", service.login)

	service.Server.Logger.Fatal(service.Server.Start(":8080"))
}
