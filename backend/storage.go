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
		  fmt.Println("User founded")
			return &user[0], nil
		} else {
			return nil, fmt.Errorf("User not found")
		}
}
