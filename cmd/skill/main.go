package main

import (
	"net/http"

	"github.com/Nexadis/alice-skill/internal/api"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := Run(); err != nil {
		panic(err)
	}
}

func Run() error {
	e := newServer()
	return e.Start(":8080")
}

func newServer() *echo.Echo {
	e := echo.New()
	e.POST("/", root)
	return e
}

func root(e echo.Context) error {
	api := &api.API{
		Response: &api.Response{
			Text: api.CanNothing,
		},
		Version: api.Version,
	}

	return e.JSON(http.StatusOK, api)
}
