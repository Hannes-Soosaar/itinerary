package main

import (
	"fmt"
	"os"
)

// global variable to track errors
var noErrors bool 

// global variable to hold the airport data
var airportData = make(map[string][]string)

func validateStartup() {
	if len(os.Args) == 2 && (HELP == os.Args[1]) {
		// fmt.Println(os.Args[1])
		fmt.Println(HELP_TEXT)
		return
	} else if len(os.Args) == 4 {
		validateInput()
		return
	} else {
		argumentNumberError()
		return
	}
}

func validateInput() {
	noErrors=true

	/*! the order of this function is important the airportData must be created before
	you can call readFile*/

	airportData, err := readCSVFile(os.Args[3])
	if err != nil {
		fmt.Println(err)
		noErrors = false
		return
	}
	
	if airportData == nil { // if data is malformed the CSV file reader will return nil
		return
	}
	parsedInput := getParseFileContent(os.Args[1])

	if noErrors {
	createClientOutput(parsedInput)
	}

}
func main() {
	validateStartup()
}
