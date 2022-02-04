package profiler

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/felixge/fgprof"
	"go.uber.org/zap"
)

// fgprofサーバーを起動
func StartFgprof(logger *zap.Logger) error {
	// ハンドラを登録
	http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())

	// サーバーを起動
	logger.Info("Starting fgprof server")
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		return err
	}

	return nil
}
