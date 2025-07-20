package cmdsearchrestaurant

import (
	"fmt"

	"github.com/spf13/cobra"
)

// SearchCmd ­= sr search ...
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search restaurants and export xlsx",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🎉 這裡已經可以使用 config 與 env 內的變數了")
		// 你的搜尋邏輯...
		return nil
	},
}
