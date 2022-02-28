package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	gRPCAddr string // gRPCサーバーがリッスンするアドレス
	restAddr string // REST APIサーバーがリッスンするアドレス
	devMode  bool   // 開発モード
)

// serveCmd Q'n'Aサーバーの起動
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
		s, err := setupServer(cfg)
		if err != nil {
			log.Panicf("failed to setup server: %v", err)
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
