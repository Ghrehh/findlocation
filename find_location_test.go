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
	result := locationFinder.FindLocation("countriesAndCodesKey")
	expectedResult := "countriesAndCodesValue"

	if result != expectedResult {
		t.Error(errorMessage(result, expectedResult))
	}
}

func TestFindLocationCitiesWithPopulationsOver15000(t *testing.T) {
	result := locationFinder.FindLocation("citiesWithPopulationsOver15000Key")
	expectedResult := "citiesWithPopulationsOver15000Value"

	if result != expectedResult {
		t.Error(errorMessage(result, expectedResult))
	}
}

func TestFindLocationAlternativeNamesForCitiesAndCountries(t *testing.T) {
	result := locationFinder.FindLocation("alternativeNamesForCitiesAndCountriesKey")
	expectedResult := "alternativeNamesForCitiesAndCountriesValue"

	if result != expectedResult {
		t.Error(errorMessage(result, expectedResult))
	}
}

func TestFindLocationUsStates(t *testing.T) {
	result := locationFinder.FindLocation("usStates")
	expectedResult := "US"

	if result != expectedResult {
		t.Error(errorMessage(result, expectedResult))
	}
}

func TestFindLocationUkCounties(t *testing.T) {
	result := locationFinder.FindLocation("ukCounties")
	expectedResult := "GB"

	if result != expectedResult {
		t.Error(errorMessage(result, expectedResult))
	}
}

func TestFindLocationCanadianProvinces(t *testing.T) {
	result := locationFinder.FindLocation("canadianProvincesAndCodesKey")
	expectedResult := "CA"

	if result != expectedResult {
		t.Error(errorMessage(result, expectedResult))
	}
}

func TestFindLocationAustralianStatesAndCodes(t *testing.T) {
	result := locationFinder.FindLocation("australianStatesAndCodesKey")
	expectedResult := "AU"

	if result != expectedResult {
		t.Error(errorMessage(result, expectedResult))
	}
}

func TestFindLocationCountriesAndCodesByKey(t *testing.T) {
	result := locationFinder.FindLocation("countriesAndCodesValue")
	expectedResult := "countriesAndCodesValue"

	if result != expectedResult {
		t.Error(errorMessage(result, expectedResult))
	}
}

func TestFindLocationUsStateCodes(t *testing.T) {
	result := locationFinder.FindLocation("usStateCodes")
	expectedResult := "US"

	if result != expectedResult {
		t.Error(errorMessage(result, expectedResult))
	}
}
