package cmd

import (
	"github.com/Q-n-A/Q-n-A/util/profiler"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	gRPCAddr string
	restAddr string
	devMode  bool
)

// Serveコマンド - Q'n'Aサーバーの起動
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Q'n'A server",
	Run: func(cmd *cobra.Command, args []string) {
		// フラグによる設定の上書き
		if gRPCAddr != "" {
			cfg.Server.GRPCAddr = gRPCAddr
		}
		if restAddr != "" {
			cfg.Server.RESTAddr = restAddr
		}
		if devMode {
			cfg.DevMode = true
		}

		// wireを使ってサーバーを生成
		s, err := setupServer(cfg, zapLog)
		if err != nil {
			zapLog.Panic("failed to setup server", zap.Error(err))
		}

		// DevModeがtrueならfgprofサーバーを起動
		if cfg.DevMode {
			go profiler.StartFgprof(zapLog)
		}

		// サーバーを起動
		s.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&gRPCAddr, "grpc_port", "g", "", "gRPC address to listen")
	serveCmd.Flags().StringVarP(&restAddr, "rest_port", "r", "", "REST API address to listen")
	serveCmd.Flags().BoolVarP(&devMode, "dev", "d", false, "Development mode")
}
