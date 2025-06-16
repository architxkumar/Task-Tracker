package commands

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
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
	characterRegex, err := regexp.MatchString(`^[a-zA-Z]+$`, args[0])
	if err != nil {
		log.Fatal("Unable to valid agrument input")
	}
	if characterRegex {
		log.Fatal("Invalid Task Id")
	}
}
