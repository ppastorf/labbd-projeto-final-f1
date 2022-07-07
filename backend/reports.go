package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Reports
func (s *Service) StatusReport(c echo.Context) error {
	userId, userType, err := getAuthData(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to authenticate user: %s\n", err.Error())
	}
	results, err := s.Store.GetStatusReport(userId, userType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get data: %s\n", err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

func (s *Service) AdminReport(c echo.Context) error {
	cidade := c.QueryParam("cidade")
	results, err := s.Store.GetAdminReport(cidade)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get Admin Report data: %s\n", err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

func (s *Service) ConstructorReport(c echo.Context) error {
	userId, _, err := getAuthData(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to authenticate user: %s\n", err.Error())
	}
	results, err := s.Store.GetConstructorReport(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get Constructor Report data: %s\n", err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

func (s *Service) DriverReport(c echo.Context) error {
	userId, _, err := getAuthData(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to authenticate user: %s\n", err.Error())
	}
	results, err := s.Store.GetDriverReport(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get Driver Report data: %s\n", err.Error())
	}

	return c.JSON(http.StatusOK, results)
}
