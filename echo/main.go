package main

import (
	"os"
	"strconv"
	"time"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

func init() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

var port = os.Getenv("PORT")

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "static/")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "test ok")
	})
	e.GET("/name/:name", func(c echo.Context) error {
		name := c.Param("name")
		age := c.QueryParam("age")
		if age == "" {
			age = "0"
		}
		ageint, _ := strconv.Atoi(age)
		jsonData := Data{Name: name, Age: ageint}
		return c.JSON(http.StatusOK, jsonData)
	})
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
