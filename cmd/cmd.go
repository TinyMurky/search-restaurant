package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	cmdsearchrestaurant "github.com/TinyMurky/search-restaurant/cmd/searchrestaurant"
	"github.com/TinyMurky/search-restaurant/internal/app/searchrestaurant/config"
)

// 這兩個變數會由旗標注入
var (
	cfgPath string
	envPath string
)

// Command wrao cobraCommand
type Command struct {
	cobraCommand *cobra.Command
}

// NewRootCmd create a Command pointer that has all command
func NewRootCmd() *Command {
	rootCmd := &cobra.Command{
		Use:   "sr",
		Short: "search restaurants near you",
		Long:  "A restaurant search tool that integrates AI and Google API.",
		// 在 PersistentPreRun 裡統一載入 env 與 config，
		// 子指令就不用重複做這些初始化工作
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// 1. 載入 .env（可選）

			if err := godotenv.Load(envPath); err != nil && !os.IsNotExist(err) {
				return fmt.Errorf("load env: %w", err)
			}

			// 2. 載入 YAML 設定（可選）
			if _, err := config.LoadConfig(cfgPath); err != nil {
				return fmt.Errorf("load config: %w", err)
			}

			return nil
		},
	}

	/* ---------------- Flag 區 ---------------- */
	rootCmd.PersistentFlags().StringVarP(
		&cfgPath,
		"config",
		"c",
		"configs/config.yaml",
		"Path to config.yaml",
	)

	rootCmd.PersistentFlags().StringVarP(
		&envPath,
		"env",
		"e",
		".env",
		"Path to .env file",
	)
	/* ---------------------------------------- */

	// 掛上子指令
	rootCmd.AddCommand(cmdsearchrestaurant.SearchCmd)

	return &Command{cobraCommand: rootCmd}
}

// Execute is the entry that main() will call
func (c *Command) Execute() {
	if err := c.cobraCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
