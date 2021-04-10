package main

import (
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", home)
	e.GET("/ls", ls)

	e.Logger.Fatal(e.Start(":5000"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func ls(c echo.Context) error {
	args := c.QueryParam("args")

	output, error := exec.Command("ls", args).Output()

	if error != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	return c.String(http.StatusOK, string(output))
}
