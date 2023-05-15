package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
	SSLMode      string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.DatabaseName,
		cfg.Password,
		cfg.SSLMode,
	))
	if err != nil {
		return nil, fmt.Errorf("error occured while initializing db: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err.Error())
	}

	return db, nil
}
