package commands

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// UpdateTaskDescription updates the description of the task with the supplied id
func UpdateTaskDescription(args []string) {
	inputArgumentValidator(args)
	taskId := args[0]
	updatedDescription := args[1]
	// File is opened only O_RDWR mode, as updation should only be performed on already exists file
	jsonArray, jsonFile := readUnmarshallBytesFromFile(os.O_RDWR)
	index := getTaskIndex(jsonArray, taskId)
	jsonArray[index].Description = updatedDescription
	jsonArray[index].UpdationTime = time.Now()
	output, err := json.Marshal(jsonArray)
	if err != nil {
		log.Fatal("Unable to marshall to json", err.Error())
	}
	truncateAndWriteContent(jsonFile, output)
	err = jsonFile.Close()
	if err != nil {
		log.Fatal("Error closing File.", err.Error())
	}
	log.Print("Task Updated Successfully")
}

// inputArgumentValidator checks for the validity of the
// program arguments during [UpdateTaskDescription] function usage
func inputArgumentValidator(args []string) {
	if len(args) != 2 {
		log.Fatal("Invalid command Usage")
	}
	// Input Validation: The id shouldn't contain alphabets
	TaskIdValidator(args[1])
}

func UpdateTaskProgress(args []string, status string) {
	if len(args) != 1 {
		log.Fatal("Invalid command Usage")
	}
	TaskIdValidator(args[0])
	taskId := args[0]
	jsonArray, jsonFile := readUnmarshallBytesFromFile(os.O_RDWR)
	index := getTaskIndex(jsonArray, taskId)
	jsonArray[index].Status = status
	jsonArray[index].UpdationTime = time.Now()
	output, err := json.Marshal(jsonArray)
	if err != nil {
		log.Fatal("Unable to marshall to json", err.Error())
	}
	truncateAndWriteContent(jsonFile, output)
	log.Print("Task Status updated successfully")
}
