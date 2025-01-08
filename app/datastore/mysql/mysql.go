package mysql

import (
	"backend/app/repository"
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var _ repository.DB = (*DB)(nil)

type DB struct {
	*sql.DB
}

func (db *DB) BeginTx(ctx context.Context) (repository.Tx, error) {
	return db.DB.BeginTx(ctx, nil)
}

func (db *DB) BeginReadOnlyTx(ctx context.Context) (repository.Tx, error) {
	return db.DB.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: true,
	})
}

func New(config *mysql.Config) (*DB, error) {
	conn, err := mysql.NewConnector(config)
	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(conn)

	return &DB{
		DB: db,
	}, nil
}
