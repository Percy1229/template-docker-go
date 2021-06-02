package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	// Name  string `json:"name"`
	// Email string `json:"email"`
	ID string `json:"id"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))

}
