package logger

import (
	"context"

	"github.com/antihax/optional"
	traq "github.com/sapphi-red/go-traq"
	"go.uber.org/zap/zapcore"
)

// traQログ投稿フック
type traQHook struct {
	cli     *traq.APIClient
	auth    context.Context
	channel string
}

// traQHookを生成
func newTraQHook(accessToken string, channel string) *traQHook {
	// traQクライアントの生成
	client := traq.NewAPIClient(traq.NewConfiguration())
	// アクセストークンから認証情報を生成
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, accessToken)

	return &traQHook{
		cli:     client,
		auth:    auth,
		channel: channel,
	}
}

// フックの対象レベル
var fireLevels = []zapcore.Level{
	zapcore.ErrorLevel,
	zapcore.FatalLevel,
	zapcore.PanicLevel,
}

// フックを実行
func (h *traQHook) Fire(e zapcore.Entry) error {
	// 対象レベルに一致する物がある場合ログを送信
	for _, fireLevel := range fireLevels {
		if e.Level == fireLevel {
			return h.sendLog(e)
		}
	}

	return nil
}

// traQにログを送信
func (h *traQHook) sendLog(e zapcore.Entry) error {
	// ログのフォーマット
	msg := "## " + e.Level.CapitalString() + " log\n" + e.Message + "\n" + e.Time.Format("2006-01-02T15:04:05+MST") + "\n```\n" + e.Stack + "\n```"

	// リクエスト用オプション生成
	req := traq.PostMessageRequest{
		Content: msg,
		Embed:   true,
	}
	opts := &traq.MessageApiPostMessageOpts{
		PostMessageRequest: optional.NewInterface(req),
	}

	// メッセージを送信
	_, _, err := h.cli.MessageApi.PostMessage(h.auth, h.channel, opts)
	if err != nil {
		return err
	}

	return nil
}
