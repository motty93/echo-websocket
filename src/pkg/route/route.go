package route

import (
	"app/pkg/handler"

	"app/pkg/db/validation"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Router is echo REST routing
func Router() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(handler.BodyDumper))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost:8000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Validator = &validation.Custom{Validator: validator.New()}

	e.GET("/", handler.RootHandler)
	e.GET("/test", handler.TestHandler)
	e.GET("/ws", handler.Websocket)

	// us := service.NewServiceUser(db.DB)
	// uh := handler.NewUserHandler(us)
	return e
}
