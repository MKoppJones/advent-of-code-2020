package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readPasswordPolicies(path string) ([]string, error) {
	policies := []string{}
	bytesRead, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		println(readFileErr)
		return nil, readFileErr
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	policies = lines[:len(lines)-1]
	return policies, nil
}

func oldPasswordPolicy(policyParts []string) bool {
	password := policyParts[2]
	letter := strings.TrimRight(policyParts[1], ":")
	counts := strings.Split(policyParts[0], "-")
	min, minError := strconv.Atoi(counts[0])
	if minError != nil {
		println(minError)
		return false
	}
	max, maxError := strconv.Atoi(counts[1])
	if maxError != nil {
		println(maxError)
		return false
	}

	letterCount := strings.Count(password, letter)
	return letterCount >= min && letterCount <= max
}

func newPasswordPolicy(policyParts []string) bool {
	password := policyParts[2]
	letter := strings.TrimRight(policyParts[1], ":")
	counts := strings.Split(policyParts[0], "-")
	positionOne, posOneError := strconv.Atoi(counts[0])
	if posOneError != nil {
		println(posOneError)
		return false
	}
	positionTwo, posTwoError := strconv.Atoi(counts[1])
	if posTwoError != nil {
		println(posTwoError)
		return false
	}

	posOneCheck := false
	if positionOne-1 < len(password) {
		posOneCheck = string(password[positionOne-1]) == letter
	}

	posTwoCheck := false
	if positionTwo-1 < len(password) {
		posTwoCheck = string(password[positionTwo-1]) == letter
	}
	return posOneCheck != posTwoCheck
}

func main() {
	args := os.Args[1:]
	passwords, readErr := readPasswordPolicies(args[0])
	if readErr != nil {
		println(readErr)
		return
	}

	oldValidPasswords := 0
	newValidPasswords := 0
	for _, policy := range passwords {
		policyParts := strings.Split(policy, " ")
		if oldPasswordPolicy(policyParts) {
			oldValidPasswords++
		}
		if newPasswordPolicy(policyParts) {
			newValidPasswords++
		}
	}
	fmt.Printf("Day 02 - Password Philosophy\nPart 1:\t%d\nPart 2:\t%d\n\n", oldValidPasswords, newValidPasswords)
}
