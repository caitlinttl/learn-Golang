package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/json", func(c echo.Context) error {
		return c.JSONBlob(
			http.StatusOK,
			[]byte(`{ "id": "1", "msg": "Hello, Boatswain!" }`),
		)
	})

	e.GET("/html", func(c echo.Context) error {
		return c.HTML(
			http.StatusOK,
			"<h1>Hello, Boatswain!</h1>",
		)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
