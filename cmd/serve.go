package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	gRPCPort int = 0
	restPort int = 0
)

// Serveコマンド - Q'n'Aサーバーの起動
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Q'n'A server",
	Run: func(cmd *cobra.Command, args []string) {
		// フラグによる設定の上書き
		if gRPCPort != 0 {
			cfg.GRPCPort = gRPCPort
		}
		if restPort != 0 {
			cfg.RESTPort = restPort
		}

		// ロガーを生成
		logger, _ := zap.NewProduction()

		// wireを使ってサーバーを生成
		s, err := SetupServer(cfg, logger)
		if err != nil {
			log.Panicf("failed to setup gRPC server: %v", err)
		}

		// サーバーを起動
		s.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&gRPCPort, "grpc_port", "g", 0, "gRPC Port to listen")
	serveCmd.Flags().IntVarP(&restPort, "rest_port", "r", 0, "REST API Port to listen")
}
