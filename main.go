package main

import (
	"ahmaruff/wschat/user"
	"ahmaruff/wschat/wsservice"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	e := echo.New()

	logger := zerolog.New(os.Stdout)

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().Str("URI", v.URI).Int("Status", v.Status).Msg("request")

			return nil
		},
	}))

	e.Use(middleware.Recover())
	user.InitUserRoutes(e)
	wsservice.InitWsRoutes(e)

	port := os.Getenv("WSCHAT_SERVER_PORT")
	if port == "" {
		port = "5000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
