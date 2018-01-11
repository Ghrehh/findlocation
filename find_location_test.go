package findlocation

import (
	"testing"
)

var countriesAndCodes = map[string]string{
	"countriesAndCodesKey": "countriesAndCodesValue",
}

var citiesWithPopulationsOver15000 = map[string]string{
	"citiesWithPopulationsOver15000Key": "citiesWithPopulationsOver15000Value",
}

var alternativeNamesForCitiesAndCountries = map[string]string{
	"alternativeNamesForCitiesAndCountriesKey": "alternativeNamesForCitiesAndCountriesValue",
}

var usStates = []string{"usStates"}

var ukCounties = []string{"ukCounties"}

var canadianProvincesAndCodes = map[string]string{
	"canadianProvincesAndCodesKey": "canadianProvincesAndCodesValue",
}

var australianStatesAndCodes = map[string]string{
	"australianStatesAndCodesKey": "australianStatesAndCodesValue",
}

var usStateCodes = []string{"usStateCodes"}

var locationFinder = &LocationFinder {
	AlternativeNamesForCitiesAndCountries: alternativeNamesForCitiesAndCountries,
	AustralianStatesAndCodes: australianStatesAndCodes,
	CanadianProvincesAndCodes: canadianProvincesAndCodes,
	CitiesWithPopulationsOver15000: citiesWithPopulationsOver15000,
	CountriesAndCodes: countriesAndCodes,
	UkCounties: ukCounties,
	UsStates: usStates,
	UsStateCodes: usStateCodes,
}

func errorMessage(result string, expectedResult string) string {
	return "expected '" + result + "' to equal '" + expectedResult + "'"
}

func TestFindLocationCountriesAndCodesByValue(t *testing.T) {
	tests := map[string]string{
		"countriesAndCodesKey": "countriesAndCodesValue",
		"citiesWithPopulationsOver15000Key": "citiesWithPopulationsOver15000Value",
		"alternativeNamesForCitiesAndCountriesKey" : "alternativeNamesForCitiesAndCountriesValue",
		"usStates": "US",
		"canadianProvincesAndCodesKey": "CA",
		"australianStatesAndCodesKey": "AU",
		"countriesAndCodesValue": "countriesAndCodesValue",
		"usStateCodes": "US",
		"canadianProvincesAndCodesValue": "CA",
		"australianStatesAndCodesValue": "AU",
		"random gibberish": "",
	}

	for searchString, expectedResult := range tests {
		result := locationFinder.FindLocation(searchString)

		if result != expectedResult {
			t.Error(errorMessage(result, expectedResult))
		}
	}
}
