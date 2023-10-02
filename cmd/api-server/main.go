package main

import (
	"flag"
	"github.com/gorilla/mux"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"web-server-yt/internal/apiserver"
	"web-server-yt/internal/slogger"

	"github.com/joho/godotenv"
)

var (
	configPath = ""
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Cant load .env file, err: %s", err)
	}

	flag.StringVar(&configPath, "config-path", os.Getenv("CONFIG_PATH"), "path to config file!")
}

func main() {
	flag.Parse()

	cfg := apiserver.MustLoadConfig(&configPath)
	logger := slogger.New(cfg.Env)
	s := apiserver.New(cfg)

	_ = s
	_ = logger

	router := mux.NewRouter()
	router.HandleFunc("/hello", tester(logger))

	_ = s.Start(router)
}

func tester(logger *slog.Logger) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Somebody come in!")
		_, _ = io.WriteString(w, "hi!")
	}
}
