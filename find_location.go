package findlocation

import "strings"

type LocationFinder struct {
	CountriesAndCodes map[string]string	
	CitiesWithPopulationsOver15000 map[string]string	
	AlternativeNamesForCitiesAndCountries map[string]string	
	UsStates []string
	UkCounties []string
	CanadianProvincesAndCodes map[string]string	
	AustralianStatesAndCodes map[string]string	
	UsStateCodes []string
}

func NewLocationFinder() *LocationFinder {
	return &LocationFinder {
		AlternativeNamesForCitiesAndCountries: loadTwoColumnFile("data/alternative_names_for_cities_and_countries.tsv"),
		AustralianStatesAndCodes: loadTwoColumnFile("data/australian_states_and_Codes.tsv"),
		CanadianProvincesAndCodes: loadTwoColumnFile("data/canadian_provinces_and_codes.tsv"),
		CitiesWithPopulationsOver15000: loadTwoColumnFile("data/cities_with_populations_over_15000.tsv"),
		CountriesAndCodes: loadTwoColumnFile("data/countries_and_codes.tsv"),
		UkCounties: loadSingleColumnFile("data/uk_counties.txt"),
		UsStates: loadSingleColumnFile("data/us_states.txt"),
		UsStateCodes: loadSingleColumnFile("data/us_state_codes.txt"),
	}
}

func (lf *LocationFinder) FindLocation(search string) string {
	splitSearchString := strings.Split(search, " ")

	// Search for the presence of a country name
	searchResult := searchMultiColumnByKey(lf.CountriesAndCodes, splitSearchString)
	if searchResult != "" {
		return searchResult
	}

	// Search through a list of cities with populations over 15,000
	searchResult = searchMultiColumnByKey(lf.CitiesWithPopulationsOver15000, splitSearchString)
	if searchResult != "" {
		return searchResult
	}

	// Search through a list alternative/non-English names for cities/countries
	searchResult = searchMultiColumnByKey(lf.AlternativeNamesForCitiesAndCountries, splitSearchString)
	if searchResult != "" {
		return searchResult
	}

	// Search for state/province/county name
	searchResult = searchSingleColumn(lf.UsStates, splitSearchString)
	if searchResult != "" {
		return "US"
	}

	searchResult = searchSingleColumn(lf.UkCounties, splitSearchString)
	if searchResult != "" {
		return "GB"
	}

	searchResult = searchMultiColumnByKey(lf.CanadianProvincesAndCodes, splitSearchString)
	if searchResult != "" {
		return "CA"
	}

	searchResult = searchMultiColumnByKey(lf.AustralianStatesAndCodes, splitSearchString)
	if searchResult != "" {
		return "AU"
	}

	// Search for abbreviations/codes
	searchResult = searchMultiColumnByValue(lf.CountriesAndCodes, splitSearchString)
	if searchResult != "" {
		return searchResult
	}

	searchResult = searchSingleColumn(lf.UsStateCodes, splitSearchString)
	if searchResult != "" {
		return "US"
	}

	searchResult = searchMultiColumnByValue(lf.CanadianProvincesAndCodes, splitSearchString)
	if searchResult != "" {
		return "CA"
	}

	searchResult = searchMultiColumnByValue(lf.AustralianStatesAndCodes, splitSearchString)
	if searchResult != "" {
		return "AU"
	}

	return ""
}

func searchMultiColumnByKey(file map[string]string, searchStrings []string) string {
	for _, segment := range searchStrings {
		result := file[segment]

		if result != "" {
			return result
		}
	}

	return ""
}

func searchMultiColumnByValue(file map[string]string, searchStrings []string) string {
	for _, segment := range searchStrings {
		for _, value := range file {
			if value == segment {
				return value
			}
		}
	}

	return ""
}

func searchSingleColumn(file []string, searchStrings []string) string {
	for _, segment := range searchStrings {
		for _, value := range file {
			if value == segment {
				return value
			}
		}
	}

	return ""
}
