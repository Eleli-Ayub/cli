package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Manage tasks on the terminal",
	Long:  `TODO(())`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("running root")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
