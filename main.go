package main

import (
	"fmt"
	"log"
	"context"
	"github.com/eleliayub/cli/db"
	"github.com/joho/godotenv"
)

import "github.com/eleliayub/cli/cmd"

func main() {
	fmt.Println("CLI TASK MANAGER")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if _, err =	db.ConnectToDb();err != nil{
		panic(err)
	}
	defer func() {
		if err = db.DBClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	cmd.Execute()
}
