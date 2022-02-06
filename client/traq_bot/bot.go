package traq_bot

import (
	"context"

	"github.com/sapphi-red/go-traq"
)

// traQ Botクライアント
type TraQBotClient struct {
	cli                 *traq.APIClient
	auth                context.Context
	logChannel          string
	notificationChannel string
}

// traQ Botクライアント用設定
type Config struct {
	DevMode             bool
	AccessToken         string
	LogChannel          string
	NotificationChannel string
}

// traQ Botクライアントを生成
func NewTraQBotClient(c *Config) *TraQBotClient {
	// traQクライアントの生成
	client := traq.NewAPIClient(traq.NewConfiguration())
	// アクセストークンから認証情報を生成
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, c.AccessToken)

	return &TraQBotClient{
		cli:                 client,
		auth:                auth,
		logChannel:          c.LogChannel,
		notificationChannel: c.NotificationChannel,
	}
}
