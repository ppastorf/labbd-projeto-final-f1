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

	fmt.Println(query)

	if err := s.DB.Select(&user, query); err != nil {
		return nil, err
	} else if len(user) == 1 {
		fmt.Println("User found")
		return &user[0], nil
	} else {
		return nil, fmt.Errorf("User not found")
	}
}
