package main

import (
	"flag"
	"log"
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
}
