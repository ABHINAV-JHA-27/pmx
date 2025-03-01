package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pmx",
	Short: "pmx is a process manager",
	Long:  "pmx is a CLI process manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to pmx!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
