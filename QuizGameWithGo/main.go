package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// flag the file
	csvFilename := flag.String("csv", "problems.csv", "a csv in a form of 'question,answer'")
	flag.Parse()
	// creating a timelimit (flag)
	timelimit := flag.Int("limit", 1, "the time limit for the quiz in seconds")
	flag.Parse()

	// opening the file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file %s", *csvFilename))
	}

	// reading the file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	correct := 0
problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}

// exit function is for closing the program and exit
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
