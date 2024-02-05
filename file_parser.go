package main

import (
	"strings"
	"unicode"
)

func paseLineFromScanner(lineAsString string) string {
	words := strings.Fields(lineAsString) // uses spaces to separate meaning full units of text
	return parseWords(words)              // add a line brake after each line
}

func isIataCode(wordToEvaluate string) bool {

	if len(wordToEvaluate) < 3 {
		return false
	}
	if len(wordToEvaluate) > 3 && wordToEvaluate[0:1] == "#" && isCapitalLetter(wordToEvaluate[1:4]) {
		return true
	}
	return false
}

func isIcaoCode(wordToEvaluate string) bool {
	if len(wordToEvaluate) < 6 {
		return false
	}
	if len(wordToEvaluate) >= 6 && wordToEvaluate[0:2] == "##" && isCapitalLetter(wordToEvaluate[2:6]) {
		return true
	}
	return false
}

func setWord(word string) string {
	return word
}

func isIataCityCode(wordToEvaluate string) bool {

	if len(wordToEvaluate) < 3 {
		return false
	}
	if len(wordToEvaluate) > 3 && wordToEvaluate[0:1] == "*" && isCapitalLetter(wordToEvaluate[1:4]) {
		return true
	}
	return false
}

func isIcaoCityCode(wordToEvaluate string) bool {
	if len(wordToEvaluate) < 6 {
		return false
	}
	if len(wordToEvaluate) >= 6 && wordToEvaluate[0:2] == "*#" && isCapitalLetter(wordToEvaluate[2:6]) {
		return true
	}
	return false
}

func is24Time(wordToEvaluate string) bool {
	if len(wordToEvaluate) < 20 {
		return false
	}
	if wordToEvaluate[0:3] == "T24" && wordToEvaluate[(len(wordToEvaluate)-1):] == ")" {
		return true
	}
	return false
}

func is12Time(wordToEvaluate string) bool {
	if len(wordToEvaluate) < 20 {
		return false
	}
	if wordToEvaluate[0:3] == "T12" && wordToEvaluate[(len(wordToEvaluate)-1):] == ")" {
		return true
	}
	return false

}

func isDate(wordToEvaluate string) bool {
	if len(wordToEvaluate) < 20 {
		return false
	}
	if wordToEvaluate[0:2] == "D(" && wordToEvaluate[(len(wordToEvaluate)-1):] == ")" {
		return true
	}
	return false

}

func replaceWhiteSpaces(line string) (cleanLine string) {

	replaceWhiteSpaces := map[string]string{"\\v": "\n", "\\f": "\n", "\\r": "\n", "\\n": "\n"}
	cleanLine = line

	for oldVal, newVal := range replaceWhiteSpaces {
		cleanLine = strings.Replace(cleanLine, oldVal, newVal, -1)
	}
	return cleanLine
}

func isCapitalLetter(inputToAnalyze string) bool {
	for _, char := range inputToAnalyze {
		if !unicode.IsUpper(char) {
			return false
		}
	}
	return true
}

func evaluateWord(wordToEvaluate string) int {

	if isIataCode(wordToEvaluate) {
		return IS_IATA
	} else if isIcaoCode(wordToEvaluate) {
		return IS_ICAO
	} else if isIataCityCode(wordToEvaluate) {
		return IS_IATA_CITY
	} else if isIcaoCityCode(wordToEvaluate) {
		return IS_ICAO_CITY
	} else if is24Time(wordToEvaluate) {
		return IS_T24_TIME
	} else if is12Time(wordToEvaluate) {
		return IS_T12_TIME
	} else if isDate(wordToEvaluate) {
		return IS_DATE
	} else {
		return NO_PARSE
	}
}

// will return a parsed line.
func parseWords(wordsFromLine []string) string {

	parsedLine := ""

	for i, word := range wordsFromLine {

		switch evaluateWord(word) { // starts needed
		case IS_ICAO:
			parsedLine += findAirportNameByIcao(word)
		case IS_IATA:
			parsedLine += findAirportNameByIata(word)
		case IS_ICAO_CITY:
			parsedLine += findAirportMunicipalityByIcao(word)
		case IS_IATA_CITY:
			parsedLine += findAirportMunicipalityByIata(word)
		case IS_T12_TIME:
			parsedLine += parseTimeFormat12(word)
		case IS_T24_TIME:
			parsedLine += parseTimeFormat24(word)
		case IS_DATE:
			parsedLine += parseDate(word)
		case NO_PARSE:
			parsedLine += setWord(word)
		default:

		}

		if i < len(wordsFromLine)-1 {
			parsedLine += " "
		}
	}

	return parsedLine
}
