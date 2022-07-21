package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetLoginUser(login string, password string) (*User, error)
	Close() error

	GetStatusReport(id int, tipo string) ([]StatusResult, error)

	InsertConstructor(values CreateConstructorRequest) error
	InsertDriver(values CreateDriverRequest) error
	SearchPilotByForename(userId int, forename string) (PilotInfo, error)

	GetAdminOverviewInfo() (*AdminOverviewInfo, error)
	GetConstructorOverviewInfo(userId int) (*ConstructorOverviewInfo, error)
	GetDriverOverviewInfo(userId int) (*DriverOverviewInfo, error)

	GetAdminReport(city string) ([]AdminReport, error)
	GetConstructorReport(userId int) ([]ConstructorReport, error)
	GetDriverReport(driverId int) ([]DriverReport, error)
}

type StoreImpl struct {
	DB *sqlx.DB
}

func (s *StoreImpl) Close() error {
	return s.DB.Close()
}

func (s *StoreImpl) insert(table, query string) error {
	fmt.Printf("Insert query: %s\n\n\n\n", query)
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
    INSERT INTO public.Constructors (ConstructorId, ConstructorRef, Name, Nationality, Url)
      VALUES((SELECT MAX(ConstructorId) FROM Constructors) + 1, '%s', '%s', '%s',' %s');`,
		values.ConstructorRef, values.Name, values.Nationality, values.Url,
	)
	return s.insert("Constructor", query)
}

func (s *StoreImpl) InsertDriver(values CreateDriverRequest) error {
	query := fmt.Sprintf(`
    INSERT INTO public.Driver (DriverId, Driverref, Number, Code, Forename, Surname, Dob, Nationality)
      VALUES((SELECT MAX(DriverId) FROM Driver) + 1, '%s', '%s', '%s', '%s', '%s', '%s', '%s');`,
		values.DriverRef, values.Number, values.Code, values.Forename, values.Surname, values.DateOfBirth, values.Nationality,
	)
	return s.insert("Drivers", query)
}

// Constructor
type PilotInfo struct {
	Nome           *string `json:"name" db:"nome_completo"`
	DataNascimento *string `json:"data_nascimento" db:"data_nascimento"`
	Nacionalidade  *string `json:"nacionalidade" db:"nacionalidade"`
}

func (s *StoreImpl) SearchPilotByForename(userId int, forename string) (PilotInfo, error) {
	result := []PilotInfo{}
	query := fmt.Sprintf(`
	SELECT y.nome_completo,
			y.data_nascimento,
			y.nacionalidade
		FROM
		(SELECT DISTINCT driverid
		FROM results
		WHERE constructorid = (select idoriginal from users where userid = %d)) x
		INNER JOIN
		(SELECT driverid,
			CONCAT(forename, ' ', surname) AS nome_completo,
			dob AS data_nascimento,
			nationality AS nacionalidade
		FROM driver
		WHERE forename='%s') y ON x.driverid = y.driverid;
	`, userId, forename)

	fmt.Printf("Search Pilot query: %s\n", query)

	if err := s.DB.Select(&result, query); err != nil {
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
	var pilotos []int
	var escuderias []int
	var corridas []int
	var temporadas []int
	var err error

	err = s.DB.Select(&pilotos, fmt.Sprintf(`select count(*) from driver`))
	if err != nil {
		log.Printf("Failed to get #drivers info for Admin: %s\n", err.Error())
		return nil, err
	}
	err = s.DB.Select(&escuderias, fmt.Sprintf(`select count(*) from constructors`))
	if err != nil {
		log.Printf("Failed to get #escuderias info for Admin: %s\n", err.Error())
		return nil, err
	}
	err = s.DB.Select(&corridas, fmt.Sprintf(`select count(*) from races`))
	if err != nil {
		log.Printf("Failed to get #corridas info for Admin: %s\n", err.Error())
		return nil, err
	}
	err = s.DB.Select(&temporadas, fmt.Sprintf(`select count(*) from seasons`))
	if err != nil {
		log.Printf("Failed to get #temporadas info for Admin: %s\n", err.Error())
		return nil, err
	}
	return &AdminOverviewInfo{
		Pilotos:    pilotos[0],
		Escuderias: escuderias[0],
		Corridas:   corridas[0],
		Temporadas: temporadas[0],
	}, nil
}

// Constructor
type ConstructorOverviewInfo struct {
	Vitorias      int    `json:"Vitorias"`
	PilotosUnicos int    `json:"PilotosUnicos"`
	PrimeiroAno   string `json:"PrimerioAno"`
	UltimoAno     string `json:"UltimoAno"`
}

func (s *StoreImpl) GetConstructorOverviewInfo(userId int) (*ConstructorOverviewInfo, error) {
	var err error
	var vitorias []int
	var pilotos []int
	var anos []string

	err = s.DB.Select(&vitorias, fmt.Sprintf(`SELECT encontrar_quantidade_vitorias_escuderia((select idoriginal from users where userid = %d))`, userId))
	if err != nil {
		log.Printf("Failed to get encontrar_quantidade_vitorias_escuderia info for Constructor of userId %d: %s\n", userId, err.Error())
		return nil, err
	}
	err = s.DB.Select(&pilotos, fmt.Sprintf(`SELECT encontrar_quantidade_pilotos_escuderia((select idoriginal from users where userid = %d))`, userId))
	if err != nil {
		log.Printf("Failed to get encontrar_quantidade_pilotos_escuderia info for Constructor of userId %d: %s\n", userId, err.Error())
		return nil, err
	}
	err = s.DB.Select(&anos, fmt.Sprintf(`SELECT encontrar_primeiro_ultimo_ano_escuderia((select idoriginal from users where userid = %d))`, userId))
	if err != nil {
		log.Printf("Failed to get encontrar_primeiro_ultimo_ano_escuderia info for Constructor of userId %d: %s\n", userId, err.Error())
		return nil, err
	}

	re, err := regexp.Compile(`[()]`)
	if err != nil {
		log.Fatal(err)
	}
	anos2 := strings.Split(re.ReplaceAllString(anos[0], ""), ",")

	return &ConstructorOverviewInfo{
		Vitorias:      vitorias[0],
		PilotosUnicos: pilotos[0],
		PrimeiroAno:   anos2[0],
		UltimoAno:     anos2[1],
	}, nil
}

// Driver
type DriverOverviewInfo struct {
	Vitorias    int
	PrimeiroAno string
	UltimoAno   string
}

func (s *StoreImpl) GetDriverOverviewInfo(userId int) (*DriverOverviewInfo, error) {
	var err error
	var vitorias []int
	var anos []string

	err = s.DB.Select(&vitorias, fmt.Sprintf(`SELECT encontrar_quantidade_vitorias_piloto((select idoriginal from users where userid = %d))`, userId))
	if err != nil {
		log.Printf("Failed to get encontrar_quantidade_vitorias_piloto info for Driver of userId %d: %s\n", userId, err.Error())
		return nil, err
	}
	err = s.DB.Select(&anos, fmt.Sprintf(`SELECT encontrar_primeiro_ultimo_ano_piloto((select idoriginal from users where userid = %d))`, userId))
	if err != nil {
		log.Printf("Failed to get encontrar_primeiro_ultimo_ano_piloto info for Driver of userId %d: %s\n", userId, err.Error())
		return nil, err
	}

	re, err := regexp.Compile(`[()]`)
	if err != nil {
		log.Fatal(err)
	}
	anos2 := strings.Split(re.ReplaceAllString(anos[0], ""), ",")

	return &DriverOverviewInfo{
		Vitorias:    vitorias[0],
		PrimeiroAno: anos2[0],
		UltimoAno:   anos2[1],
	}, nil
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
				WHERE constructors.constructorid = (select idoriginal from users where userid = %d)
				GROUP BY status.statusid, constructors.constructorid
				ORDER BY count DESC
		`, userId)
	case "piloto":
		query = fmt.Sprintf(`
			SELECT status.status, COUNT(results.statusid)
				FROM results INNER JOIN status ON results.statusid = status.statusid INNER JOIN driver ON driver.driverid = results.driverid
				WHERE driver.driverid = (select idoriginal from users where userid = %d)
				GROUP BY status.statusid, driver.driverid
				ORDER BY count DESC
		`, userId)
	default:
		return nil, fmt.Errorf("GetResultsByEachStatus default case")
	}

	fmt.Printf("Status Report query: %s\n", query)
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
			FROM (SELECT name, lat, long FROM geocities15k WHERE name='%s') c
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

