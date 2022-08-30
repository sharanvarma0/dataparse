package json

import (
	"log"
	"io/ioutil"
	"os"
	"encoding/json"
)

func readJsonFromFile(filename string) (interface{}, error) {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error occurred while reading file: " + err.Error())
		return nil, err
	}

	var parsedData interface{}

	err = json.Unmarshal(filedata, &parsedData)
	if err != nil {
		log.Fatal("Error occurred while unmarshaling: " + err.Error())
		return nil, err
	}

	return parsedData, nil
}
