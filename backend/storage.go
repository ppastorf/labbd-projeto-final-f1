package main

import (
	"fmt"
	"crypto/md5"
)

func md5Encrypt(password string) string{
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func (s *StoreImpl) Close() error {
	return s.DB.Close()
}

func (s *StoreImpl) reportAllUsers() ([]User, error) {
  query := `SELECT * FROM users`

  users := []User{}

  if err := s.DB.Select(&users, query); err != nil {
    fmt.Println(err)
    return nil, err
	} else {
    fmt.Println("Returning all users")
    return users, nil
  }
}

func (s *StoreImpl) GetResultsByEachStatus(userId int, tipo string) ([]GetResultsByEachStatus, error) {
  GetResultsByEachStatus := []GetResultsByEachStatus{}

  fmt.Println("Entrei no storage")
  switch tipo {
	case "Admin":
    query := `select status.status, count(results.statusid)
    from results inner join status on results.statusid = status.statusid
    group by status.statusid
    order by count desc`

    if err := s.DB.Select(&GetResultsByEachStatus, query); err != nil {
      fmt.Println(err)
      return nil, err
    } else {
      return GetResultsByEachStatus, nil
    }
	case "Escuderia":
    query := fmt.Sprintf(`select status.status, count(results.statusid) 
    from results inner join status on results.statusid = status.statusid inner join constructors on constructors.constructorid = results.constructorid
    where constructors.constructorid = %v
    group by status.statusid, constructors.constructorid
    order by count desc`, userId)

    if err := s.DB.Select(&GetResultsByEachStatus, query); err != nil {
      fmt.Println(err)
      return nil, err
    } else {
      return GetResultsByEachStatus, nil
    }
	case "PILOTO":
    query := fmt.Sprintf(`select status.status, count(results.statusid)
    from results inner join status on results.statusid = status.statusid inner join driver on driver.driverid = results.driverid
    where driver.driverid = %v
    group by status.statusid, driver.driverid
    order by count desc`, userId)

    if err := s.DB.Select(&GetResultsByEachStatus, query); err != nil {
      fmt.Println(err)
      return nil, err
    } else {
      return GetResultsByEachStatus, nil
    }

	default:
	  return nil, fmt.Errorf("GetResultsByEachStatus default case")
	}
}


func (s *StoreImpl) rawSQL(input InputRawSQL) (interface{}, error) {
  out := []map[string]interface{}{}
  
  rows, _ := s.DB.Queryx(input.SQL)
  for rows.Next() {
    m := map[string]interface{}{}
    err := rows.MapScan(m)
    if err != nil {
      fmt.Println("err", err)
    }
    out = append(out, m)
  }
  return out, nil
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
