package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"temperature-sensor-receiver/internal/config"
)

func NewDb() (context.Context, *pgxpool.Pool, error) {
	ctx := context.Background()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Host, config.Config.DB.Port, config.Config.DB.Dbname)

	pool, err := pgxpool.New(ctx, connStr)

	return ctx, pool, err
}
