package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Flags
	csvFilename := flag.String("csv", "problems.csv", "a csv filename in the format of 'question,answer'")
	flag.Parse()

	// Read file from OS
	file, err := os.Open(*csvFilename)

	// Catch error
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}

	// Init CSV encoder
	r := csv.NewReader(file)
	// Try read from r
	lines, err := r.ReadAll()

	// Catch error
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	// Do make new array
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		var answer string // Declaring 'answer' var

		// Get input from user input and write to answer var
		fmt.Scanf("%s\n", &answer)

		// Correctly?
		if answer == p.a {
			correct++ // fmt.Println("Correct!")
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

// Parsing/Formatting data to structs' array
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	// fmt.Println(lines)

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

// Declaring problem struct
type problem struct {
	q string
	a string
}

// Catch logger
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
