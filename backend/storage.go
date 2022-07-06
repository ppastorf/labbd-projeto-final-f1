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
    fmt.Println(err)
    return nil, err
	} else if len(user) == 1 {
		fmt.Println("User found")
		return &user[0], nil
	} else {
    fmt.Println("User not found")
		return nil, fmt.Errorf("User not found")
	}
}

func (s *StoreImpl) GetAdminReport2(city string) ([]Report2, error) {

	report := []Report2{}
	fmt.Println(city)

	// !!
	query := fmt.Sprintf(`
	SELECT c.name AS nome_da_cidade,
		   iatacode AS codigo_iata,
		   a.name AS nome_aeroporto,
		   city AS cidade_aeroporto,
		   TYPE AS tipo_aeroporto,
				   earth_distance(ll_to_earth(a.latdeg, a.longdeg), ll_to_earth(c.lat, c.long)) AS distancia
	FROM
	  (SELECT name,
			  lat, long
	   FROM geocities15k
	   WHERE name='%v') c
	LEFT JOIN
	  (SELECT iatacode,
			  name,
			  city,
			  TYPE,
			  latdeg,
			  longdeg
	   FROM airports
	   WHERE isocountry='BR'
		 AND (TYPE='medium_airport'
			  OR TYPE='large_airport')) a ON earth_distance(ll_to_earth(a.latdeg, a.longdeg), ll_to_earth(c.lat, c.long)) <= 100000;
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

	// !!
	query := fmt.Sprintf(`
	SELECT nome_completo,
       COALESCE(vitorias, 0) AS vitorias
FROM
  (SELECT p.driverid,
          nome_completo
   FROM
     (SELECT DISTINCT driverid
      FROM results
      WHERE constructorid=9) p --consulta pilotos distintos que correram pela escuderia

   LEFT JOIN
     (SELECT driverid,
             CONCAT(forename, ' ', surname) AS nome_completo
      FROM driver) n --obtém nome completo
 ON p.driverid = n.driverid) parcial
LEFT JOIN
  (SELECT driverid,
          count(*) AS vitorias
   FROM results
   WHERE constructorid=9
     AND POSITION=1
   GROUP BY driverid) AS v --obtém vitorias por aquela escuderia
ON parcial.driverid = v.driverid;`)

	fmt.Println(query)

	if err := s.DB.Select(&report, query); err != nil {
    	fmt.Println(err)
    	return nil, err
	}

	return report, nil
}

func (s *StoreImpl) GetAdminReport5(userId string) ([]Report5, error) {
	report := []Report5{}

	// !!
	query := fmt.Sprintf(`
	SELECT YEAR AS ano,
               name AS corrida,
               SUM(POSITION) AS vitoria
FROM
  (SELECT POSITION,
          driverid,
          raceid
   FROM results
   WHERE driverid=%v
     AND POSITION=1) x
LEFT JOIN
  (SELECT YEAR,
          raceid,
          name
   FROM races) y ON x.raceid = y.raceid
GROUP BY ROLLUP(ano, corrida)
ORDER BY ano,
         corrida;
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
