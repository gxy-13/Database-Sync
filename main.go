package main

import (
	"databasesync/cmd"
	"log"
)

func main() {
	cmd.Execute()
	log.Println("success")
}
