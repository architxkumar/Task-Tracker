package commands

import (
	"log"
	"os"
	"regexp"
)

// TruncateAndWriteContent overwrites the content of the json file with
// the supplied JsonMarshall bytes and takes in file pointer
func TruncateAndWriteContent(jsonFile *os.File, output []byte) {
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

// TaskIdValidator checks the supplied Task id for updation, deletion of task
// by comparing the length of supplied argument and non-alphabetical values
func TaskIdValidator(args []string) {
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
