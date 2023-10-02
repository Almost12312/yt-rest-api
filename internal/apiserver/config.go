package apiserver

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

func init() {

}

type Config struct {
	BindAddr string `yaml:"bind_addr"`
	Env      string `env:"ENV" yaml:"env" env=required:"true"`
}

type Env struct {
	ConfigPath string `env:"CONFIG_PATH" env-required:"true"`
}

func MustLoadConfig(path *string) *Config {
	path = mustValidatePath(path)

	var cfg Config
	if err := cleanenv.ReadConfig(*path, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", path)
	}

	return &cfg
}

func mustValidatePath(path *string) *string {
	if *path == "" {
		log.Fatal("Parse CONFIG_PATH")
	}

	if _, err := os.Stat(*path); os.IsNotExist(err) {
		log.Fatalf("File not found: %s", path)
	}

	return path
}
