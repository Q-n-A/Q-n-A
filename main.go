/*
Q'n'A - traP Anonymous Question Box Service
traP内部向けの質問箱サービスです。
*/
package main

import (
	"log"
	"time"

	"github.com/Q-n-A/Q-n-A/cmd"
)

// タイムゾーンの設定
func init() {
	const location = "Asia/Tokyo"

	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}

	time.Local = loc
}

func main() {
	// CLI実行
	if err := cmd.Execute(); err != nil {
		log.Panicf("failed to start Q'n'A application: %v", err)
	}
}
