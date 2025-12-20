package db

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PgxIface interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

func Connect(log *zap.Logger) (*pgxpool.Pool, error) {

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Error("db URL is empty")
		return nil, errors.New("db URL is empty")
	}

	dbPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Error("cant create db pool")
		return nil, err
	}

	if err := dbPool.Ping(context.Background()); err != nil {
		log.Error("cant ping to db")
		return nil, err
	}

	log.Info("connected to DB", zap.Time("At", time.Now()))
	return dbPool, nil
}
