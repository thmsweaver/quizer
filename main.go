package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

const defaultFilename = "problems.csv"

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String(
		"csv",
		defaultFilename,
		"file containing questions and answers",
	)
	flag.Parse()

	f, err := os.Open(*csvFilename)
	if err != nil {
		panic(fmt.Errorf("failed to open CSV file: %s", *csvFilename))
	}

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		tmplt := "Unable to parse CSV: %s, error: %s"
		panic(fmt.Errorf(tmplt, csvFilename, err.Error()))
	}

	var correct int
	problems := parseLines(lines)

	for i, p := range problems {
		fmt.Printf("%d: %s\n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("you answered %d out %d correctly\n", correct, len(problems))
}

func parseLines(lines [][]string) []*problem {
	ret := make([]*problem, len(lines))
	for i, line := range lines {
		ret[i] = &problem{q: line[0], a: strings.TrimSpace(line[1])}
	}
	return ret
}
