package cmd

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/felixge/fgprof"
	"github.com/spf13/cobra"
)

var (
	gRPCAddr string = ""
	restAddr string = ""
	devMode  bool   = false
)

// Serveコマンド - Q'n'Aサーバーの起動
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Q'n'A server",
	Run: func(cmd *cobra.Command, args []string) {
		// フラグによる設定の上書き
		if gRPCAddr != "" {
			cfg.GRPCAddr = gRPCAddr
		}
		if restAddr != "" {
			cfg.RESTAddr = restAddr
		}
		if devMode {
			cfg.DevMode = true
		}

		// wireを使ってサーバーを生成
		s, err := setupServer(cfg)
		if err != nil {
			log.Panicf("failed to setup server: %v", err)
		}

		// DevModeがtrueならfgprofサーバーを起動
		if cfg.DevMode {
			go func() {
				// ハンドラを登録
				http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())

				// サーバーを起動
				log.Print("Starting fgprof server")
				err := http.ListenAndServe(":6060", nil)
				if err != nil {
					log.Panicf("failed to start fgprof server: %v", err)
				}
			}()
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
