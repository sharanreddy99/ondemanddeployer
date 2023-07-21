package utils

import (
	"fmt"
	"os"
)

func Log(params ...interface{}) {
	// Write to stdout
	// log.Println(params...)

	//Write to file
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%v", params)); err != nil {
		panic(err)
	}
}
