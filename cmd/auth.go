package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/eleliayub/cli/utils"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "login or signup",
	Long: "TODO()",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		authType := 	args[0]

		switch authType {
			case "login":
				username, _ := cmd.Flags().GetString("u")
				password, _ := cmd.Flags().GetString("p")
				fmt.Println(username,password)
				break;
			case "register":
				utils.RegisterUser()
				break;
			default: 
				fmt.Println("no value chosen")
				break;
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().String("u", "", "username for user")
	authCmd.Flags().String("p", "", "password for user")
}
