package main

import "fmt"

/*
custom error handling
when any of the below errors are called, the output file will not be created
*/

func argumentNumberError() {
	fmt.Println(ERR_INVALID_USE)
	noErrors = false
}

func doesNotExist() {
	fmt.Println(ERR_INPUT)
	noErrors = false
}

func airportLookupFailed() {
	fmt.Println(ERR_AIRPORT_DB_FILE)
	noErrors = false
}

func airportLookupMalformed() {
	fmt.Println(ERR_DATA_MALFORMED)
	noErrors = false
}

func fileNotCreated() {
	fmt.Println(ERR_FILE_NOT_CREATED)
	noErrors = false
}

func writingToFileFailed() {
	fmt.Print(ERR_WRITING_TO_FILE)
	noErrors = false
}

func errorReadingFile(err error) {
	fmt.Println("Error reading file:", err)
	noErrors = false
}

func setError() {
	noErrors = false
}
