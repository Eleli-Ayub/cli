/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// printnameCmd represents the printname command
var printnameCmd = &cobra.Command{
	Use:   "printname",
	Short: "very brienf intro",
	Long: `Just another explanation of this code`,
  Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
    authFlag, _ := cmd.Flags().GetString("auth")
    if authFlag == "login"{
      fmt.Println(args[0] + " is login in")
    }else{
      fmt.Println("please login to use")
    }
	},
}

func init() {
	rootCmd.AddCommand(printnameCmd)
  printnameCmd.Flags().String("auth","","this is a flag for authentication" )
}