func (s *StoreImpl) GetConstructorReport(userId int) ([]ConstructorReport, error) {
	report := []ConstructorReport{}
	query := fmt.Sprintf(`
		SELECT nome_completo,
		COALESCE(vitorias, 0) AS vitorias
		FROM
		(SELECT p.driverid,
			nome_completo
		FROM
		(SELECT DISTINCT driverid
		FROM results
		WHERE constructorid=(select idoriginal from users where userid = %d)) p

		LEFT JOIN
		(SELECT driverid,
				CONCAT(forename, ' ', surname) AS nome_completo
		FROM driver) n 
		ON p.driverid = n.driverid) parcial
		LEFT JOIN
		(SELECT driverid,
			count(*) AS vitorias
		FROM results
		WHERE constructorid=(select idoriginal from users where userid = %d)
		AND POSITION=1
		GROUP BY driverid) AS v
		ON parcial.driverid = v.driverid;
	`, userId, userId)

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

func (s *StoreImpl) GetDriverReport(userId int) ([]DriverReport, error) {
	report := []DriverReport{}
	query := fmt.Sprintf(`
		SELECT YEAR AS ano, name AS corrida, SUM(POSITION) AS vitoria
			FROM (SELECT POSITION, driverid, raceid FROM results WHERE driverid=(select idoriginal from users where userid = %d) AND POSITION=1) x
			LEFT JOIN (SELECT YEAR, raceid, name FROM races) y ON x.raceid = y.raceid
			GROUP BY ROLLUP(ano, corrida)
			ORDER BY ano, corrida;
	`, userId)

	fmt.Printf("Driver Report query: %s\n", query)
	if err := s.DB.Select(&report, query); err != nil {
		log.Printf("Error getting Driver Report data: %s\n", err.Error())
		return nil, err
	}
	return report, nil
}
