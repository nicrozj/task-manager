package config

import (
	"github.com/jmoiron/sqlx"
)

type AppConfig struct {
	DB *sqlx.DB
}

func NewAppConfig() *AppConfig {
	db, err := NewDBClient()
	if err != nil {
		panic("failed to get db client")
	}
	return &AppConfig{
		DB: db,
	}
}
