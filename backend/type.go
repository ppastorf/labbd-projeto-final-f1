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
	Login      string `form:"login" json:"login" db:"login"`
	Password   string `form:"password" json:"password" db:"userpassword"`
	Tipo       string `db:"tipo"`
	IdOriginal int    `db:"idoriginal"`
}

type Report2Input struct {
	Cidade string `query:"cidade"`
}

type Report3 struct {
	NomeCompleto string `json:"nome_completo" db:"nome_completo"`
	Vitorias int `json:"vitorias" db:"vitorias"`
}

type Report5 struct {
	Ano *int `json:"ano" db:"ano"`
	Corrida *string `json:"corrida" db:"corrida"`
	Vitoria *int `json:"vitoria" db:"vitoria"`
}

type Search struct {
	Nome 			*string `form:"name" json:"name" db:"name"`
	DataNascimento  *string `form:"data_nascimento" json:"data_nascimento" db:"data_nascimento"`
	Nacionalidade	*string	`form:"nacionalidade" json:"nacionalidade" db:"nacionalidade"`
}

type SearchInput struct {
	Sobrenome string `form:"sobrenome" json:"sobrenome" db:"sobrenome"`
}

type GetResultsByEachStatus struct {
	Status string `json:"status" db:"status"`
	Count  int    `json:"count" db:"count"`
}

type Report2 struct {
	NomeDaCidade *string	`json:"nome_da_cidade" db:"nome_da_cidade"`
	CodigoIATA *string `json:"codigo_iata" db:"codigo_iata"`
	NomeAeroporto *string `json:"nome_aeroporto" db:"nome_aeroporto"`
	CidadeAeroporto *string `json:"cidade_aeroporto" db:"cidade_aeroporto"`
	TipoAeroporto *string `json:"tipo_aeroporto" db:"tipo_aeroporto"`
	Distancia *string `json:"distancia" db:"distancia"`
}
