package server

import (
	"database/sql"

	"github.com/brpaz/echozap"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/srinathgs/mysqlstore"
	"go.uber.org/zap"
)

// 新しいEchoインスタンスを生成
func newEcho(store sessions.Store, logger *zap.Logger) *echo.Echo {
	e := echo.New()

	// Echo起動時のログを無効化
	e.HideBanner = true
	e.HidePort = true

	// loggerの設定
	e.Use(echozap.ZapLogger(logger))

	// セッションの設定
	e.Use(session.Middleware(store))

	// ハンドラを登録
	registerHandlers(e)

	return e
}

// ハンドラをEchoインスタンスに登録
func registerHandlers(e *echo.Echo) {
	api := e.Group("/api")
	{
		api.GET("/ping", func(c echo.Context) error {
			return c.String(200, "pong")
		})
	}
}

// 新しいMySQLセッションストアを生成
func newMySQLStore(db *sql.DB) (sessions.Store, error) {
	store, err := mysqlstore.NewMySQLStoreFromConnection(db, "sessions", "/", 60*60*24*14, []byte("secret-key"))
	if err != nil {
		return nil, err
	}

	return store, nil
}
