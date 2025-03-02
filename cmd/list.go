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
		_name := cmd.Flag("name").Value.String()
		_id := cmd.Flag("id").Value.String()
		_status := cmd.Flag("status").Value.String()
		_pid := cmd.Flag("pid").Value.String()

		if _name != "" {
			fmt.Printf("Name: %s\n", _name)
		}
		if _id != "" {
			fmt.Printf("ID: %s\n", _id)
		}
		if _status != "" {
			fmt.Printf("Status: %s\n", _status)
		}
		if _pid != "" {
			fmt.Printf("PID: %s\n", _pid)
		}

		if _name == "" && _id == "" && _status == "" && _pid == "" {
			fmt.Println("All processes")
		}

	},
}

func init() {
	listCmd.Flags().StringP("name", "n", "", "Name of the process")
	listCmd.Flags().StringP("id", "i", "", "ID of the process")
	listCmd.Flags().StringP("status", "s", "", "Status of the process")
	listCmd.Flags().StringP("pid", "p", "", "PID of the process")
	rootCmd.AddCommand(listCmd)
}
