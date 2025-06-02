package config

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type AppConfig struct {
	DB *sqlx.DB
}

func NewAppConfig() *AppConfig {
	log.Println("🚀 [Init] Creating new DB client...")

	db, err := NewDBClient()
	if err != nil {
		log.Fatalf("❌ [Error] Failed to create DB client: %v", err)
	}

	log.Println("✅ [Init] DB client successfully created!")

	return &AppConfig{
		DB: db,
	}
}
