package main

import (
	"Task-Tracker/commands"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Arguments not supplied\n")
		os.Exit(1)
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
		fmt.Printf("Invalid command\n")
		os.Exit(1)
	}
	os.Exit(0)
}
