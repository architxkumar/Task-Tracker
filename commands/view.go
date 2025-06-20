package commands

import (
	"Task-Tracker/model"
	"fmt"
	"log"
	"os"
)

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
		if len(jsonArray) == 0 {
			print("No tasks found\n")
			os.Exit(0)
		}
		iteratingAndPrintingTask(status, jsonArray)
		err := jsonFile.Close()
		if err != nil {
			log.Fatal("Unable to close file.\n", err.Error())
		}
	} else {
		log.Fatal("Invalid arguments.\n")
	}
}

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
