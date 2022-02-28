package cmd

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Q-n-A/Q-n-A/server/protobuf"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

// healthcheckCmd サーバーへのpingが正常に帰ってくるかを確認する
var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Server healthcheck",
	Run: func(cmd *cobra.Command, args []string) {
		// フラグによる設定の上書き
		if gRPCAddr != "" {
			cfg.Server.GRPCAddr = gRPCAddr
		}
		if restAddr != "" {
			cfg.Server.RESTAddr = restAddr
		}

		// gRPCコネクションを作成
		conn, err := grpc.Dial(cfg.Server.GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Panicf("failed to dial too gRPC server: %v", err)
		}
		defer func() { _ = conn.Close() }()

		// Pingサービスのクライアントを作成
		pingClient := protobuf.NewPingClient(conn)

		// Pingメソッドを呼び出し
		res, err := pingClient.Ping(context.Background(), &emptypb.Empty{})
		if err != nil {
			log.Panicf("failed to ping gRPC server: %v", err)
		}
		if res.GetMessage() != "pong" {
			log.Panicf("unexpected response from gRPC server: %v", err)
		}

		// REST APIサーバーの`/api/ping`にGETリクエストを送信
		res2, err := http.DefaultClient.Get(fmt.Sprintf("http://%s/api/ping", cfg.Server.RESTAddr))
		if err != nil {
			log.Panicf("failed to ping REST API server: %v", err)
		}
		if res2.StatusCode != http.StatusOK {
			log.Panicf("unexpected status code: %d", res2.StatusCode)
		}
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(res2.Body)
		defer res2.Body.Close()
		if err != nil {
			log.Panicf("failed to unmarshal REST API response: %v", err)
		}
		if buf.String() != "pong" {
			log.Panicf("unexpected response from REST API server: %s", buf.String())
		}

		log.Println("Healthcheck OK")
	},
}

func init() {
	rootCmd.AddCommand(healthcheckCmd)
	healthcheckCmd.Flags().StringVarP(&gRPCAddr, "grpc_port", "g", "", "gRPC address to dial")
	healthcheckCmd.Flags().StringVarP(&restAddr, "rest_port", "r", "", "REST API address to dial")
}
