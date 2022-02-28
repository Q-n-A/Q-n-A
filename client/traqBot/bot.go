package traqBot

import (
	"context"

	"github.com/sapphi-red/go-traq"
)

// Client traQ Botクライアント
type Client struct {
	cli                 *traq.APIClient
	accessToken         string
	logChannel          string
	notificationChannel string
}

// Config traQ Botクライアント用設定
type Config struct {
	DevMode             bool
	AccessToken         string
	LogChannel          string
	NotificationChannel string
}

// NewTraQBotClient traQ Botクライアントを生成
func NewTraQBotClient(c *Config) *Client {
	// traQクライアントの生成
	client := traq.NewAPIClient(traq.NewConfiguration())

	return &Client{
		cli:                 client,
		accessToken:         c.AccessToken,
		logChannel:          c.LogChannel,
		notificationChannel: c.NotificationChannel,
	}
}

// getAuthContext アクセストークンから認証情報を生成
func (c *Client) getAuthContext() context.Context {
	return context.WithValue(context.Background(), traq.ContextAccessToken, c.accessToken)
}
