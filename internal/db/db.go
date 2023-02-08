package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	connDB  *sql.DB
	timeout time.Duration
}

func NewDB(dataSourceName string, timeout time.Duration) *DB {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("can't connect to db: %v", err)
	}
	return &DB{
		connDB:  db,
		timeout: timeout,
	}
}

func (db *DB) SaveRequest(statusCode int) {
	ctx, _ := context.WithTimeout(context.Background(), db.timeout)
	_, err := db.connDB.ExecContext(ctx, "insert into requests_history (status_code) values ($1)", statusCode)
	if err != nil {
		log.Fatalf("can't save status code to db: %v", err)
	}
}
