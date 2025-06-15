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
		fmt.Printf("%s\n", command)
	case "delete":
		fmt.Printf("%s\n", command)
	case "mark-done":
		fmt.Printf("%s\n", command)
	case "mark-in-progress":
		fmt.Printf("%s\n", command)
	case "list":
		fmt.Printf("%s\n", command)
	default:
		fmt.Printf("Invalid command\n")
		os.Exit(1)
	}
	os.Exit(0)
}
