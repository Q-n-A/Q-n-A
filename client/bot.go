package client

import "github.com/google/uuid"

// BotClient Botクライアント
type BotClient interface {
	// ログを送信
	SendLog(msg string) error
	// メッセージを送信
	SendMessage(channelID uuid.UUID, msg string) error
}
