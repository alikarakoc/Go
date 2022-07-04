package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Printf("Hello world!")
	e := echo.New()

	e.GET("/main", func(c echo.Context) error {
		return c.String(http.StatusOK, "main endpointine get isteği yapıldı")
	})

	e.Start(":8000")
}
