package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var method = r.Method
		switch method {
		case "GET":
			io.WriteString(w, "Hello, World for GET!!")
		case "POST":
			io.WriteString(w, "Hello, World!! for POST")
		default:
			io.WriteString(w, fmt.Sprintf("Unsupported method: %s", r.Method))
		}
	})

	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Error().Err(err)
	}
}
