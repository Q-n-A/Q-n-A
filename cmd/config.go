package cmd

import (
	"fmt"
	"log"

	"github.com/Q-n-A/Q-n-A/repository/gorm2"
	"github.com/Q-n-A/Q-n-A/server"
	"github.com/Q-n-A/Q-n-A/util/logger"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 設定格納用変数
var cfg *Config

// 設定
type Config struct {
	DevMode bool `mapstructure:"dev_mode" json:"dev_mode,omitempty"` // 開発モード (default: false)

	Bot struct {
		AccessToken         string `mapstructure:"access_token" json:"access_token,omitempty"`                 // Bot用アクセストークン (default: "")
		VerificationToken   string `mapstructure:"verification_token" json:"verification_token,omitempty"`     // Bot用確認トークン (default: "")
		LogChannel          string `mapstructure:"log_channel" json:"log_channel,omitempty"`                   // ログ投稿チャンネル (default: "")
		NotificationChannel string `mapstructure:"notification_channel" json:"notification_channel,omitempty"` // 通知投稿チャンネル (default: "")
	} `mapstructure:"bot" json:"bot,omitempty"` // Bot用設定

	Client struct {
		ID    string `mapstructure:"id" json:"id,omitempty"`         // 本番環境向けのクライアントID (default: "")
		DevID string `mapstructure:"dev_id" json:"dev_id,omitempty"` // ローカル開発環境向けのクライアントID (default: "")
	} `mapstructure:"client" json:"client,omitempty"` // OAuthクライアント用設定

	Server struct {
		GRPCAddr string `mapstructure:"grpc_addr" json:"grpc_addr,omitempty"` // gRPCサーバーがリッスンするアドレス (default: :9001)
		RESTAddr string `mapstructure:"rest_addr" json:"rest_addr,omitempty"` // REST APIサーバーがリッスンするアドレス (default: :9000)
	} `mapstructure:"server" json:"server,omitempty"` // サーバー用設定

	DB struct {
		Hostname string `mapstructure:"hostname" json:"hostname,omitempty"` // DBのホスト (default: "mariadb")
		Port     int    `mapstructure:"port" json:"port,omitempty"`         // DBのポート番号 (default: 3306)
		Username string `mapstructure:"username" json:"username,omitempty"` // DBのユーザー名 (default: "root")
		Password string `mapstructure:"password" json:"password,omitempty"` // DBのパスワード (default: "password")
		Database string `mapstructure:"database" json:"database,omitempty"` // DBのDB名 (default: "Q-n-A")
	} `mapstructure:"db" json:"db,omitempty"` // DB用設定

}

// Gorm v2リポジトリ用設定の提供
func provideRepositoryConfig(c *Config) *gorm2.Config {
	return &gorm2.Config{
		Hostname: c.DB.Hostname,
		Port:     c.DB.Port,
		Username: c.DB.Username,
		Password: c.DB.Password,
		Database: c.DB.Database,
	}
}

// サーバー用設定の提供
func provideServerConfig(c *Config) *server.Config {
	return &server.Config{
		GRPCAddr: c.Server.GRPCAddr,
		RESTAddr: c.Server.RESTAddr,
	}
}

// logger用設定の提供
func provideLoggerConfig(c *Config) *logger.Config {
	return &logger.Config{
		DevMode:     c.DevMode,
		AccessToken: c.Bot.AccessToken,
		LogChannel:  c.Bot.LogChannel,
	}
}

// 設定を読み込む
func loadConfig(cfgFile string) error {
	// デフォルト値の設定
	viper.SetDefault("dev_mode", false)

	viper.SetDefault("bot.access_token", "")
	viper.SetDefault("bot.verification_token", "")
	viper.SetDefault("bot.log_channel", "")
	viper.SetDefault("bot.notification_channel", "")

	viper.SetDefault("client.id", "")
	viper.SetDefault("client.dev_id", "")

	viper.SetDefault("server.grpc_addr", ":9001")
	viper.SetDefault("server.rest_addr", ":9000")

	viper.SetDefault("db.hostname", "mariadb")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.username", "root")
	viper.SetDefault("db.password", "password")
	viper.SetDefault("db.database", "Q-n-A")

	// 環境変数の取得
	viper.AutomaticEnv()

	if cfgFile != "" {
		// 引数で渡された設定ファイルをセット
		viper.SetConfigFile(cfgFile)
	} else {
		// デフォルトの設定ファイルをセット
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("json")
	}

	// 設定ファイルの読み込み
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Print("Unable to find config file, default settings or environmental variables are to be used.")
		} else {
			return fmt.Errorf("Error: failed to load config file - %s ", err)
		}
	}

	// 構造体にバインド
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return fmt.Errorf("Error: failed to parse configs - %s ", err)
	}

	return nil
}

// Configコマンド - 現在の設定を表示する
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
