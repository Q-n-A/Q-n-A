package logger

import (
	"github.com/Q-n-A/Q-n-A/client"
	"go.uber.org/zap/zapcore"
)

// traQログ投稿フック
type traQHook struct {
	cli client.BotClient
}

// traQHookを生成
func newTraQHook(cli client.BotClient) *traQHook {
	return &traQHook{
		cli: cli,
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

	// メッセージを送信
	err := h.cli.SendLog(msg)
	if err != nil {
		return err
	}

	return nil
}
