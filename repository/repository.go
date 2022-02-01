package repository

import (
	"context"
	"database/sql"
)

// Repositoryインターフェース
// repositoryとしての基本動作を定義
// 全てのsub-repositoryインターフェースに埋め込まれる
type Repository interface {
	// Transactionの中でメソッドを実行する
	Do(context.Context, *sql.TxOptions, func(context.Context) error) error
}
