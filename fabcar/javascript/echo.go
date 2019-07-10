package main

import (
	"net/http"
	"github.com/labstack/echo"
	"fmt"
	"os/exec"
)

func queryall(e echo.Context) error {
	c := "node querystudents.js"
	cmd := exec.Command("sh","-c", c)
	f,err := cmd.Output()
	if err!=nil {
		fmt.Printf("Error:OUTPUTING CMD FAILED")
	}
	return e.String(http.StatusOK,string(f))
}

func querystudent(e echo.Context) error {
	name := e.QueryParam("name")
	c := "node querystudent.js "+name
	cmd := exec.Command("sh","-c", c)
	f,err := cmd.Output()
	if err!=nil {
		fmt.Printf("Error:OUTPUTING CMD FAILED")
	}
	return e.String(http.StatusOK,string(f))
}

func create(e echo.Context) error {
	id := e.QueryParam("id")

    name := e.QueryParam("name")

   	mark := e.QueryParam("mark")
	
	record := e.QueryParam("model")
	
	cmd := exec.Command("sh","-c",id,name,mark,record)

	t,err := cmd.Output()
	if err!=nil {
		fmt.Printf("Error:OUTPUTING CMD FAILED")
	}
	return e.String(http.StatusOK,string(t))
}

func main() {
	e := echo.New()
	e.GET("/", queryall) 
	e.GET("/querystudent",querystudent)
	e.GET("/create",create)
	e.Logger.Fatal(e.Start(":1323"))
}
