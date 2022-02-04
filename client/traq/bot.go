package traq

import (
	"context"

	"github.com/sapphi-red/go-traq"
)

// traQ Botクライアント
type traQBotClient struct {
	cli                 *traq.APIClient
	auth                context.Context
	notificationChannel string
}

// traQ Botクライアント用設定
type Config struct {
	DevMode             bool
	AccessToken         string
	NotificationChannel string
}

// traQ Botクライアントを生成
func NewTraQBotClient(c *Config) *traQBotClient {
	// traQクライアントの生成
	client := traq.NewAPIClient(traq.NewConfiguration())
	// アクセストークンから認証情報を生成
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, c.AccessToken)

	return &traQBotClient{
		cli:                 client,
		auth:                auth,
		notificationChannel: c.NotificationChannel,
	}
}
