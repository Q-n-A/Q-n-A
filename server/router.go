package server

import (
	"database/sql"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/srinathgs/mysqlstore"
	"google.golang.org/grpc"
)

func newEcho(sessions.Store) *echo.Echo {
	e := echo.New()
	// ログの設定
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${time_rfc3339} ${prefix} ${short_file} ${line} |")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "${time_rfc3339} method = ${method} | uri = ${uri} | status = ${status} ${error}\n"}))

	// セッションの設定
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret-key"))))

	return e
}

func newMySQLStore(db *sql.DB) (sessions.Store, error) {
	store, err := mysqlstore.NewMySQLStoreFromConnection(db, "sessions", "/", 60*60*24*14, []byte("secret-key"))
	if err != nil {
		return nil, err
	}

	return store, nil
}

func setupHandlers(e *echo.Echo, s *grpc.Server) {
	api := e.Group("/api")
	{
		api.GET("/ping", func(c echo.Context) error {
			return c.String(200, "pong")
		})
	}

	e.GET("/grpc", echo.WrapHandler(convertGRPCServer(s)))
}
