package main

import (
	"crypto/md5"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func md5Encrypt(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

type Store interface {
	Login(login string, password string) (*User, error)
	Close() error
	GetResultsByEachStatus(id int, tipo string) ([]ResultStatus, error)
	InsertConstructor(values CreateConstructorRequest) error
	InsertDriver(values CreateDriverRequest) error
	// GetAdminReport2(id int, tipo string) ([])
}

//StoreImpl struct
type StoreImpl struct {
	DB *sqlx.DB
}

func (s *StoreImpl) Close() error {
	return s.DB.Close()
}

// GetResultsByEachStatus
type ResultStatus struct {
	Status string `json:"status" db:"status"`
	Count  int    `json:"count" db:"count"`
}

func (s *StoreImpl) GetResultsByEachStatus(userId int, tipo string) ([]ResultStatus, error) {
	resultStatus := []ResultStatus{}

	switch tipo {
	case "admin":
		query := `select status.status, count(results.statusid)
    from results inner join status on results.statusid = status.statusid
    group by status.statusid
    order by count desc`

		if err := s.DB.Select(&resultStatus, query); err != nil {
			log.Println(err)
			return nil, err
		} else {
			return resultStatus, nil
		}

	case "escuderia":
		query := fmt.Sprintf(`select status.status, count(results.statusid) 
    from results inner join status on results.statusid = status.statusid inner join constructors on constructors.constructorid = results.constructorid
    where constructors.constructorid = %v
    group by status.statusid, constructors.constructorid
    order by count desc`, userId)

		if err := s.DB.Select(&resultStatus, query); err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			return resultStatus, nil
		}
	case "piloto":
		query := fmt.Sprintf(`select status.status, count(results.statusid)
    from results inner join status on results.statusid = status.statusid inner join driver on driver.driverid = results.driverid
    where driver.driverid = %v
    group by status.statusid, driver.driverid
    order by count desc`, userId)

		if err := s.DB.Select(&resultStatus, query); err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			return resultStatus, nil
		}

	default:
		return nil, fmt.Errorf("GetResultsByEachStatus default case")
	}
}

func (s *StoreImpl) Login(login string, password string) (*User, error) {
	user := []User{}

	// !!
	query := fmt.Sprintf(`SELECT * FROM users WHERE login='%v' AND userpassword='%v'`, login, md5Encrypt(password))

	if err := s.DB.Select(&user, query); err != nil {
		return nil, err
	} else if len(user) == 1 {
		fmt.Println("User found")
		return &user[0], nil
	} else {
		return nil, fmt.Errorf("User not found")
	}
}

func (s *StoreImpl) insert(table, query string) error {
	result, err := s.DB.Exec(query)
	if err != nil {
		log.Printf("Storage error: %s\n", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		log.Printf("Failed to insert into table %s\n", err.Error())
	}
	log.Printf("Inserted row into table %s\n", table)

	return nil
}

func (s *StoreImpl) InsertConstructor(values CreateConstructorRequest) error {
	query := fmt.Sprintf(`
    INSERT INTO public.Constructors (ConstructorRef, Name, Nationality, Url)
      VALUES('%s', '%s', '%s',' %s',);`,
		values.ConstructorRef, values.Name, values.Nationality, values.Url,
	)
	return s.insert(query, "Constructor")
}

func (s *StoreImpl) InsertDriver(values CreateDriverRequest) error {
	query := fmt.Sprintf(`
    INSERT INTO public.Drivers (Driverref, Number, Code, Forename, Surname, Date of Birth, Nationality)
      VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', );`,
		values.DriverRef, values.Number, values.Code, values.Forename, values.Surname, values.DateOfBirth, values.Nationality,
	)
	return s.insert(query, "Drivers")
}
