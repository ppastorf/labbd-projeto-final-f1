package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newService() *Service {
	service := &Service{
		Server: echo.New(),
	}

	return service
}

func main() {
	
	service := newService()

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
	



	// v1.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	// Be careful to use constant time comparison to prevent timing attacks
	// 	// faz a query e tenta pega as info

	// 	// valida se a senha é igual
	// 	if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
	// 		subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// e.Static("/", "/app")

	// e.GET("/login", routes.LoginPage)

	service.Server.Logger.Fatal(service.Server.Start(":8080"))
}
