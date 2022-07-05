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

func (s *StoreImpl) Login(login string, password string) ([]User, error) {
	
  user := []User{}

  // !!
	query := fmt.Sprintf(`SELECT * FROM users WHERE login='%v' AND password='%v'`, login, md5Encrypt(password))

	if err := s.DB.Select(&user, query); err != nil {
    return nil, err
	} else if len(user) > 0 {
		  fmt.Println("User founded")
			return user, nil
		} else {
			return nil, fmt.Errorf("User not found")
		}
}
