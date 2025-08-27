package postgres

import (
	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
)

func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)
	return db
}
