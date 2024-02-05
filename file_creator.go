package main

import (
	"fmt"
	"os"
)

func createClientOutput(dataToOutput string) {

	file, err := os.Create("./test/output.txt")
	if err != nil {
		fileNotCreated()
	}

	defer file.Close()

	_, err = fmt.Fprint(file, dataToOutput)

	if err != nil {
		writingToFileFailed()
	}
}
