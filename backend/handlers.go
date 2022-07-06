package main

import (
	"github.com/labstack/echo/v4"
	"encoding/json"
	"net/http"
	"fmt"
  "time"
)

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

func (s *Service) GetAdminView(c echo.Context) error {
  usernameCookie, _ := c.Cookie("username")
  
  if usernameCookie == nil {
    return c.Redirect(http.StatusFound, "/login")
  }
  return c.HTML(http.StatusOK, renderHTML("frontend_test/admin.html", map[string]string{
    "nome": usernameCookie.Value,
  }))
}

func (s *Service) GetEscuderiaView(c echo.Context) error {
  usernameCookie, _ := c.Cookie("username")
  
  if usernameCookie == nil {
    return c.Redirect(http.StatusFound, "/login")
  }
  return c.HTML(http.StatusOK, renderHTML("frontend_test/escuderia.html", map[string]string{
    "nome": usernameCookie.Value,
  }))
}

func (s *Service) GetPilotoView(c echo.Context) error {
  usernameCookie, _ := c.Cookie("username")
  
  if usernameCookie == nil {
    return c.Redirect(http.StatusFound, "/login")
  }
  return c.HTML(http.StatusOK, renderHTML("frontend_test/piloto.html", nil))
}

func (s *Service) GetLoginView(c echo.Context) error {
  return c.HTML(http.StatusOK, renderHTML("frontend_test/login.html", map[string]string{}))
}

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
    cookie := new(http.Cookie)
	  cookie.Name = "username"
	  cookie.Value = user.Login
	  cookie.Expires = time.Now().Add(24 * time.Hour)
	  
    c.SetCookie(cookie)
    return c.Redirect(http.StatusFound, selectProperUserPage(user))
	}
}