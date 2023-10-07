package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "we13"
	dbname   = "rest"
)

type Storage struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Storage {
	return &Storage{config: config}
}

func (s *Storage) Open() error {
	const op = "storage.Open"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	//log.Printf("read data from storage %s in op: %s \n", s.config.DatabaseUrl, op)

	for i := 0; i < 10000; i++ {
		if err := db.Ping(); err != nil {
			return err
		}
	}

	s.db = db

	return nil
}

func (s *Storage) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}

	return nil

}
