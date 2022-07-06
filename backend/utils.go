package main

import (
	"fmt"
  "os"
  "strings"
)

func selectProperUserPage(user *User) string{
	switch user.Tipo {
	case "Admin":
	  return "/admin"
	case "Escuderia":
	  return "/escuderia"
	case "Piloto":
	  return "/piloto"
	default:
	  return "/404.html"
	}
}

func renderHTML(filename string, nome map[string]string) string {
  html := readFile(filename)
  for k,v := range nome{
	  html = strings.ReplaceAll(html, "{{"+k+"}}", v)
  }

  return html
}

func checkFile(e error) {
	if e != nil {
		panic(e)
	}
}
  
func readFile(file string) string{
	dat, err := os.ReadFile(file)
	checkFile(err)
	fmt.Print(string(dat))
	
  return string(dat)
}