package main

import (
	"net/http"

	"github.com/Nexadis/alice-skill/internal/api"
	"github.com/labstack/echo/v4"
)

func main() {
	c := NewConfig()
	if err := Run(c); err != nil {
		panic(err)
	}
}

func Run(c *Config) error {
	e := newServer()
	return e.Start(c.ListenAddr)
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
