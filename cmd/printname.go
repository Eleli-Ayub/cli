/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// printnameCmd represents the printname command
var printnameCmd = &cobra.Command{
	Use:   "printname",
	Short: "very brienf intro",
	Long:  `Just another explanation of this code`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		authFlag, _ := cmd.Flags().GetString("auth")
		cmdFlag, _ := cmd.Flags().GetString("open")
		if authFlag == "login" {
			fmt.Println(args[0] + " is login in")
		} else {
			fmt.Println("please login to use")
		}
		if cmdFlag == "docs" {
			path, err := filepath.Abs("./index.html")
			if err != nil {
				log.Printf("error finding absolute path: %v\n", err)
				return
			}
			comm := exec.Command("xdg-open", path)

			if err := comm.Start(); err != nil {
				log.Printf("error: %v\n", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(printnameCmd)
	printnameCmd.Flags().String("auth", "", "this is a flag for authentication")
	printnameCmd.Flags().String("open", "", "this is a flag for list c")
}
