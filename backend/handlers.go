package main

import (
	"github.com/labstack/echo/v4"
	"encoding/json"
	"net/http"
	"fmt"
)

func (s *Service) login(c echo.Context) error {
	
	input := InputLogin{}
	user := []User{}
	var err error

	if err := json.NewDecoder(c.Request().Body).Decode(&input); err != nil {
		return err
	}

	if user, err = s.Store.Login(input.Login, input.Password); err != nil {
		return c.JSON(http.StatusBadRequest, "Username or password  not found")
	}

	fmt.Println(user) // test
	return c.JSON(http.StatusOK, "Logged In")
}