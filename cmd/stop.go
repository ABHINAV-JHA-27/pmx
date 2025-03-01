package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"kill"},
	Short:   "Stop a process",
	Long:    "Stop a process with the given ID",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stopping process with ID:", args[0])
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
