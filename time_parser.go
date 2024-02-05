package main

import (
	"time"
)

func parseTimeFormat12(timeToParse string) string {
	trimmedTimeToParse := timeToParse[4:(len(timeToParse) - 1)]

	var parsedTime time.Time
	var formattedTime string

	if trimmedTimeToParse[(len(trimmedTimeToParse)-6)] == '-' || trimmedTimeToParse[len(trimmedTimeToParse)-6] == '+' {
		parsedTime, _ = time.Parse("2006-01-02T15:04-07:00", trimmedTimeToParse)
	} else {
		parsedTime, _ = time.Parse("2006-01-02T15:04Z", trimmedTimeToParse)
	}

	_, offset := parsedTime.Zone()

	if offset != 0 {
		formattedTime = parsedTime.Format("03:04PM (-07:00)")
	} else {
		formattedTime = parsedTime.Format("03:04AM (+00:00)")
	}
	return formattedTime
}

func parseTimeFormat24(timeToParse string) string {
	trimmedTimeToParse := timeToParse[4:(len(timeToParse) - 1)]

	var parsedTime time.Time
	var formattedTime string

	if trimmedTimeToParse[(len(trimmedTimeToParse)-6)] == '-' || trimmedTimeToParse[len(trimmedTimeToParse)-6] == '+' {
		parsedTime, _ = time.Parse("2006-01-02T15:04-07:00", trimmedTimeToParse)
	} else {
		parsedTime, _ = time.Parse("2006-01-02T15:04Z", trimmedTimeToParse)
	}

	_, offset := parsedTime.Zone()

	if offset != 0 {
		formattedTime = parsedTime.Format("15:04 (-07:00)")
	} else {
		formattedTime = parsedTime.Format("15:04 (+00:00)")
	}
	return formattedTime
}

func parseDate(dateToParse string) string {
	trimmedTimeToParse := dateToParse[2:12]
	parsedDate, err := time.Parse("2006-01-02", trimmedTimeToParse)
	if err != nil {
		panic(err)
	}
	formattedTime := parsedDate.Format("02 Jan 2006")
	return formattedTime
}
