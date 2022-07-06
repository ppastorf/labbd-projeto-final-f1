package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

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
