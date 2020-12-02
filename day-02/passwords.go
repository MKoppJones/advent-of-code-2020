package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readPasswordPolicies() ([]string, error) {
	policies := []string{}

	bytesRead, readFileErr := ioutil.ReadFile("passwords.txt")
	if readFileErr != nil {
		println(readFileErr)
		return []string{}, readFileErr
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	policies = lines[:len(lines)-1]
	return policies, nil
}

func main() {
	passwords, readErr := readPasswordPolicies()
	if readErr != nil {
		println(readErr)
		return
	}

	validPasswords := 0
	for _, policy := range passwords {
		splitPolicy := strings.Split(policy, " ")
		password := splitPolicy[2]
		letter := strings.TrimRight(splitPolicy[1], ":")
		counts := strings.Split(splitPolicy[0], "-")
		min, minError := strconv.Atoi(counts[0])
		if minError != nil {
			println(minError)
			return
		}
		max, maxError := strconv.Atoi(counts[1])
		if maxError != nil {
			println(maxError)
			return
		}

		letterCount := strings.Count(password, letter)
		if letterCount >= min && letterCount <= max {
			validPasswords++
		}
	}
	println(validPasswords)
}
