package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "changeme"
	dbname   = "postgres"
)

func getConnString() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return psqlInfo
}

type Person struct {
  FirstName string `db:"first_name"`
  LastName  string `db:"last_name"`
  Email     string
}

func createDB() *sqlx.DB {
	db, err := sqlx.Open("postgres", getConnString())
	if err != nil {
		log.Fatal(err)
	}

  // Ping test
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  
  return db
}


