package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// GetJSON ...
func GetJSON(jsonFileName string) []byte {
	jsonPath := fmt.Sprintf("./utils/fixtures/%s.json", jsonFileName)
	file, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("File error: %v\n", err)
		os.Exit(1)
	}
	return file
}
