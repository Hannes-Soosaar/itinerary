package main

import (
	"bufio"
	"encoding/csv"
	"os"
)

type FileReader struct {
}

func getParseFileContent(filePath string) string {

	inputData := ""

	file, err := os.Open(filePath)
	if err != nil {
		doesNotExist()
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// var emptyLine int
	var previousLine string
	var addToInput bool
	for scanner.Scan() {

		line := scanner.Text()
		cleanedLine := replaceWhiteSpaces(line)

		if previousLine == cleanedLine && cleanedLine == "" || previousLine == cleanedLine && cleanedLine == "\n" {
			addToInput = false
		}

		if addToInput {
			if inputData != "" {
				inputData += "\n"
			}
			inputData += paseLineFromScanner(cleanedLine) // takes in a scanner line and parses it
		}

		previousLine = cleanedLine
		addToInput = true
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return inputData // returns one single string as the the output to be Written to the file
}

func readCSVFile(filePath string) (map[string][]string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		airportLookupFailed()
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		airportLookupMalformed()
		return nil, err
	}

	columnNames := records[0]

	dataMap := make(map[string][]string)

	for _, columnName := range columnNames {
		dataMap[columnName] = nil
	}

	for _, record := range records[1:] {
		for i, value := range record {
			columnName := columnNames[i]
			dataMap[columnName] = append(dataMap[columnName], value)
		}
	}

	for _, columnName := range dataMap {
		for _, value := range columnName {
			if value == "" {
				airportLookupMalformed()
				return nil, nil
			}
		}
	}
	return dataMap, nil
}
