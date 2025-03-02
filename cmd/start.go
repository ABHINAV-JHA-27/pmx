package cmd

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"run"},
	Short:   "Start a process",
	Long:    "Start a process with the given command and arguments",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_name := cmd.Flag("name").Value.String()
		_restart, _ := cmd.Flags().GetBool("restart")
		_cmd := args[0]
		manager.StartProcess(_name, _cmd, _restart)
	},
}

func init() {
	startCmd.Flags().StringP("name", "n", "", "Name of the process")
	startCmd.Flags().BoolP("restart", "r", false, "Restart process on failure")
	rootCmd.AddCommand(startCmd)
}
