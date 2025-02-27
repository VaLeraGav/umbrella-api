package mv

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	roleAdmin = "admin"
	roleUser  = "user"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")

		if strings.Contains(val, roleAdmin) {
			log.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
