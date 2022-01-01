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

// Healthcheckコマンド - サーバーへのpingが正常に帰ってくるかを確認する
var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Server healthcheck",
	Run: func(cmd *cobra.Command, args []string) {
		// フラグによる設定の上書き
		if gRPCPort != 0 {
			cfg.GRPCPort = gRPCPort
		}
		if restPort != 0 {
			cfg.RESTPort = restPort
		}

		// gRPCコネクションを作成
		conn, err := grpc.Dial(fmt.Sprintf(":%d", cfg.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Panicf("failed to dial too gRPC server: %v", err)
		}
		defer conn.Close()

		// Pingサービスのクライアントを作成
		pingClient := protobuf.NewPingClient(conn)

		// Pingメソッドを実行
		res, err := pingClient.Ping(context.Background(), &emptypb.Empty{})
		if err != nil {
			log.Panicf("failed to ping gRPC server: %v", err)
		}
		if res.GetMessage() != "pong" {
			log.Panicf("unexpected response from gRPC server: %v", res.GetMessage())
		}

		// REST APIサーバーの`/api/ping`にGETリクエストを送信
		res2, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:%d/api/ping", cfg.RESTPort))
		if err != nil {
			log.Panicf("failed to ping REST API server: %v", err)
		}
		defer res2.Body.Close()
		if res2.StatusCode != http.StatusOK {
			log.Panicf("unexpected status code: %d", res2.StatusCode)
		}
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(res2.Body)
		if err != nil {
			log.Panicf("failed to unmarshal REST API responce: %v", err)
		}
		if buf.String() != "pong" {
			log.Panicf("unexpected response from REST API server: %s", buf.String())
		}

		fmt.Println("Healthcheck OK")
	},
}

func init() {
	rootCmd.AddCommand(healthcheckCmd)
	healthcheckCmd.Flags().IntVarP(&gRPCPort, "grpc_port", "g", 0, "gRPC Port to dial")
	healthcheckCmd.Flags().IntVarP(&restPort, "rest_port", "r", 0, "REST API Port to dial")
}
