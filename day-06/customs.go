package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readInput(path string) ([]([]string), error) {
	bytesRead, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		println(readFileErr)
		return nil, readFileErr
	}
	fileContent := string(bytesRead)
	allLines := strings.Split(fileContent, "\n")
	lines := []([]string){}
	groups := []string{}
	for _, line := range allLines {
		if len(line) != 0 {
			groups = append(groups, line)
		} else {
			lines = append(lines, groups)
			groups = []string{}
		}
	}

	return lines, nil
}

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}

	questionCount := 0
	allQuestionCount := 0
	for _, group := range input {
		questions := map[rune]int{}
		for _, person := range group {
			for _, answer := range person {
				_, exists := questions[answer]
				if !exists {
					questions[answer] = 0
				}
				questions[answer]++
			}
		}
		for _, question := range questions {
			if question == len(group) {
				allQuestionCount++
			}
		}
		questionCount += len(questions)
	}

	fmt.Printf("Day 06 - Custom Customs\nPart 1:\t%d\nPart 2:\t%d\n\n", questionCount, allQuestionCount)
}
