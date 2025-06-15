package commands

import (
	"log"
	"os"
)

func TruncatingAndWriteContents(jsonFile *os.File, output []byte) {
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
