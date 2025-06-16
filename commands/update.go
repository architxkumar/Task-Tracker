package commands

import (
	"encoding/json"
	"log"
	"os"
)

// UpdateTaskDescription updates the description of the task with the supplied id
func UpdateTaskDescription(args []string) {
	TaskIdValidator(args)
	taskId := args[0]
	description := args[1]
	// Will be updated in the loop in case the task is present
	taskPresent := false
	// File is opened only O_RDWR mode, as updation should only be performed on already exists file
	jsonArray, jsonFile := ReadUnmarshallBytesFromFile(os.O_RDWR)
	for i, e := range jsonArray {
		if e.Id == taskId {
			jsonArray[i].Description = description
			taskPresent = true
		}
	}
	if !taskPresent {
		log.Fatal("Task not present in the list with the specific id")
	}
	output, err := json.Marshal(jsonArray)
	if err != nil {
		log.Fatal("Unable to marshall to json", err.Error())
	}
	TruncateAndWriteContent(jsonFile, output)
	err = jsonFile.Close()
	if err != nil {
		log.Fatal("Error closing File.", err.Error())
	}
	log.Print("Task Updated Successfully")
}
