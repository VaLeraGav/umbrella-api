package app

import (
	"fmt"
	"log"
	"umbrella/internal/app/endpoint"
	mv "umbrella/internal/app/mw"
	"umbrella/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	endpoint *endpoint.Endpoints
	service  *service.Service
	echo     *echo.Echo
}

func New() (*App, error) {
	s := service.New()
	e := endpoint.New(s)

	a := &App{
		endpoint: e,
		service:  s,
		echo:     echo.New(),
	}

	a.echo.Use(mv.RoleCheck)

	a.echo.GET("/status", a.endpoint.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server running")

	err := a.echo.Start(":1323")
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("failed to start http server: %w", err)
	}

	return nil
}
