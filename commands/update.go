package commands

import (
	"Task-Tracker/model"
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
)

func UpdateTaskDescription(args []string) {
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
	taskId := args[0]
	description := args[1]
	taskPresent := false
	jsonFile, err := os.OpenFile("tasks.json", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("Error opening tasks.json.", err.Error())
	}
	content, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error reading contents from tasks.json.", err.Error())
	}
	var jsonArray []model.Task
	if err = json.Unmarshal(content, &jsonArray); err != nil {
		log.Fatal("Error decoding contents from tasks.json", err.Error())
	}
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
	err = jsonFile.Truncate(0)
	if err != nil {
		log.Fatal("Unable to truncate json file.", err.Error())
	}
	_, err = jsonFile.Seek(0, 0)
	if err != nil {
		log.Fatal("Unable to reset pointer location.", err.Error())
	}
	_, err = jsonFile.Write(output)
	if err != nil {
		log.Fatal("Unable to write to json.", err.Error())
	}
	err = jsonFile.Close()
	if err != nil {
		log.Fatal("Error closing File.", err.Error())
	}
	log.Print("Task Updated Successfully")
}
