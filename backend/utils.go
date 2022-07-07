package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func md5Encrypt(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func strArrayContains(arr []string, val string) bool {
	for _, e := range arr {
		if strings.ToLower(e) == strings.ToLower(val) {
			fmt.Printf("tipo %s presente na lista de acesso\n", val)
			return true
		}
	}
	fmt.Printf("tipo nao %s presente na lista de acesso\n", val)
	return false
}

func getAuthData(c echo.Context) (userId int, userType string, err error) {
	userIdCookie, noUserId := c.Cookie("userid")
	userTipoCookie, noTipo := c.Cookie("tipo")

	if noUserId != nil || noTipo != nil || userIdCookie.Value == "" || userTipoCookie.Value == "" {
		return 0, "", fmt.Errorf("Missing authentication cookies")
	}

	userId, err = strconv.Atoi(userIdCookie.Value)
	if err != nil {
		return 0, "", fmt.Errorf("Invalid value for 'userid' cookie")
	}

	userType = strings.ToLower(userTipoCookie.Value)
	if userType != "admin" && userType != "escuderia" && userType != "piloto" {
		return 0, "", fmt.Errorf("Invalid value for 'tipo' cookie")
	}

	return userId, userType, nil
}
