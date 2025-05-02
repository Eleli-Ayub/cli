package cmd

import (
	"bufio"
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"os"
	"github.com/eleliayub/cli/db"
)


var addTasksCmd = &cobra.Command{
	Use:   "add-task",
	Short: "flag to add tasks",
	Long:  "TODO()",
	Run: func(cmd *cobra.Command, args []string) {
		newTask := db.Task{}
		fmt.Println("-----------------Adding task----------------")
		fmt.Println("*Task Title")

		reader := bufio.NewReader(os.Stdin)
		newTask.Title, _ = reader.ReadString('\n')
		newTask.Title = strings.TrimSuffix(newTask.Title, "\n")

		fmt.Println("*Priority")
		newTask.Priority, _ = reader.ReadString('\n')
		newTask.Priority = strings.TrimSuffix(newTask.Priority, "\n")

		if err := db.InsertTask(newTask); err != nil {
			panic(err)
		}
		fmt.Println("Task added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addTasksCmd)
}
