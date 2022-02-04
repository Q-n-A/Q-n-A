package client

import "github.com/google/uuid"

// Botクライアント
type BotClient interface {
	// ログを送信
	SendMessage(channelID uuid.UUID, message string) error
}
