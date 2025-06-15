package commands

import (
	"Task-Tracker/model"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// AddTask is used to add tasks in .json file.
// If the file exists, it reads all the contents from it,
// then overwrites the file by truncating and then adding
// the new updated json array to the .json file
func AddTask(args []string) {
	if len(args) != 1 {
		log.Fatal("Invalid command usage\n")
	}
	description := args[0]
	// HACK: using hardcoded string instead of typed enum
	// TODO: create status enum with three possible values: Todo, in-progress, done
	status := "todo"
	task := model.Task{Id: strconv.FormatInt(time.Now().UnixNano(), 10), Status: status, Description: description, CreationTime: time.Now(), UpdationTime: time.Now()}
	jsonFile, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
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
	jsonObjects = append(jsonObjects, task)
	output, err := json.Marshal(jsonObjects)
	if err != nil {
		log.Fatal("Unable to marshall to json", err.Error())
	}
	TruncateAndWriteContent(jsonFile, output)
	err = jsonFile.Close()
	if err != nil {
		log.Fatal("Error closing File.", err.Error())
	}
	log.Print("Task successfully created")
}
