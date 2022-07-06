package main

import (
	"github.com/labstack/echo/v4"
)

//Service struct
type Service struct {
	Server *echo.Echo
	Store  Store
}

type User struct {
	UserId     int    `db:"userid"`
	Login      string `json:"login" db:"login"`
	Password   string `json:"password" db:"userpassword"`
	Tipo       string `db:"tipo"`
	IdOriginal int    `db:"idoriginal"`
}

type GetResultsByEachStatus struct {
	Status string `json:"status" db:"status"`
	Count  int    `json:"count" db:"count"`
}
