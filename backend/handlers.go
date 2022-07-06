package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// Login
func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userIdCookie, noIdCookie := c.Cookie("userid")
			userTipoCookie, noTipoCookie := c.Cookie("tipo")
			if noIdCookie != nil || noTipoCookie != nil || userIdCookie.Value == "" || userTipoCookie.Value == "" {
				return c.Redirect(http.StatusPermanentRedirect, "/loginPage")
			}
			return next(c)
		}
	}
}

type InputLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (s *Service) Login(c echo.Context) error {
	input := InputLogin{}
	user := &User{}
	var err error

	if err = c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if user, err = s.Store.Login(input.Login, input.Password); err != nil {
		return c.JSON(http.StatusBadRequest, "Username or password  not found")
	}

	log.Printf("User %s logged", user.Login)

	tipoCookie := new(http.Cookie)
	tipoCookie.Name = "tipo"
	tipoCookie.Value = user.Tipo
	tipoCookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(tipoCookie)

	idCookie := new(http.Cookie)
	idCookie.Name = "userid"
	idCookie.Value = strconv.Itoa(user.UserId)
	idCookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(idCookie)

	return c.Redirect(http.StatusFound, "/overview")
}
