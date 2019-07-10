package main

import (
	"net/http"
	"github.com/labstack/echo"
	"fmt"
	"os/exec"
)

func queryall(e echo.Context) error {
	c := "node 20190708queryall.js"
	cmd := exec.Command("sh","-c", c)
	f,err := cmd.Output()
	
	g:=[]rune(string(f))
	h:=string(g[119:])
	
	if err!=nil {
		fmt.Printf("Error:OUTPUTING CMD FAILED")
	}
	
	return e.String(http.StatusOK,h)
}

func querystudent(e echo.Context) error {
	id := e.QueryParam("id")
	c := "node 20190708query.js "+id
	cmd := exec.Command("sh","-c", c)
	f,err := cmd.Output()
	g:=[]rune(string(f))
	h:=string(g[119:])
	if err!=nil {
		fmt.Printf("Error:OUTPUTING CMD FAILED")
	}
	return e.String(http.StatusOK,h)
}

func create(e echo.Context) error {
	id := e.QueryParam("id")
    name := e.QueryParam("name")
	sex := e.QueryParam("sex")
	institute := e.QueryParam("institute")
	major := e.QueryParam("major")
	c:= "node 20190708create.js "+id+ " "+id+" "+name+" "+sex+" "+institute+" "+major+""
	cmd := exec.Command("sh","-c",c)

	t,err := cmd.Output()
	if err!=nil {
		fmt.Printf("Error:OUTPUTING CMD FAILED")
	}
	return e.String(http.StatusOK,string(t))
}

func main() {
	e := echo.New()
	e.GET("/queryall", queryall) 
	e.GET("/querystudent",querystudent)
	e.GET("/create",create)
	e.Logger.Fatal(e.Start(":1323"))
}
