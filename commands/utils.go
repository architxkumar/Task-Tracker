package commands

import (
	"Task-Tracker/model"
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
)

// truncateAndWriteContent overwrites the content of the json file with
// the supplied JsonMarshall bytes and takes in file pointer
func truncateAndWriteContent(jsonFile *os.File, output []byte) {
	err := jsonFile.Truncate(0)
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
}

// readUnmarshallBytesFromFile opens the tasks.json in the specific mode
// supplied with the flag argument, reads the contents, unmarshal them into task model
// and then returns it along with file pointer.
func readUnmarshallBytesFromFile(flag int) ([]model.Task, *os.File) {
	jsonFile, err := os.OpenFile("tasks.json", flag, 0644)
	if err != nil {
		log.Fatal("Error opening or creating tasks.json.", err.Error())
	}
	content, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error reading contents from tasks.json.", err.Error())
	}
	var jsonObjects []model.Task
	if len(content) != 0 {
		err = json.Unmarshal(content, &jsonObjects)
		if err != nil {
			log.Fatal("Error parsing contents from tasks.json.", err.Error())
		}
	}
	return jsonObjects, jsonFile
}

// getTaskIndex returns the index of task from the array,
// if present else will exit the code with "Task not present" message.
func getTaskIndex(jsonArray []model.Task, taskId string) int {
	index := -1
	for i, e := range jsonArray {
		if e.Id == taskId {
			index = i
		}
	}
	if index < 0 {
		log.Fatal("Task not present in the list with the specific id")
	}
	return index
}

// TaskIdValidator checks the supplied id against regex pattern for alphabets.
// It ensures task id doesn't contain characters.
func TaskIdValidator(taskId string) {
	characterRegex, err := regexp.MatchString(`^[a-zA-Z]+$`, taskId)
	if err != nil {
		log.Fatal("Unable to validate task id")
	}
	if characterRegex {
		log.Fatal("Invalid command Usage: Invalid Task Id")
	}
}
