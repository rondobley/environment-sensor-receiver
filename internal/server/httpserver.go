package server

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type httpServer struct {
	Ctx context.Context
	Db  *pgxpool.Pool
}
