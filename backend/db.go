package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "172.20.0.2"
	port     = 5432
	user     = "postgres"
	password = "changeme"
	dbname   = "postgres"
)

var schema = `
	CREATE TABLE member (
		id text PRIMARY KEY,
		name text NOT NULL,
		phone text NOT NULL,
		email text NOT NULL UNIQUE,
		password text NOT NULL,
		republica text,
		isAdmin	bool DEFAULT FALSE
);`

func getConnString() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return psqlInfo
}

func createDB() *sqlx.DB {
	db, err := sqlx.Open("postgres", getConnString())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

  // Ping test
  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("TO AQUI")
  _, err = db.Exec(schema)
  if err != nil {
    panic(err)
  }


	return db
}


