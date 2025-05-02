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
		tasks, err := db.GetTasks()
		if err != nil {
			panic(err)
		}
		fmt.Println("-----------------Tasks Lists----------------")
		fmt.Println("ID. Title")
		for i, task := range tasks {
			fmt.Printf("%d. %s (%s) -> %s\n", i + 1, task.Title, task.Priority, task.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(getTasksCmd)
}
