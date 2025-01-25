package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/qsmsoft/tg-service/config"
	"log"
	"time"
)

type Database struct {
	Conn *sqlx.DB
}

func NewDatabase(cfg config.Config) *Database {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSslMode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &Database{Conn: db}
}

func (db *Database) Close() {
	if err := db.Conn.Close(); err != nil {
		log.Printf("failed to close database connection: %v", err)
	} else {
		log.Println("database connection closed")
	}
}
