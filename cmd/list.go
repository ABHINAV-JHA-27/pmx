package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all processes",
	Long:  "List all processes currently running",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing all processes")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
