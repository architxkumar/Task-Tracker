package main

import (
	"Task-Tracker/commands"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Arguments not provided")
	}
	command := strings.ToLower(os.Args[1])
	argsArray := os.Args[2:]

	switch command {
	case "add":
		commands.AddTask(argsArray)
	case "update":
		commands.UpdateTaskDescription(argsArray)
	case "delete":
		commands.DeleteTask(argsArray)
	case "mark-done":
		commands.UpdateTaskProgress(argsArray, "done")
	case "mark-in-progress":
		commands.UpdateTaskProgress(argsArray, "in-progress")
	case "list":
		commands.ViewTask(argsArray)
	default:
		log.Fatal("Invalid command")
	}
}
