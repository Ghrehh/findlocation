package main

import (
	"strings"
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/csv"
)

func loadSingleColumnFile(fileName string) []string {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		errString := []string{"error opening file '", fileName, "'"};
    fmt.Println(strings.Join(errString, " "));
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	output := make([]string, 0)

	for _, line := range lines {
		output = append(output, line)
	}

	return output
}

func loadTwoColumnFile(fileName string) map[string]string {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		errString := []string{"error opening file '", fileName, "'"};
    fmt.Println(strings.Join(errString, " "));
		os.Exit(1)
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
			errString := []string{"error parsing file '", fileName, "'"};
			fmt.Println(strings.Join(errString, " "));
			os.Exit(1)
		}

		output[record[1]] = record[0]
	}

	return output
}
