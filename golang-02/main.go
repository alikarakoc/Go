package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Username string "json:'username'"
	Name     string "json:'name'"
	Surname  string "json:'surname'"
}

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Main endpointine GET isteği yapıldı.")
}
func userHandler(c echo.Context) error {

	dataType := c.Param("data")
	username := c.QueryParam("username")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Username %s, Name: %s, Surname: %s", username, name, surname))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"username": username,
			"name":     name,
			"surname":  surname,
		})
	}

	fmt.Println(username)
	fmt.Println(name)
	fmt.Println(surname)

	return c.String(http.StatusBadRequest, "Yalnızca JSON veya String parametresini kullanabilirsiniz.")
}
func addUserHandler(c echo.Context) error {

	user := User{}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	fmt.Println(user)

	return c.String(http.StatusOK, "Başarılı")
}
func mainAdminHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Admin endpointindesin.")
}
func main() {
	fmt.Printf("Hello World!")

	e := echo.New()
	//e.Use(middleware.Logger()) //klasik kullanım
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	})) //custom kullanım

	e.GET("/main", mainHandler)

	adminGroup := e.Group("/admin", middleware.Logger())
	adminGroup.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		if username == "admin" && password == "123" {
			return true, nil
		}
		return false, nil
	}))
	adminGroup.GET("/main", mainAdminHandler)

	e.GET("/user/:data", userHandler)
	e.POST("/user", addUserHandler)

	e.Start(":8000")
}
