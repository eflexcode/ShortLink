package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/internal/config"
	_ "github.com/lib/pq"
)

type DatabaseRepo struct {
	db *sql.DB
}

func NewDatabaseRepo(db *sql.DB) *DatabaseRepo {

	databaseRepo := DatabaseRepo{
		db: db,
	}

	return &databaseRepo
}

func ConnectDb(config config.DatabaseConfig) (*sql.DB, error) {

	db, err := sql.Open(config.DbType, config.Addr)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(config.MaxIdealConn)
	db.SetMaxOpenConns(config.MaxOpenConn)
	duration, err := time.ParseDuration(config.MaxIdealTime)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
