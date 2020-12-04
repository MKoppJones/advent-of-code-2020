package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readPassports(path string) (string, error) {
	passports := ""

	bytesRead, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		println(readFileErr)
		return "", readFileErr
	}
	passports = string(bytesRead)
	return passports, nil
}

func processBatch(batch string, fields []string) int {
	validPassports := 0

	lines := strings.Split(batch, "\n")
	// lines = lines[:len(lines)-1]

	fieldMap := map[string]int{}
	for _, line := range lines {
		if len(line) == 0 {
			if len(fieldMap) == len(fields) {
				validPassports++
			}
			fieldMap = map[string]int{}
		} else {
			pairs := strings.Split(line, " ")
			for _, pair := range pairs {
				keyValue := strings.Split(pair, ":")
				_, exists := fieldMap[keyValue[0]]
				if keyValue[0] == "cid" {
					continue
				}
				if !exists {
					fieldMap[keyValue[0]] = 0
				}
				fieldMap[keyValue[0]]++
			}
		}
	}

	return validPassports
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	passportBatch, err := readPassports(args[0])
	if err != nil {
		println(err)
		return
	}

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	validPassports := processBatch(passportBatch, fields)

	fmt.Printf("Day 04 - Passport Processing\nPart 1:\t%d\nPart 2:\t%d\n\n", validPassports, 0)
}
