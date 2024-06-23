package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server interface {
	DaysLeft() int
}

type Endpoints struct {
	// s *service.Server // лучше создать интерфейс
	service Server
}

func New(service Server) *Endpoints {
	return &Endpoints{
		service: service,
	}
}

func (e *Endpoints) Status(ctx echo.Context) error {

	d := e.service.DaysLeft()

	s := fmt.Sprintf("Количество днейЖ %d", d)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}
