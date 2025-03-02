package cmd

import (
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"kill"},
	Short:   "Stop a process",
	Long:    "Stop a process with the given ID",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_name := cmd.Flag("name").Value.String()
		_ID := args[0]
		if _ID == "" {
			if _name != "" {
			} else {
				cmd.Help()
			}
		} else {
			manager.StopProcess(_ID)
		}
	},
}

func init() {
	stopCmd.Flags().StringP("name", "n", "", "Name of the process")
	rootCmd.AddCommand(stopCmd)
}
