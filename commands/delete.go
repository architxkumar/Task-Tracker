package commands

import (
	"encoding/json"
	"log"
	"os"
)

// DeleteTask removes the task from "tasks.json"
// by supplying the associate task id if present, else
// it exists with an error. It does by reading the contents
// into array of tasks model, removing the task from the array
// and then overwriting the array contents into the file.
func DeleteTask(args []string) {
	if len(args) != 1 {
		log.Fatal("Invalid command usage: Invalid number of arguments.\n")
	}
	TaskIdValidator(args[0])
	taskId := args[0]
	jsonArray, jsonFile := readUnmarshallBytesFromFile(os.O_RDWR)
	taskIndex := getTaskIndex(jsonArray, taskId)
	jsonArray = append(jsonArray[:taskIndex], jsonArray[taskIndex+1:]...)
	output, err := json.Marshal(jsonArray)
	if err != nil {
		log.Fatal("Unable to marshall to json.\n", err.Error())
	}
	truncateAndWriteContent(jsonFile, output)
	log.Print("Task Deleted Successfully.\n")
}
