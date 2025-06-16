package commands

import (
	"encoding/json"
	"log"
	"os"
)

func DeleteTask(args []string) {
	if len(args) != 1 {
		log.Fatal("Invalid Arguments")
	}
	TaskIdValidator(args[0])
	taskId := args[0]
	jsonArray, jsonFile := readUnmarshallBytesFromFile(os.O_RDWR)
	taskIndex := getTaskIndex(jsonArray, taskId)
	jsonArray = append(jsonArray[:taskIndex], jsonArray[taskIndex+1:]...)
	output, err := json.Marshal(jsonArray)
	if err != nil {
		log.Fatal("Unable to marshall to json", err.Error())
	}
	truncateAndWriteContent(jsonFile, output)
	log.Print("Task Deleted Successfully")
}
