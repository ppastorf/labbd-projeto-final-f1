package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type InputLogin struct {
	Login    string `json:"login" form:"login"`
	Password string `json:"password" form:"password"`
}

type User struct {
	UserId     int    `db:"userid"`
	Login      string `form:"login" json:"login" db:"login"`
	Password   string `form:"password" json:"password" db:"userpassword"`
	Tipo       string `db:"tipo"`
	IdOriginal int    `db:"idoriginal"`
}

func (s *Service) Login(c echo.Context) error {
	input := InputLogin{}
	user := &User{}
	var err error

	if err = c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if user, err = s.Store.GetLoginUser(input.Login, input.Password); err != nil {
		return c.JSON(http.StatusBadRequest, "login or password  not found")
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

	redirectUrl := fmt.Sprintf("/%s/", strings.ToLower(user.Tipo))

	return c.Redirect(http.StatusFound, redirectUrl)
}
