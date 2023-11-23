package main

import (
	"fmt"
	"os"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

var upgrader = websocket.Upgrader{}

func WebsocketHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		return err
	}

	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			return
		}
	}(ws)

	for {
		// WRITE
		err := ws.WriteMessage(websocket.TextMessage, []byte("hola"))

		if err != nil {
			c.Logger().Error(err)
		}

		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("message from client: %s\n", msg)
	}
}

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

	e.GET("/ws", WebsocketHandler)

	port := os.Getenv("WSCHAT_SERVER_PORT")
	if port == "" {
		port = "5000"
	}

	e.Start(":" + port)
}
