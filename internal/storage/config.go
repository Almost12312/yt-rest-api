package storage

type Config struct {
	DatabaseUrl string `env:"DATABASE_URL" yaml:"url"`
}
