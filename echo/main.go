package main

import (
	"os"
	"time"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
	// "github.com/gin-gonic/gin"
)

func init() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

var port = os.Getenv("PORT")

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "static/")
	e.GET("/", func(c echo.Context) error {
		// log.Info().Str("path", "/").Str("method", "GET").Send()
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/", func(c echo.Context) error {
		// log.Info().Str("path", "/").Str("method", "POST").Send()
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/test", func(c echo.Context) error {
		// log.Info().Str("path", "/test").Str("method", "GET").Send()
		return c.String(http.StatusOK, "test ok")
	})
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
