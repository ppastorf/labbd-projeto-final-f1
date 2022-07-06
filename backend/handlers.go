package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
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

// Create driver
type CreateDriverRequest struct {
	DriverRef   string `json:"driverref"`
	Number      string `json:"number"`
	Code        string `json:"code"`
	Forename    string `json:"forename"`
	Surname     string `json:"surname"`
	DateOfBirth string `json:"date_of_birth"`
	Nationality string `json:"nationality"`
}

func (s *Service) CreateDriver(c echo.Context) error {
	request := CreateDriverRequest{}
	var err error

	if err = c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := s.Store.InsertDriver(request); err != nil {
		log.Printf("Error: %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Printf("Driver '%s %s' created", request.Forename, request.Surname)

	return c.NoContent(http.StatusCreated)
}

// Create constructor
type CreateConstructorRequest struct {
	ConstructorRef string `json:"constructorref"`
	Name           string `json:"name"`
	Nationality    string `json:"nationality"`
	Url            string `json:"url"`
}

func (s *Service) CreateConstructor(c echo.Context) error {
	request := CreateConstructorRequest{}
	var err error

	if err = c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := s.Store.InsertConstructor(request); err != nil {
		log.Printf("Error: %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Printf("Constructor '%s' created", request.Name)

	return c.NoContent(http.StatusCreated)
}

// Admin overview
func (s *Service) AdminInfo(c echo.Context) error {
	response := &AdminInfo{}
	var err error

	userTipoCookie, _ := c.Cookie("tipo")
	tipo := strings.ToLower(userTipoCookie.Value)

	userIdCookie, _ := c.Cookie("userid")
	userId, _ := strconv.Atoi(userIdCookie.Value)

	if tipo != "admin" && userId != 0 {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	response, err = s.Store.GetAdminInfo()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get admin data")
	}

	return c.JSON(http.StatusOK, response)
}

// Constructor overview
func (s *Service) ConstructorInfo(c echo.Context) error {
	response := &ConstructorInfo{}
	var err error

	userTipoCookie, _ := c.Cookie("tipo")
	tipo := strings.ToLower(userTipoCookie.Value)

	userIdCookie, _ := c.Cookie("userid")
	userId, _ := strconv.Atoi(userIdCookie.Value)

	if tipo != "escuderia" {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	response, err = s.Store.GetConstructorInfo(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get constructor data")
	}

	return c.JSON(http.StatusOK, response)
}

// Constructor overview
func (s *Service) DriverInfo(c echo.Context) error {
	response := &DriverInfo{}
	var err error

	userTipoCookie, _ := c.Cookie("tipo")
	tipo := strings.ToLower(userTipoCookie.Value)

	userIdCookie, _ := c.Cookie("userid")
	userId, _ := strconv.Atoi(userIdCookie.Value)

	if tipo != "piloto" {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	response, err = s.Store.GetDriverInfo(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get driver data")
	}

	return c.JSON(http.StatusOK, response)
}
