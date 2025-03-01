package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"run"},
	Short:   "Start a process",
	Long:    "Start a process with the given command and arguments",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting process with command:", args[0])
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
