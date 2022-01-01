package cmd

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 設定格納用変数
var cfg = &Config{}

// 設定
type Config struct {
	ClientID        string `mapstructure:"client_id" json:"client_id,omitempty"`               // 本番環境向けのクライアントID (default: "")
	DevClientID     string `mapstructure:"dev_client_id" json:"dev_client_id,omitempty"`       // ローカル開発環境向けのクライアントID (default: "")
	RESTPort        int    `mapstructure:"rest_port" json:"rest_port,omitempty"`               // REST APIサーバーのポート番号 (default: 9000)
	GRPCPort        int    `mapstructure:"grpc_port" json:"grpc_port,omitempty"`               // gRPCサーバーのポート番号 (default: 9001)
	MariaDBHostname string `mapstructure:"mariadb_hostname" json:"mariadb_hostname,omitempty"` // DBのホスト (default: "mariadb")
	MariaDBPort     int    `mapstructure:"mariadb_port" json:"mariadb_port,omitempty"`         // DBのポート番号 (default: 3306)
	MariaDBUsername string `mapstructure:"mariadb_username" json:"mariadb_username,omitempty"` // DBのユーザー名 (default: "root")
	MariaDBPassword string `mapstructure:"mariadb_password" json:"mariadb_password,omitempty"` // DBのパスワード (default: "password")
	MariaDBDatabase string `mapstructure:"mariadb_database" json:"mariadb_database,omitempty"` // DBのDB名 (default: "Q-n-A")
}

// 設定を読み込む
func loadConfig(cfgFile string) error {
	// デフォルト値の設定
	viper.SetDefault("client_id", "")
	viper.SetDefault("dev_client_id", "")
	viper.SetDefault("rest_port", 9000)
	viper.SetDefault("grpc_port", 9001)
	viper.SetDefault("mariadb_hostname", "mariadb")
	viper.SetDefault("mariadb_port", 3306)
	viper.SetDefault("mariadb_username", "root")
	viper.SetDefault("mariadb_password", "password")
	viper.SetDefault("mariadb_database", "Q-n-A")

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
	err := viper.Unmarshal(cfg)
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
		fmt.Println("Current Configurations:")

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
