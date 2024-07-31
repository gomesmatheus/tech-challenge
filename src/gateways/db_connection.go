package gateways

import (
	"github.com/jackc/pgx/v5"
)

type DbConnection struct {
	Db *pgx.Conn
}
