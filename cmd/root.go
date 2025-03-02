package cmd

import (
	"os"

	"github.com/ABHINAV-JHA-27/pmx/internal/logger"
	"github.com/ABHINAV-JHA-27/pmx/internal/process"
	"github.com/spf13/cobra"
)

var manager *process.Manager

func init() {
	manager = process.NewManager()
}

var rootCmd = &cobra.Command{
	Use:   "pmx",
	Short: "pmx is a process manager",
	Long:  "pmx is a CLI process manager",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Log.Println("Welcome to pmx")
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Log.Fatalf("Error executing root command: %v", err)
		os.Exit(1)
	}
}
