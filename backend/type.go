package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

//Service struct
type Service struct {
	Server *echo.Echo
	Store  Store
}

type Store interface {}

//StoreImpl struct
type StoreImpl struct {
	DB *sqlx.DB
}