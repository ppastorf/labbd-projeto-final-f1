package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetLoginUser(login string, password string) (*User, error)
	Close() error

	GetStatusReport(id int, tipo string) ([]StatusResult, error)

	InsertConstructor(values CreateConstructorRequest) error
	InsertDriver(values CreateDriverRequest) error
	SearchPilotByForename(constructorId int, forename string) (PilotInfo, error)

	GetAdminOverviewInfo() (*AdminOverviewInfo, error)
	GetConstructorOverviewInfo(userId int) (*ConstructorOverviewInfo, error)
	GetDriverOverviewInfo(userId int) (*DriverOverviewInfo, error)

	GetAdminReport(city string) ([]AdminReport, error)
	GetConstructorReport(constructorId int) ([]ConstructorReport, error)
	GetDriverReport(driverId int) ([]DriverReport, error)
}

type StoreImpl struct {
	DB *sqlx.DB
}

func (s *StoreImpl) Close() error {
	return s.DB.Close()
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

// ============================================================================
// Login
// ============================================================================
func (s *StoreImpl) GetLoginUser(login string, password string) (*User, error) {
	user := []User{}
	query := fmt.Sprintf(`SELECT * FROM users WHERE login='%v' AND userpassword='%v'`, login, md5Encrypt(password))
	fmt.Printf("Login query: %s\n", query)

	if err := s.DB.Select(&user, query); err != nil {
		log.Println("Error db: ", err.Error())
		return nil, err
	}
	if len(user) != 1 {
		log.Printf("User %s not found\n", login)
		return nil, fmt.Errorf("User not found")
	}
	return &user[0], nil
}

// ============================================================================
// Funcionalidades dos usuarios
// ============================================================================
// Admin
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

// Constructor
type PilotInfo struct {
	Nome           *string `form:"name" json:"name" db:"name"`
	DataNascimento *string `form:"data_nascimento" json:"data_nascimento" db:"data_nascimento"`
	Nacionalidade  *string `form:"nacionalidade" json:"nacionalidade" db:"nacionalidade"`
}

func (s *StoreImpl) SearchPilotByForename(constructorId int, forename string) (PilotInfo, error) {
	result := []PilotInfo{}

	query := fmt.Sprintf(`SELECT forename, data_nascimento, nacionalidade FROM Results WHERE sobrenome=%s`, forename)
	if err := s.DB.Select(&result, query); err != nil {
		log.Println("Failed to search Driver by forename: ", err.Error())
		return PilotInfo{}, err
	}
	return result[0], nil
}

// ============================================================================
// Overview
// ============================================================================

// Admin
type AdminOverviewInfo struct {
	Pilotos    int
	Escuderias int
	Corridas   int
	Temporadas int
}

func (s *StoreImpl) GetAdminOverviewInfo() (*AdminOverviewInfo, error) {
	// query := fmt.Sprintf(``)
	return nil, nil
}

// Constructor
type ConstructorOverviewInfo struct {
	Vitorias      int
	PilotosUnicos int
	PrimeiroAno   int
	UltimoAno     int
}

func (s *StoreImpl) GetConstructorOverviewInfo(userId int) (*ConstructorOverviewInfo, error) {
	// query := fmt.Sprintf(``)
	return nil, nil
}

// Driver
type DriverOverviewInfo struct {
	Vitorias    int
	PrimeiroAno int
	UltimoAno   int
}

func (s *StoreImpl) GetDriverOverviewInfo(userId int) (*DriverOverviewInfo, error) {
	// query := fmt.Sprintf(``)
	return nil, nil
}

// ============================================================================
// Reports
// ============================================================================
type StatusResult struct {
	Status string `json:"status" db:"status"`
	Count  int    `json:"count" db:"count"`
}

func (s *StoreImpl) GetStatusReport(userId int, userType string) ([]StatusResult, error) {
	statusResult := []StatusResult{}
	var query string
	switch userType {
	case "admin":
		query = `
			SELECT status.status, COUNT(results.statusid)
				FROM results INNER JOIN status ON results.statusid = status.statusid
				GROUP BY status.statusid
				ORDER BY count DESC`
	case "escuderia":
		query = fmt.Sprintf(`
			SELECT status.status, COUNT(results.statusid) 
				FROM results INNER JOIN status ON results.statusid = status.statusid INNER JOIN constructors ON constructors.constructorid = results.constructorid
				WHERE constructors.constructorid = %v
				GROUP BY status.statusid, constructors.constructorid
				ORDER BY count DESC
		`, userId)
	case "piloto":
		query = fmt.Sprintf(`
			SELECT status.status, COUNT(results.statusid)
				FROM results INNER JOIN status ON results.statusid = status.statusid INNER JOIN driver ON driver.driverid = results.driverid
				WHERE driver.driverid = %v
				GROUP BY status.statusid, driver.driverid
				ORDER BY count DESC
		`, userId)
	default:
		return nil, fmt.Errorf("GetResultsByEachStatus default case")
	}

	if err := s.DB.Select(&statusResult, query); err != nil {
		log.Printf("Failed to get status report for %s: %s\n", userType, err.Error())
		return nil, err
	}
	return statusResult, nil
}

type AdminReport struct {
	NomeDaCidade    *string `json:"nome_da_cidade" db:"nome_da_cidade"`
	CodigoIATA      *string `json:"codigo_iata" db:"codigo_iata"`
	NomeAeroporto   *string `json:"nome_aeroporto" db:"nome_aeroporto"`
	CidadeAeroporto *string `json:"cidade_aeroporto" db:"cidade_aeroporto"`
	TipoAeroporto   *string `json:"tipo_aeroporto" db:"tipo_aeroporto"`
	Distancia       *string `json:"distancia" db:"distancia"`
}

func (s *StoreImpl) GetAdminReport(city string) ([]AdminReport, error) {
	report := []AdminReport{}
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

	fmt.Printf("Admin Report query: %s\n", query)

	if err := s.DB.Select(&report, query); err != nil {
		log.Printf("Error getting Admin Report data: %s\n", err.Error())
		return nil, err
	}
	return report, nil
}

type ConstructorReport struct {
	NomeCompleto string `json:"nome_completo" db:"nome_completo"`
	Vitorias     int    `json:"vitorias" db:"vitorias"`
}

func (s *StoreImpl) GetConstructorReport(constructorId int) ([]ConstructorReport, error) {
	report := []ConstructorReport{}
	query := fmt.Sprintf(`
		SELECT nome_completo, COALESCE(vitorias, 0) AS vitorias
			FROM (SELECT p.driverid, nome_completo
				FROM (SELECT DISTINCT driverid FROM results WHERE constructorid=%d) p
				LEFT JOIN (SELECT driverid, CONCAT(forename, ' ', surname) AS nome_completoFROM driver) n ON p.driverid = n.driverid
			) parcial
			LEFT JOIN (
				SELECT driverid, count(*) AS vitorias
				FROM results WHERE constructorid=%d AND POSITION=1
				GROUP BY driverid
			) AS v ON parcial.driverid = v.driverid;
	`, constructorId, constructorId)

	fmt.Printf("Constructor Report query: %s\n", query)
	if err := s.DB.Select(&report, query); err != nil {
		log.Printf("Error getting Constructor Report data: %s\n", err.Error())
		return nil, err
	}
	return report, nil
}

type DriverReport struct {
	Ano     *int    `json:"ano" db:"ano"`
	Corrida *string `json:"corrida" db:"corrida"`
	Vitoria *int    `json:"vitoria" db:"vitoria"`
}

func (s *StoreImpl) GetDriverReport(driverId int) ([]DriverReport, error) {
	report := []DriverReport{}
	query := fmt.Sprintf(`
		SELECT YEAR AS ano, name AS corrida, SUM(POSITION) AS vitoria
			FROM (SELECT POSITION, driverid, raceid FROM results WHERE driverid=%d AND POSITION=1) x
			LEFT JOIN (SELECT YEAR, raceid, name FROM races) y ON x.raceid = y.raceid
			GROUP BY ROLLUP(ano, corrida)
			ORDER BY ano, corrida;
	`, driverId)

	fmt.Printf("Driver Report query: %s\n", query)
	if err := s.DB.Select(&report, query); err != nil {
		log.Printf("Error getting Driver Report data: %s\n", err.Error())
		return nil, err
	}
	return report, nil
}
