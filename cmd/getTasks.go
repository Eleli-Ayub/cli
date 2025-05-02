package cmd

import (
	"fmt"
	"github.com/eleliayub/cli/db"
	"github.com/spf13/cobra"
)

var getTasksCmd = &cobra.Command{
	Use:   "get-tasks",
	Short: "get all tasks",
	Long: "TODO()",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("-----------------Getting tasks----------------")
		// Call the function to get tasks from the database
		tasks, err := db.GetTasks()
		if err != nil {
			panic(err)
		}
		for i, task := range tasks {
			fmt.Printf("%d. Title: %s, Priority: %s, Status: %s\n", i + 1, task.Title, task.Priority, task.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(getTasksCmd)
}
