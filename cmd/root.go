package cmd

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cfgFile string // 設定ファイルのパス

// ルートコマンド - ダミーコマンド
var rootCmd = &cobra.Command{
	Use:  "Q-n-A",
	Long: "Q'n'A - traP Anonymous Question Box Service",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// バナーを表示
		printBanner()

		// 設定を読み込む
		err := loadConfig(cfgFile)
		if err != nil {
			log.Panicf("failed to load config: %v", err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is config.json)")
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
		color.HiWhiteString(`  _       _  `),
		color.HiRedString(`___`),
	)
	fmt.Println()
	fmt.Print(
		color.HiBlueString(`  / __ \`),
		color.HiWhiteString(`( )____ ( )`),
		color.HiRedString(`/   |`),
	)
	fmt.Println()
	fmt.Print(
		color.HiBlueString(` / / / /`),
		color.HiWhiteString(`|// __ \|/`),
		color.HiRedString(`/ /| |`),
	)
	fmt.Println()
	fmt.Print(
		color.HiBlueString(`/ /_/ /`),
		color.HiWhiteString(`  / / / / `),
		color.HiRedString(`/ ___ |`),
	)
	fmt.Println()
	fmt.Print(
		color.HiBlueString(`\___\_\`),
		color.HiWhiteString(` /_/ /_/ `),
		color.HiRedString(`/_/  |_|`),
	)
	fmt.Println()
	fmt.Println()
}
