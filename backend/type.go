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

type User struct {
  UserId      int     `db:"userid"`
  Login       string  `json:"login" db:"login"`
  Password    string  `json:"password" db:"password"`
  Tipo        string  `db:"tipo"`
  IdOriginal  int     `db:"idoriginal"`
}

type InputLogin struct {
  Login       string `json:"login"`
  Password    string `json:"password"`
}

type Store interface {
  Login(login string, password string) ([]User, error)
  Close() error
}

//StoreImpl struct
type StoreImpl struct {
	DB *sqlx.DB
}