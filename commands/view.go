package commands

import (
	"Task-Tracker/model"
	"fmt"
	"log"
	"os"
)

// ViewTask filters and prints task from JSON file based on their task status.
//
// Possible arguments:
//
// "all" - Shows all the tasks
//
// "done" - Shows completed tasks
//
// "todo" - Shows pending tasks
//
// "in-progress" - Shows tasks in progress
//
// If no arguments are provided then defaults to "all".
// Panics on invalid arguments or file operation failures.
func ViewTask(args []string) {
	var status string
	if len(args) == 0 {
		status = "all"
	} else if len(args) == 1 {
		status = args[0]
	} else {
		log.Fatal("Invalid number of arguments.\n")
	}
	if status == "all" || status == "done" || status == "todo" || status == "in-progress" {
		jsonArray, jsonFile := readUnmarshallBytesFromFile(os.O_RDONLY)
		defer func() {
			err := jsonFile.Close()
			if err != nil {
				log.Fatal("Unable to close file.\n", err.Error())
			}
		}()
		if len(jsonArray) == 0 {
			print("No tasks found\n")
			return
		}
		iteratingAndPrintingTask(status, jsonArray)

	} else {
		log.Fatal("Invalid arguments.\n")
	}
}

// iteratingAndPrintingTask iterates over each task and prints its detail using printTask
// for formatted output.
func iteratingAndPrintingTask(taskStatus string, jsonArray []model.Task) {
	if taskStatus == "all" {
		for _, task := range jsonArray {
			printTask(task)
		}
	} else {
		for _, task := range jsonArray {
			if task.Status == taskStatus {
				printTask(task)
			}
		}
	}
}

// printTask prints the details of a task in formatted layout.
// It prints the task id, description, status, creation date and updated date
func printTask(task model.Task) {
	fmt.Printf("\n\n\n")
	fmt.Println("--------------------------------")
	fmt.Println("id:			\t", task.Id)
	fmt.Println("description:	\t", task.Description)
	fmt.Println("status:		\t", task.Status)
	fmt.Println("creation date:	\t", task.CreationTime)
	fmt.Println("updated date:	\t", task.UpdatingTime)
	fmt.Println("--------------------------------")
	fmt.Printf("\n\n\n")
}
