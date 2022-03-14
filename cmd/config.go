package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Q-n-A/Q-n-A/client/traqBot"
	"github.com/Q-n-A/Q-n-A/repository/gorm2"
	"github.com/Q-n-A/Q-n-A/server"
	"github.com/Q-n-A/Q-n-A/util/logger"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cfg 設定格納用変数
var cfg *config

// config サーバー設定
type config struct {
	DevMode bool `mapstructure:"dev_mode"` // 開発モード (default: false)

	Bot struct {
		AccessToken         string `mapstructure:"access_token"`         // Bot用アクセストークン (default: "")
		VerificationToken   string `mapstructure:"verification_token"`   // Bot用確認トークン (default: "")
		LogChannel          string `mapstructure:"log_channel"`          // ログ投稿チャンネル (default: "")
		NotificationChannel string `mapstructure:"notification_channel"` // 通知投稿チャンネル (default: "")
	} `mapstructure:"bot"` // Bot用設定

	OAuth struct {
		ID    string `mapstructure:"id"`     // 本番環境向けのクライアントID (default: "")
		DevID string `mapstructure:"dev_id"` // ローカル開発環境向けのクライアントID (default: "")
	} `mapstructure:"oauth"` // OAuthクライアント用設定

	Server struct {
		GRPCAddr string `mapstructure:"grpc_addr"` // gRPCサーバーがリッスンするアドレス (default: :9001)
		RESTAddr string `mapstructure:"rest_addr"` // REST APIサーバーがリッスンするアドレス (default: :9000)
	} `mapstructure:"server"` // サーバー用設定

	MariaDB struct {
		Hostname string `mapstructure:"hostname"` // DBのホスト (default: "mariadb")
		Port     int    `mapstructure:"port"`     // DBのポート番号 (default: 3306)
		Username string `mapstructure:"username"` // DBのユーザー名 (default: "root")
		Password string `mapstructure:"password"` // DBのパスワード (default: "password")
		Database string `mapstructure:"database"` // DBのDB名 (default: "Q-n-A")
	} `mapstructure:"mariadb"` // MariaDB用設定
}

// provideRepositoryConfig Gorm v2リポジトリ用設定の提供
func provideRepositoryConfig(c *config) *gorm2.Config {
	return &gorm2.Config{
		Hostname: c.MariaDB.Hostname,
		Port:     c.MariaDB.Port,
		Username: c.MariaDB.Username,
		Password: c.MariaDB.Password,
		Database: c.MariaDB.Database,
	}
}

// provideServerConfig サーバー用設定の提供
func provideServerConfig(c *config) *server.Config {
	return &server.Config{
		DevMode:  c.DevMode,
		GRPCAddr: c.Server.GRPCAddr,
		RESTAddr: c.Server.RESTAddr,
	}
}

// provideLoggerConfig logger用設定の提供
func provideLoggerConfig(c *config) *logger.Config {
	return &logger.Config{
		DevMode: c.DevMode,
	}
}

// provideTraQBotClientConfig traQ Botクライアント用設定の提供
func provideTraQBotClientConfig(c *config) *traqBot.Config {
	return &traqBot.Config{
		DevMode:             c.DevMode,
		AccessToken:         c.Bot.AccessToken,
		LogChannel:          c.Bot.LogChannel,
		NotificationChannel: c.Bot.NotificationChannel,
	}
}

// loadConfig 設定を読み込む
func loadConfig(cfgFile string) error {
	// デフォルト値の設定
	viper.SetDefault("dev_mode", false)

	viper.SetDefault("bot.access_token", "")
	viper.SetDefault("bot.verification_token", "")
	viper.SetDefault("bot.log_channel", "")
	viper.SetDefault("bot.notification_channel", "")

	viper.SetDefault("oauth.id", "")
	viper.SetDefault("oauth.dev_id", "")

	viper.SetDefault("server.grpc_addr", ":9001")
	viper.SetDefault("server.rest_addr", ":9000")

	viper.SetDefault("mariadb.hostname", "mariadb")
	viper.SetDefault("mariadb.port", 3306)
	viper.SetDefault("mariadb.username", "root")
	viper.SetDefault("mariadb.password", "password")
	viper.SetDefault("mariadb.database", "Q-n-A")

	// 環境変数の取得
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if cfgFile != "" {
		// 引数で渡された設定ファイルをセット
		viper.SetConfigFile(cfgFile)
	} else {
		// デフォルトの設定ファイルをセット
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	// 設定ファイルの読み込み
	if err := viper.ReadInConfig(); err != nil {
		if ok := errors.As(err, &viper.ConfigFileNotFoundError{}); ok {
			log.Print("Unable to find config file, default settings or environmental variables are to be used.")
		} else {
			return fmt.Errorf("failed to load config file - %w", err)
		}
	}

	// 構造体にバインド
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to parse configs - %w", err)
	}

	return nil
}

// configCmd 現在の設定を表示する
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print current configurations to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current Config:")

		// spewのダンプ設定
		scs := spew.ConfigState{
			Indent:                  "\t",
			DisablePointerAddresses: true,
		}
		// Config構造体の表示
		scs.Dump(cfg)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
