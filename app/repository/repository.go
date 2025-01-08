package repository

import (
	"context"
	"database/sql"
)

// Execer 行の取得が不要なクエリを実行するためのインターフェース
type Execer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Preparer
}

// Queryer 行を取得するクエリを実行するためのインターフェース
type Queryer interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Preparer
}

// Preparer プリペアードステートメントを生成するインターフェイス
type Preparer interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// QueryExecer クエリを実行するためのインターフェース
type QueryExecer interface {
	Queryer
	Execer
}

type TxBeginner interface {
	// BeginTx トランザクションを開始する
	BeginTx(ctx context.Context) (Tx, error)
}

type ReadOnlyTxBeginner interface {
	// BeginReadOnlyTx 読み込み専用トランザクションを開始する
	BeginReadOnlyTx(ctx context.Context) (Tx, error)
}

type Tx interface {
	QueryExecer
	Commit() error
	Rollback() error
}

// DB データーベース
type DB interface {
	Queryer

	// *sql.DB は Execer を実装しているが、Exec はトランザクションの中でやって欲しいので、インターフェースからは除外
	// Execer

	TxBeginner
	ReadOnlyTxBeginner
}
