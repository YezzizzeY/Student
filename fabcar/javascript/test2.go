package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func login(c echo.Context) error {

	u := new(User)
  	_ = u
  	return c.String(http.StatusOK, "ok")
}
type User struct {
	account  string `json:"account " form:"account " query:"account "`
	password string `json:"password" form:"password" query:"password"`
}
func main() {
	e := echo.New()
	e.GET("/login", login) 
	e.Logger.Fatal(e.Start(":8080"))
}