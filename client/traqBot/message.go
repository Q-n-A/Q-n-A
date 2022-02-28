package traqBot

import (
	"github.com/antihax/optional"
	"github.com/google/uuid"
	"github.com/sapphi-red/go-traq"
)

// SendLog ログを送信
func (c *Client) SendLog(msg string) error {
	// リクエスト用オプション生成
	req := traq.PostMessageRequest{
		Content: msg,
		Embed:   true,
	}
	opts := &traq.MessageApiPostMessageOpts{
		PostMessageRequest: optional.NewInterface(req),
	}

	// メッセージを送信
	_, _, err := c.cli.MessageApi.PostMessage(c.getAuthContext(), c.logChannel, opts)
	if err != nil {
		return err
	}

	return nil
}

// SendMessage 指定したチャンネルにメッセージを送信
func (c *Client) SendMessage(channelID uuid.UUID, msg string) error {
	// リクエスト用オプション生成
	req := traq.PostMessageRequest{
		Content: msg,
		Embed:   true,
	}
	opts := &traq.MessageApiPostMessageOpts{
		PostMessageRequest: optional.NewInterface(req),
	}

	// メッセージを送信
	_, _, err := c.cli.MessageApi.PostMessage(c.getAuthContext(), channelID.String(), opts)
	if err != nil {
		return err
	}

	return nil
}
