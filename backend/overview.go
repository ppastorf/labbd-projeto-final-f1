package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) AdminOverview(c echo.Context) error {
	response := &AdminOverviewInfo{}
	var err error

	response, err = s.Store.GetAdminOverviewInfo()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get admin data")
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Service) ConstructorOverview(c echo.Context) error {
	response := &ConstructorOverviewInfo{}
	var err error

	userId, _, err := getAuthData(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to authenticate user: %s\n", err.Error())
	}

	response, err = s.Store.GetConstructorOverviewInfo(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get constructor data")
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Service) DriverOverview(c echo.Context) error {
	response := &DriverOverviewInfo{}
	var err error

	userId, _, err := getAuthData(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to authenticate user: %s\n", err.Error())
	}

	response, err = s.Store.GetDriverOverviewInfo(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get driver data")
	}
	return c.JSON(http.StatusOK, response)
}
