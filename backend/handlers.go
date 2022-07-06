package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

  userId, err := strconv.Atoi(userTipoCookie.Value)
	if err != nil {
		fmt.Println("Erro Atoi: ", err)
	}

  tipo := strings.ToLower(userIdCookie.Value)
  if tipo != "admin" && tipo != "escuderia" && tipo != "piloto" {
    return c.NoContent(http.StatusForbidden)
  }

	resultsByEachStatus, err := s.Store.GetResultsByEachStatus(intUserId, tipo.Value)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resultsByEachStatus)
}

//
func (s *Service) GetStatusRepor(c echo.Context) error {
	userId, err := c.Cookie("userid")
	tipo, _ := c.Cookie("tipo")c.Cookie("tipo")
	if err != nil {
		fmt.Println(err)
	}

	if tipo.Value != "Escuderia" {
		return c.NoContent(http.StatusForbidden)
	}

	intUserId, err := strconv.Atoi(userId.Value)
	if err != nil {
		fmt.Println("Erro Atoi: ", err)
	}
	GetResultsByEachStatus, err := s.Store.GetResultsByEachStatus(intUserId, tipo.Value)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, GetResultsByEachStatus)
}

//
func (s *Service) GetPilotoStatusReport(c echo.Context) error {
	userId, err := c.Cookie("userid")
	tipo, _ := c.Cookie("tipo")
	if err != nil {
		fmt.Println(err)
	}

	if tipo.Value != "PILOTO" {
		return c.NoContent(http.StatusForbidden)
	}

	intUserId, err := strconv.Atoi(userId.Value)
	if err != nil {
		fmt.Println("Erro Atoi: ", err)
	}
	GetResultsByEachStatus, err := s.Store.GetResultsByEachStatus(intUserId, tipo.Value)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, GetResultsByEachStatus)
}

// Test
func (s *Service) ReportAllUsers(c echo.Context) error {
	users, _ := s.Store.reportAllUsers()

	return c.JSON(http.StatusOK, users)
}

func (s *Service) RawSQL(c echo.Context) error {
	usernameCookie, _ := c.Cookie("username")

	if usernameCookie == nil {
		return c.NoContent(http.StatusForbidden)
	}

	if usernameCookie.Value != "mclaren_c" {
		return c.NoContent(http.StatusForbidden)
	}

	input := InputRawSQL{}

	if err := json.NewDecoder(c.Request().Body).Decode(&input); err != nil {
		return err
	}

	users, _ := s.Store.rawSQL(input)

	return c.JSON(http.StatusOK, users)
}
