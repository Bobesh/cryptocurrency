package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/bridge"
)

type EchoApi struct {
	echo *echo.Echo
	app  bridge.App
	port string
}

func NewEchoApi(app bridge.App, port string) *EchoApi {
	return &EchoApi{
		echo: echo.New(),
		app:  app,
		port: port,
	}
}

func (e *EchoApi) Serve() {
	// endpoints
	e.echo.GET("/currency", e.getCryptocurrency)
	e.echo.GET("/find_hash", e.findHash)

	// run
	e.echo.Logger.Fatal(e.echo.Start(fmt.Sprintf(":%s", e.port)))
}

func respond(c echo.Context, res interface{}, err error) error {
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (e *EchoApi) getCryptocurrency(c echo.Context) error {
	currency := c.FormValue("currency")
	app := e.app
	res, err := app.GetCryptoCurrencies(currency)
	return respond(c, res, err)
}

func (e *EchoApi) findHash(c echo.Context) error {
	hashStr := c.FormValue("hash")
	app := e.app
	res, err := app.FindHash(hashStr)
	return respond(c, res, err)
}
