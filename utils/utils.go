package utils

import (
	"os"
)

func ReadFromFile(fileName string) []byte {
	dat, err := os.ReadFile(fileName)
    if err != nil {
		Log("Error while reading file data: ",err.Error())
		return []byte{}
	}

	return dat
}

func WriteToFile(fileName string, data []byte) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		Log("Error while opening file for writing: ",err.Error())
	}

	if _, err = f.WriteString(string(data)); err != nil {
		Log("Error while writing data to file: ",err.Error())
	}
}