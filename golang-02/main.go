package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Main endpointine GET isteği yapıldı.")
}
func userHandler(c echo.Context) error {

	username := c.QueryParam("username")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")

	fmt.Println(username)
	fmt.Println(name)
	fmt.Println(surname)

	return c.String(http.StatusOK, "User endpointine GET isteği yapıldı.")
}
func main() {
	fmt.Printf("Hello World!")

	e := echo.New()

	e.GET("/main", mainHandler)
	e.GET("/user", userHandler)

	e.Start(":8000")
}
