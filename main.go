package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// helloWorld := "Hello world"

	// fmt.Printf("%s\n", helloWorld)
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Error opening %s\n", *csvFilename))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file")
	}
	problems := parseProblems(lines)

	var counter int
	var answer string
	for i, problem := range problems {
		fmt.Printf(
			"Problem number %d out of %d.\n%s = ",
			i+1,
			len(problems),
			problem.question,
		)

		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			counter++
		}
	}

	fmt.Printf("Quiz finished!\nYou scored %d out of %d\n", counter, len(problems))
}

type problem struct {
	question string
	answer   string
}

func parseProblems(lines [][]string) []problem {
	res := make([]problem, len(lines))
	for i, line := range lines {
		res[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return res
}

// No need to cause its already parsed
// func validateLine(line string) bool {
// 	validQuesiton, err := regexp.MatchString(`[!-+\--~]*,`, line)
// 	if err != nil {
// 		exit(fmt.Sprintf("Error validating question in line: %s", line))
// 	}

// 	validAnswer, err := regexp.MatchString(`,[ ]*[!-+\--~]*`, line)
// 	if err != nil {
// 		exit(fmt.Sprintf("Error validating answer in line: %s", line))
// 	}

// 	return validQuesiton && validAnswer
// }

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
