package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Service struct {
	Server *echo.Echo
	Store  Store
}

func newService() *Service {
	service := &Service{
		Server: echo.New(),
		Store:  &StoreImpl{DB: createDB()},
	}
	return service
}

func AuthMiddleware(allowedUserTypes []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, userType, err := getAuthData(c)
			if err != nil {
				return c.Redirect(http.StatusPermanentRedirect, "/index.html")
			}
			if !strArrayContains(allowedUserTypes, userType) {
				return echo.NewHTTPError(http.StatusForbidden, "Seu usuário não tem acesso a esse recurso.")
			}
			return next(c)
		}
	}
}

func main() {
	service := newService()
	defer service.Store.Close()

	// Middlewares
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

	// Rotas sem autenticacao
	service.Server.POST("/login", service.Login)

	// Rotas de Admin
	authenticated_admin := service.Server.Group("/admin", AuthMiddleware([]string{"admin"}))
	authenticated_admin.GET("/overview", service.AdminOverview)
	authenticated_admin.GET("/status-report", service.StatusReport)
	authenticated_admin.GET("/report", service.AdminReport)
	authenticated_admin.POST("/create-constructor", service.CreateConstructor)
	authenticated_admin.POST("/create-driver", service.CreateDriver)

	// Rotas de Escuderias
	authenticated_constructor := service.Server.Group("/constructor", AuthMiddleware([]string{"escuderia"}))
	authenticated_constructor.GET("/overview", service.ConstructorOverview)
	authenticated_constructor.GET("/status-report", service.StatusReport)
	authenticated_constructor.GET("/report", service.ConstructorReport)
	authenticated_constructor.GET("/search-driver", service.ConstructorOverview)

	// Rotas de Piloto
	authenticated_driver := service.Server.Group("/driver", AuthMiddleware([]string{"piloto"}))
	authenticated_driver.GET("/overview", service.DriverOverview)
	authenticated_driver.GET("/status-report", service.StatusReport)
	authenticated_driver.GET("/report", service.DriverReport)

	service.Server.Logger.Fatal(service.Server.Start(":8080"))
}
