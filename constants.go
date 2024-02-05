package main

const (
	HELP      = "-h"
	HELP_TEXT = "itinerary usage:\n go run . ./input.txt ./output.txt ./airport-lookup.csv"

	IS_ICAO      = 1
	IS_IATA      = 2
	IS_ICAO_CITY = 3
	IS_IATA_CITY = 4
	IS_T12_TIME  = 5
	IS_T24_TIME  = 6
	IS_DATE      = 7
	NO_PARSE     = 8

	ERR_INPUT            = "Input not found"
	ERR_INVALID_USE      = "Invalid use of program run go run . -h for help!"
	ERR_AIRPORT_DB_FILE  = "Airport lookup not found"
	ERR_DATA_MALFORMED   = "Airport lookup malformed"
	ERR_FILE_NOT_CREATED = "File not created!"
	ERR_WRITING_TO_FILE  = "Writing to file failed"
)
