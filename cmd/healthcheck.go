package cmd

import (
	"context"
	"fmt"
	"log"

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
			log.Panicf("unexpected response from gRPC server: %v", res)
		}

		fmt.Println("Healthcheck OK")
	},
}

func init() {
	rootCmd.AddCommand(healthcheckCmd)
}
