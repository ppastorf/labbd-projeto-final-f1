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

// Login
func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().RequestURI == "/login" {
				return next(c)
			}
			userIdCookie, noIdCookie := c.Cookie("userid")
			userTipoCookie, noTipoCookie := c.Cookie("tipo")
			if noIdCookie != nil || noTipoCookie != nil || userIdCookie.Value == "" || userTipoCookie.Value == "" {
				return c.Redirect(http.StatusPermanentRedirect, "/loginPage")
			}
			return next(c)
		}
	}
}

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
func (s *Service) GetAdminReport2(c echo.Context) error {
	report := []Report2{}
	input := Report2Input{}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	fmt.Println(input.Cidade)

	report, err := s.Store.GetAdminReport2(input.Cidade)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, report)
}

func (s *Service) GetAdminReport3(c echo.Context) error {
	report := []Report3{}

	report, err := s.Store.GetAdminReport3()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, report)
}

func (s *Service) GetAdminReport5(c echo.Context) error {
	report := []Report5{}

	userIdCookie, err := c.Cookie("userid")

	report, err = s.Store.GetAdminReport5(userIdCookie.Value)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, report)
}

func (s *Service) SearchPilot(c echo.Context) error {

	userIdCookie, err := c.Cookie("userid")
	if err != nil {
		return err
	}

	fmt.Println(userIdCookie)

	input := SearchInput{}
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	fmt.Println(input.Sobrenome)

	search, err := s.Store.SearchPilot(input.Sobrenome, userIdCookie.Value)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, search)
}
