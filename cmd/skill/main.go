package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	if err := Run(); err != nil {
		panic(err)
	}
}

func Run() error {
	e := echo.New()
	e.POST("/", root)
	return e.Start(":8080")
}

func root(e echo.Context) error {
	api := &API{
		Response: &Response{
			Text: "Извините, я пока ничего не умею",
		},
		Version: "1.0",
	}
	return e.JSON(http.StatusOK, api)
}

type Response struct {
	Text string `json:"text"`
}

type API struct {
	Response *Response `json:"response,omitempty"`
	Version  string    `json:"version"`
}
