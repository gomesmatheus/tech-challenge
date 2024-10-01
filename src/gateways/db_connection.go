package gateways

import (
	"github.com/jackc/pgx/v5"
	"github.com/go-redis/redis/v8"
)

type DbConnection struct {
	Db *pgx.Conn
	Redis *redis.Client
}
