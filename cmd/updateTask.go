package cmd

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"github.com/eleliayub/cli/db"
	"github.com/spf13/cobra"
)

var updateTaskCmd = &cobra.Command{
	Use:   "update-task",
	Short: "Updating tasks",
	Long: "TODO()",
	Args:  cobra.ExactArgs(1),
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
		fmt.Println("Enter the number of the task you want to update:")
		var taskNumber int
		fmt.Scanf("%d", &taskNumber)
		reader := bufio.NewReader(os.Stdin) 
		switch args[0] {
			case "status":	
				fmt.Println("Enter the new status:")
				status, _ := reader.ReadString('\n')
				status = strings.TrimSuffix(status, "\n")
				db.UpdateTaskStatus(tasks[taskNumber-1].Title, status)
				fmt.Printf("Task status updated to %s\n", status)
			case "priority":
				fmt.Println("Enter the new priority:")
				priority, _ := reader.ReadString('\n')
				priority = strings.TrimSuffix(priority, "\n")
				db.UpdateTaskPriority(tasks[taskNumber-1].Title, priority)
				fmt.Printf("Task priority updated to %s\n", priority)
			case "title":
				fmt.Println("enter the new title")
				title, _ := reader.ReadString('\n')
				title = strings.TrimSuffix(title, "\n")
				db.UpdateTaskTitle(tasks[taskNumber-1].Title, title)
				fmt.Printf("Task title updated to %s\n", title)
			default:
				fmt.Println("Invalid argument. Use 'status', 'priority', or 'title'.")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateTaskCmd)
}
