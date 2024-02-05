package main

import (
	"os"
)

type Airports struct {
	name         string `csv:"name"`
	iso_country  string `csv:"iso_country"`
	municipality string `csv:"municipality"`
	icao_code    string `csv:"icao_code"`
	iata_code    string `csv:"iata_code"`
	coordinates  string `csv:"coordinates"`
}

func findAirportNameByIata(iata_code string) (airportName string) {

	airportData, _ := readCSVFile(os.Args[3])
	iataIndex := -1

	for i, code := range airportData["iata_code"] {
		if code == iata_code[1:4] {
			iataIndex = i
		}
	}
	if iataIndex == -1 {
		airportLookupFailed()
		return iata_code // if not code is found returns the same code  does not crash
	}

	airportNames, exists := airportData["name"]

	if exists {

		return airportNames[iataIndex]
	}
	return iata_code
}

func findAirportNameByIcao(icao_code string) (airportName string) {
	airportData, _ := readCSVFile(os.Args[3])
	icaoIndex := -1
	// removed the constraints
	for i, code := range airportData["icao_code"] {
		if code == icao_code[2:6] {
			icaoIndex = i
		}
	}

	if icaoIndex == -1 {
		airportLookupFailed()
		return icao_code
	}

	airportNames, exists := airportData["name"]
	// happy path
	if exists {
		return airportNames[icaoIndex]
	}
	// sad path
	return icao_code
}

func findAirportMunicipalityByIcao(icao_code string) (cityName string) {

	icaoIndex := -1
	
	airportData, _ := readCSVFile(os.Args[3])
	for i, code := range airportData["icao_code"] {
		if code == icao_code[2:6] {
			icaoIndex = i
		}
	}

	if icaoIndex == -1 {
		airportLookupFailed()
		return icao_code
	}

	airportMunicipalities, exists := airportData["municipality"]

	if exists {
		return airportMunicipalities[icaoIndex]
	}

	return icao_code
}

func findAirportMunicipalityByIata(iata_code string) (cityName string) {
	iataIndex := -1
	airportData, _ := readCSVFile(os.Args[3])

	for i, code := range airportData["iata_code"] {
		if code == iata_code[1:4] {
			iataIndex = i
		}
	}

	if iataIndex == -1 {
		airportLookupFailed()
		return iata_code 
	}

	airportMunicipalities, exists := airportData["municipality"]

	if exists {
		return airportMunicipalities[iataIndex]
	}
	return iata_code 
}
