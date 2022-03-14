package gorm2

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ctxKey string // ctxKey contextに格納するDBインスタンスのkey型

const txKey ctxKey = "transaction" // txKey contextに格納するDBインスタンスのkey

var errorDBCastFailed = errors.New("failed to cast DB instance to *gorm.DB") // 型キャスト失敗エラー

// Do Transactionの中でメソッドを実行する
func (repo *Repository) Do(ctx context.Context, options *sql.TxOptions, callBack func(context.Context) error) error {
	// トランザクション内で実行される関数
	// 返り値がnilでなければロールバックされ、返り値がnilならコミットされる
	txFunc := func(tx *gorm.DB) error {
		// contextにDBインスタンスを格納
		ctx = context.WithValue(ctx, txKey, tx)

		// 引数で渡されたコールバック処理を実行
		err := callBack(ctx)
		if err != nil {
			return err
		}

		// contextがキャンセルされていないかチェック
		err = ctx.Err()
		if err != nil {
			return err
		}

		return nil
	}

	// transactionを作成し、txFuncを実行
	if options == nil {
		err := repo.db.Transaction(txFunc)
		if err != nil {
			return fmt.Errorf("failed to execute callback: %w", err)
		}
	} else {
		err := repo.db.Transaction(txFunc, options)
		if err != nil {
			return fmt.Errorf("failed to execute callback:%w", err)
		}
	}

	return nil
}

// getTX DBインスタンスをコンテキストから取得
func (repo *Repository) getTX(ctx context.Context) (*gorm.DB, error) {
	// contextからDBインスタンスを取得
	txInterface := ctx.Value(txKey)
	// contextにDBインスタンスが存在しない場合はcontextをもとに新たなセッションを開始
	if txInterface == nil {
		return repo.db.WithContext(ctx), nil
	}

	// DBインスタンスをgorm.DB型にキャスト
	tx, ok := txInterface.(*gorm.DB)
	if !ok {
		return nil, errorDBCastFailed
	}

	return tx, nil
}
