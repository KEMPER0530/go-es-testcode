package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// ②CLIを実際の処理 (Run()関数内) でフラグを参照
		// フラグが渡されていればそれを表示
		keyword, _ := cmd.Flags().GetString("keyword")
		if keyword != "" {
			fmt.Printf("hello world!  %s san!\n", keyword)
		} else {
			fmt.Printf("Not enough flags.\n")
		}
	},
}

func init() {
	var keyword string
	var size string
	var price string
	// ①CLIを初期化する処理(init()内)でフラグを定義
	// 第1引数: フラグ名、第2引数: 省略したフラグ名
	// 第3引数: デフォルト値、第4引数: フラグの説明
	searchCmd.Flags().StringVarP(&keyword, "keyword", "k", "", "検索ワード")
	searchCmd.Flags().StringVarP(&size, "size", "s", "", "取得するサイズ")
	searchCmd.Flags().StringVarP(&price, "price", "p", "", "価格")
	rootCmd.AddCommand(searchCmd)
}
