package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ============================================================================
// Admin
// ============================================================================
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

// ============================================================================
// Constructor
// ============================================================================
// Search pilot by forename
func (s *Service) SearchPilot(c echo.Context) error {
	userId, _, err := getAuthData(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to authenticate user: %s\n", err.Error())
	}

	forename := c.QueryParam("sobrenome")
	result, err := s.Store.SearchPilotByForename(userId, forename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, result)
}
