package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
  "strings"

	"github.com/labstack/echo/v4"
)

// Login
func (s *Service) Login(c echo.Context) error {

	input := InputLogin{}
	user := &User{}
	var err error

	if err := json.NewDecoder(c.Request().Body).Decode(&input); err != nil {
		return err
	}

	if user, err = s.Store.Login(input.Login, input.Password); err != nil {
		return c.JSON(http.StatusBadRequest, "Username or password  not found")
	} else {
		fmt.Printf("User %s logged", user.Login)
		usernameCookie := new(http.Cookie)
		usernameCookie.Name = "username"
		usernameCookie.Value = user.Login
		usernameCookie.Expires = time.Now().Add(24 * time.Hour)

		c.SetCookie(usernameCookie)

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

		return c.Redirect(http.StatusFound, selectProperUserPage(user))
	}
}

// Login

// Login Views
func (s *Service) GetAdminView(c echo.Context) error {
	usernameCookie, _ := c.Cookie("username")
	tipoCookie, _ := c.Cookie("tipo")

	if usernameCookie == nil {
		return c.Redirect(http.StatusFound, "/login")
	}
	if tipoCookie.Value != "Admin" {
		return c.NoContent(http.StatusForbidden)
	}

	return c.HTML(http.StatusOK, renderHTML("frontend_test/admin.html", map[string]string{
		"nome": usernameCookie.Value,
	}))
}

func (s *Service) GetConstructorView(c echo.Context) error {
	usernameCookie, _ := c.Cookie("username")
	tipoCookie, _ := c.Cookie("tipo")

	if usernameCookie == nil {
		return c.Redirect(http.StatusFound, "/login")
	}
	if tipoCookie.Value != "Escuderia" {
		return c.NoContent(http.StatusForbidden)
	}

	return c.HTML(http.StatusOK, renderHTML("frontend_test/escuderia.html", map[string]string{
		"nome": usernameCookie.Value,
	}))
}

func (s *Service) GetPilotView(c echo.Context) error {
	usernameCookie, _ := c.Cookie("username")

	if usernameCookie == nil {
		return c.Redirect(http.StatusFound, "/login")
	}
	return c.HTML(http.StatusOK, renderHTML("frontend_test/piloto.html", nil))
}

func (s *Service) GetLoginView(c echo.Context) error {
	return c.HTML(http.StatusOK, renderHTML("frontend_test/login.html", map[string]string{}))
}

// Login Views

// Reports
func (s *Service) GetStatusReport(c echo.Context) error {
	userIdCookie, err := c.Cookie("userid")
	userTipoCookie, _ := c.Cookie("tipo")
	if err != nil {
		fmt.Println(err)
	}

  userId, err := strconv.Atoi(userIdCookie.Value)
	if err != nil {
		fmt.Println("Erro Atoi: ", err)
	}

  tipo := strings.ToLower(userTipoCookie.Value)
  if tipo != "admin" && tipo != "escuderia" && tipo != "piloto" {
    fmt.Println(tipo)
    return c.NoContent(http.StatusForbidden)
  }

	resultsByEachStatus, err := s.Store.GetResultsByEachStatus(userId, tipo)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resultsByEachStatus)
}

// func (s *Service) GetAdminReport2(c echo.Context) error {
//   userId, err := c.Cookie("userid")
//   tipo, _ := c.Cookie("tipo")
//   if err != nil {
//     fmt.Println(err)
//   }
  
//   if userId == nil {
//     return c.Redirect(http.StatusForbidden, "/login")
//   }

//   if tipo.Value != "Admin" {
//     return c.NoContent(http.StatusForbidden)
//   }

//   intUserId, err := strconv.Atoi(userId.Value)
//   if err != nil {
//     fmt.Println("Erro Atoi: ", err)
//   }
//   GetResultsByEachStatus, err := s.Store.GetAdminReport2(intUserId, tipo.Value)  
//   if err != nil {
// 		return err
// 	}
  
//   return c.JSON(http.StatusOK, GetResultsByEachStatus)
// }
