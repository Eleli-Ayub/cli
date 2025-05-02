package cmd

import (
	"fmt"
	"github.com/eleliayub/cli/db"
	"github.com/spf13/cobra"
)

var deleteTaskCmd = &cobra.Command{
	Use:   "delete-task",
	Short: "Deleting tasks",
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

		fmt.Println("Enter the number of the task you want to delete:")
		var taskNumber int
		fmt.Scanf("%d", &taskNumber)
	
		fmt.Printf("Confirm deleting task %s \n(y/n)", tasks[taskNumber-1].Title)
		smallInput := ""
		fmt.Scanf("%s", &smallInput)
		if smallInput == "y" {
			db.DeleteTask(tasks[taskNumber-1].Title)
			fmt.Println("Task deleted successfully.")
		}else if smallInput == "n" {
			fmt.Println("Task deletion cancelled.")
		}else{
			fmt.Println("Invalid input. Task deletion cancelled.")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)
}
