package commands

import (
	"Task-Tracker/model"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"
)

// AddTask is used to add tasks in .json file.
// If the file exists, it reads the contents from the file,
// creates new json array with latest task and then overwrites
// the file content else, new file is created and contents are directly written.
func AddTask(args []string) {
	if len(args) != 1 {
		log.Fatal("Invalid command usage.\n")
	}
	description := args[0]
	// HACK: using hardcoded string instead of typed enum
	// TODO: create status enum with three possible values: Todo, in-progress, done
	status := "todo"
	// HACK: task id uses UNIX timestamp in nanoseconds in order to ensure unique id generation
	task := model.Task{Id: strconv.FormatInt(time.Now().UnixNano(), 10), Status: status,
		Description: description, CreationTime: time.Now(), UpdatingTime: time.Now()}
	jsonArray, jsonFile := readUnmarshallBytesFromFile(os.O_RDWR | os.O_CREATE)
	jsonArray = append(jsonArray, task)
	output, err := json.Marshal(jsonArray)
	if err != nil {
		log.Fatal("Unable to marshall to json.\n", err.Error())
	}
	truncateAndWriteContent(jsonFile, output)
	err = jsonFile.Close()
	if err != nil {
		log.Fatal("Error closing File.\n", err.Error())
	}
	log.Print("Task successfully created")
}
