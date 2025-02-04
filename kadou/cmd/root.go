package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kadou",
	Short: "稼働を管理するCLIアプリケーション",
	Long: `kadouは稼働時間を管理するためのCLIアプリケーションです。
開始・終了時刻、休憩時刻などの稼働に関する情報を管理します。`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.HelpFunc()(cmd, args)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
