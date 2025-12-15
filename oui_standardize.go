package main

import (
	"os"
	"io"
	"encoding/csv"
	"strings"
	"sort"
)

var data []string

func readCSV(file string) {

	raw, err := os.Open(file)
	if err != nil {panic(err)}
	defer raw.Close()

	reader := csv.NewReader(raw)

	// Skip header line
	reader.Read()

	for {
		// Read CSV line
		line, err := reader.Read()

		// Validate line
		if err == io.EOF {break}
		if err != nil {panic(err)}

		// Remove multi-line strings from addresses
		line[3] = strings.Replace(line[3], "\r", "", -1)
		line[3] = strings.Replace(line[3], "\n", " ", -1)

		// Null-separated strings, starting with mac prefix
		data = append(data, strings.Join([]string{line[1], line[0], line[2], line[3]}, "\x00"))
	}

}

func help() {

	os.Stdout.WriteString("Usage: oui_standardize [<file>] [<file>] [<file>]...\n")
	os.Exit(0)

}

func main() {

	//Display help and exit if not enough arguments
	if len(os.Args) < 2 {help()}

	// Treat arguments as file paths and read all files into data
	for i := 1; i < len(os.Args); i++ {readCSV(os.Args[i])}

	// Sort data
	sort.Strings(data)

	// Format and output
	for _, line := range data {
		fields := strings.Split(line, "\x00")
		os.Stdout.WriteString(
			"{"+fields[1]+"}{"+fields[0]+"}{"+fields[2]+"}{"+fields[3]+"}"+EOL)
	}

}