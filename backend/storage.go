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
	GetAdminOverviewInfo() (*AdminInfo, error)
	GetConstructorOverviewInfo(userId int) (*ConstructorInfo, error)
	GetDriverOverviewInfo(userId int) (*DriverInfo, error)
	GetAdminReport2(city string) ([]Report2, error)
	GetAdminReport3() ([]Report3, error)
	GetAdminReport5(userId string) ([]Report5, error)
	SearchPilot(userId string, sobrenome string) (Search, error)
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
		query := `
			SELECT status.status, COUNT(results.statusid)
    			FROM results INNER JOIN status ON results.statusid = status.statusid
				GROUP BY status.statusid
				ORDER BY count DESC`

		if err := s.DB.Select(&resultStatus, query); err != nil {
			log.Println(err)
			return nil, err
		} else {
			return resultStatus, nil
		}

	case "escuderia":
		query := fmt.Sprintf(`
			SELECT status.status, COUNT(results.statusid) 
				FROM results INNER JOIN status ON results.statusid = status.statusid INNER JOIN constructors ON constructors.constructorid = results.constructorid
				WHERE constructors.constructorid = %v
				GROUP BY status.statusid, constructors.constructorid
				ORDER BY count DESC
		`, userId)

		if err := s.DB.Select(&resultStatus, query); err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			return resultStatus, nil
		}
	case "piloto":
		query := fmt.Sprintf(`
		SELECT status.status, COUNT(results.statusid)
			FROM results INNER JOIN status ON results.statusid = status.statusid INNER JOIN driver ON driver.driverid = results.driverid
			WHERE driver.driverid = %v
			GROUP BY status.statusid, driver.driverid
			ORDER BY count DESC
		`, userId)

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
	query := fmt.Sprintf(`SELECT * FROM users WHERE login='%v' AND userpassword='%v'`, login, md5Encrypt(password))
	fmt.Printf("Login query: %s\n", query)

	if err := s.DB.Select(&user, query); err != nil {
		log.Println("Error db: ", err)
		return nil, err
	}
	if len(user) != 1 {
		log.Printf("User %s not found\n", login)
		return nil, fmt.Errorf("User not found")
	}
	return &user[0], nil
}

func (s *StoreImpl) GetAdminReport2(city string) ([]Report2, error) {
	report := []Report2{}

	query := fmt.Sprintf(`
	SELECT  c.name AS nome_da_cidade,
			iatacode AS codigo_iata,
			a.name AS nome_aeroporto,
			city AS cidade_aeroporto,
			TYPE AS tipo_aeroporto,
			earth_distance(ll_to_earth(a.latdeg, a.longdeg), ll_to_earth(c.lat, c.long)) AS distancia
		FROM (SELECT name, lat, long FROM geocities15k WHERE name='%v') c
		LEFT JOIN (
			SELECT iatacode, name, city, TYPE, latdeg, longdeg
	   		FROM airports
			WHERE isocountry='BR' AND (TYPE='medium_airport' OR TYPE='large_airport')
		) a ON earth_distance(ll_to_earth(a.latdeg, a.longdeg), ll_to_earth(c.lat, c.long)) <= 100000;
	`, city)

	fmt.Println(query)

	if err := s.DB.Select(&report, query); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return report, nil
}

func (s *StoreImpl) GetAdminReport3() ([]Report3, error) {
	report := []Report3{}
	query := fmt.Sprintf(`
		SELECT nome_completo, COALESCE(vitorias, 0) AS vitorias
			FROM (SELECT p.driverid, nome_completo
				FROM (SELECT DISTINCT driverid FROM results WHERE constructorid=9) p
				LEFT JOIN (SELECT driverid, CONCAT(forename, ' ', surname) AS nome_completoFROM driver) n ON p.driverid = n.driverid
			) parcial
			LEFT JOIN (
				SELECT driverid, count(*) AS vitorias
				FROM results WHERE constructorid=9 AND POSITION=1
				GROUP BY driverid
			) AS v ON parcial.driverid = v.driverid;
	`)

	fmt.Println(query)

	if err := s.DB.Select(&report, query); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return report, nil
}

func (s *StoreImpl) GetAdminReport5(userId string) ([]Report5, error) {
	report := []Report5{}
	query := fmt.Sprintf(`
		SELECT YEAR AS ano, name AS corrida, SUM(POSITION) AS vitoria
			FROM (SELECT POSITION, driverid, raceid FROM results WHERE driverid=%v AND POSITION=1) x
			LEFT JOIN (SELECT YEAR, raceid, name FROM races) y ON x.raceid = y.raceid
			GROUP BY ROLLUP(ano, corrida)
			ORDER BY ano, corrida;
	`, userId)

	fmt.Println(query)

	if err := s.DB.Select(&report, query); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return report, nil
}

func (s *StoreImpl) SearchPilot(userId string, sobrenome string) (Search, error) {
	search := []Search{}

	query := fmt.Sprintf(`SELECT forename, data_nascimento, nacionalidade FROM Results WHERE sobrenome=%v`, sobrenome)
	if err := s.DB.Select(&search, query); err != nil {
		return Search{}, err
	}

	return search[0], nil
}

// Admin overview
type AdminInfo struct {
	Pilotos    int
	Escuderias int
	Corridas   int
	Temporadas int
}

func (s *StoreImpl) GetAdminOverviewInfo() (*AdminInfo, error) {
	// query := fmt.Sprintf(``)
	return nil, nil
}

// Constructor overview
type ConstructorInfo struct {
	Vitorias      int
	PilotosUnicos int
	PrimeiroAno   int
	UltimoAno     int
}

func (s *StoreImpl) GetConstructorOverviewInfo(userId int) (*ConstructorInfo, error) {
	// query := fmt.Sprintf(``)
	return nil, nil
}

// Driver overview
type DriverInfo struct {
	Vitorias    int
	PrimeiroAno int
	UltimoAno   int
}

func (s *StoreImpl) GetDriverOverviewInfo(userId int) (*DriverInfo, error) {
	// query := fmt.Sprintf(``)
	return nil, nil
}
