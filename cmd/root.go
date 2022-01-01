package cmd

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// ルートコマンド - ダミーコマンド
var rootCmd = &cobra.Command{
	Use:  "Q-n-A",
	Long: "Q'n'A - traP Anonymous Question Box Service",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// バナーを表示
		printBanner()

		// 設定ファイルのパスを取得
		cfgFile, err := cmd.PersistentFlags().GetString("config")
		if err != nil {
			cfgFile = ""
		}

		// 設定を読み込む
		err = loadConfig(cfgFile)
		if err != nil {
			log.Panicf("failed to load config: %v", err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is config.json)")
}

// CLI実行
func Execute() error {
	return rootCmd.Execute()
}

// バナーを表示する
func printBanner() {
	// 文字部
	fmt.Println("Q'n'A - traP Anonymous Question Box Service")

	// デカ文字部
	fmt.Print(
		color.HiBlueString(`   ____`),
		color.WhiteString(`  _       _  `),
		color.HiRedString(`___`),
	)
	fmt.Println()
	fmt.Print(
		color.HiBlueString(`  / __ \`),
		color.WhiteString(`( )____ ( )`),
		color.HiRedString(`/   |`),
	)
	fmt.Println()
	fmt.Print(
		color.HiBlueString(` / / / /`),
		color.WhiteString(`|// __ \|/`),
		color.HiRedString(`/ /| |`),
	)
	fmt.Println()
	fmt.Print(
		color.HiBlueString(`/ /_/ /`),
		color.WhiteString(`  / / / / `),
		color.HiRedString(`/ ___ |`),
	)
	fmt.Println()
	fmt.Print(
		color.HiBlueString(`\___\_\`),
		color.WhiteString(` /_/ /_/ `),
		color.HiRedString(`/_/  |_|`),
	)
	fmt.Println()
	fmt.Println()
}
