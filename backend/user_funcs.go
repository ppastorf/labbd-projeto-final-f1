package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ============================================================================
// Admin
// ============================================================================
// Create constructor
type CreateConstructorRequest struct {
	ConstructorRef string `json:"constructorref" form:"constructorref"`
	Name           string `json:"name" form:"name"`
	Nationality    string `json:"nationality" form:"nationality"`
	Url            string `json:"url" form:"url"`
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
	return c.JSON(http.StatusCreated, "Escuderia criada com sucesso")
}

// Create driver
type CreateDriverRequest struct {
	DriverRef   string `json:"driverref" form:"driverref"`
	Number      string `json:"number" form:"number"`
	Code        string `json:"code" form:"code"`
	Forename    string `json:"forename" form:"forename"`
	Surname     string `json:"surname" form:"surname"`
	DateOfBirth string `json:"date_of_birth" form:"date_of_birth"`
	Nationality string `json:"nationality" form:"nationality"`
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

	return c.JSON(http.StatusCreated, "Piloto criado com sucesso")
}

// ============================================================================
// Constructor
// ============================================================================
// Search pilot by forename
func (s *Service) SearchDriver(c echo.Context) error {
	userId, _, err := getAuthData(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to authenticate user: %s\n", err.Error())
	}

	forename := c.QueryParam("nome")
	result, err := s.Store.SearchPilotByForename(userId, forename)
	if err != nil {
		log.Println("Failed to search Driver by forename: ", err.Error())
		return err
	}

	return c.JSON(http.StatusOK, result)
}
