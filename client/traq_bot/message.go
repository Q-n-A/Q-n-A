package traq_bot

import (
	"github.com/antihax/optional"
	"github.com/google/uuid"
	"github.com/sapphi-red/go-traq"
)

func (c *TraQBotClient) SendLog(msg string) error {
	// リクエスト用オプション生成
	req := traq.PostMessageRequest{
		Content: msg,
		Embed:   true,
	}
	opts := &traq.MessageApiPostMessageOpts{
		PostMessageRequest: optional.NewInterface(req),
	}

	// メッセージを送信
	_, _, err := c.cli.MessageApi.PostMessage(c.auth, c.logChannel, opts)
	if err != nil {
		return err
	}

	return nil
}

func (c *TraQBotClient) SendMessage(channelID uuid.UUID, msg string) error {
	// リクエスト用オプション生成
	req := traq.PostMessageRequest{
		Content: msg,
		Embed:   true,
	}
	opts := &traq.MessageApiPostMessageOpts{
		PostMessageRequest: optional.NewInterface(req),
	}

	// メッセージを送信
	_, _, err := c.cli.MessageApi.PostMessage(c.auth, channelID.String(), opts)
	if err != nil {
		return err
	}

	return nil
}
