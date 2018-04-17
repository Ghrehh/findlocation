package findlocation

import (
	"strings"
	"log"
	"io"
	"encoding/csv"
)

func loadSingleColumnFile(fileName string) []string {
	data, err := Asset(fileName)

	if err != nil {
    log.Fatal("error opening file '" + fileName +"'")
	}

	lines := strings.Split(string(data), "\n")
	output := make([]string, 0)

	for _, line := range lines {
		output = append(output, line)
	}

	return output
}

func loadTwoColumnFile(fileName string) map[string]string {
	data, err := Asset(fileName)

	if err != nil {
    log.Fatal("error opening file '" + fileName +"'")
	}

	r := csv.NewReader(strings.NewReader(string(data)))
	r.Comma = '\t'

	output := make(map[string]string)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
    	log.Fatal("error parsing file '" + fileName +"'")
		}

		output[record[1]] = record[0]
	}

	return output
}
