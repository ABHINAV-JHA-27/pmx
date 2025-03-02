package cmd

import (
	"github.com/ABHINAV-JHA-27/pmx/internal/logger"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"kill"},
	Short:   "Stop a process",
	Long:    "Stop a process with the given ID",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_name := cmd.Flag("name").Value.String()
		if len(args) == 0 && _name == "" {
			logger.Log.Error("Please provide a process ID or name")
			cmd.Help()
			return
		}
		_ID := args[0]
		if _ID != "" {
			manager.StopProcess(_ID)
			logger.Log.Println("Process stopped with ID: ", _ID)
			return
		}
		if _name != "" {
			for _, p := range manager.GetProcesses() {
				if p.Name == _name {
					manager.StopProcess(p.ID)
					logger.Log.Println("Process stopped with name: ", _name)
					return
				}
			}
		}
	},
}

func init() {
	stopCmd.Flags().StringP("name", "n", "", "Name of the process")
	rootCmd.AddCommand(stopCmd)
}
