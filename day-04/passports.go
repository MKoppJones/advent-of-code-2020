package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
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

func insecureProcessBatch(batch string, fields []string) int {
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

func secureProcessBatch(batch string, fields []string) int {
	validPassports := 0

	lines := strings.Split(batch, "\n")

	fieldMap := map[string]int{}
	for _, line := range lines {
		if len(line) == 0 {
			// jsonString, _ := json.Marshal(fieldMap)

			// println(string(jsonString), len(fieldMap), len(fields))
			if len(fieldMap) == len(fields) {
				validPassports++
			}
			fieldMap = map[string]int{}
		} else {
			pairs := strings.Split(line, " ")
			for _, pair := range pairs {
				keyValue := strings.Split(pair, ":")
				_, exists := fieldMap[keyValue[0]]
				switch keyValue[0] {
				case "byr":
					value, err := strconv.Atoi(keyValue[1])
					if err != nil {
						println(err.Error, "byr", keyValue[1])
					} else if value >= 1920 && value <= 2002 {
						if !exists {
							fieldMap[keyValue[0]] = 0
						}
						fieldMap[keyValue[0]]++
					}
				case "iyr":
					value, err := strconv.Atoi(keyValue[1])
					if err != nil {
						println(err.Error, "iyr", keyValue[1])
					} else if value >= 2010 && value <= 2020 {
						if !exists {
							fieldMap[keyValue[0]] = 0
						}
						fieldMap[keyValue[0]]++
					}
				case "eyr":
					value, err := strconv.Atoi(keyValue[1])
					if err != nil {
						println(err.Error, "eyr", keyValue[1])
					} else if value >= 2020 && value <= 2030 {
						if !exists {
							fieldMap[keyValue[0]] = 0
						}
						fieldMap[keyValue[0]]++
					}
				case "hgt":
					if strings.HasSuffix(keyValue[1], "cm") {
						value, err := strconv.Atoi(strings.ReplaceAll(keyValue[1], "cm", ""))
						if err != nil {
							println(err.Error, "hgt cm", keyValue[1])
						} else if value >= 150 && value <= 193 {
							if !exists {
								fieldMap[keyValue[0]] = 0
							}
							fieldMap[keyValue[0]]++
						}
					} else if strings.HasSuffix(keyValue[1], "in") {
						value, err := strconv.Atoi(strings.ReplaceAll(keyValue[1], "in", ""))
						if err != nil {
							println(err.Error, "hgt in", keyValue[1])
						} else if value >= 59 && value <= 76 {
							println(exists)
							if !exists {
								fieldMap[keyValue[0]] = 0
							}
							fieldMap[keyValue[0]]++
						}
					}
				case "hcl":
					found, err := regexp.MatchString("^#[0-9a-f]{6}$", keyValue[1])
					if err != nil {
						println(err.Error, "hcl", keyValue[1])
					} else if found {
						if !exists {
							fieldMap[keyValue[0]] = 0
						}
						fieldMap[keyValue[0]]++
					}
				case "ecl":
					if contains([]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}, keyValue[1]) {
						if !exists {
							fieldMap[keyValue[0]] = 0
						}
						fieldMap[keyValue[0]]++
					}
				case "pid":
					found, err := regexp.MatchString("^[0-9]{9}$", keyValue[1])
					if err != nil {
						println(err.Error, "pid", keyValue[1])
					} else if found {
						if !exists {
							fieldMap[keyValue[0]] = 0
						}
						fieldMap[keyValue[0]]++
					}
				}
			}
		}
	}

	return validPassports
}

func main() {
	args := os.Args[1:]
	passportBatch, err := readPassports(args[0])
	if err != nil {
		println(err)
		return
	}

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	insecureValidPassports := insecureProcessBatch(passportBatch, fields)
	secureValidPassports := secureProcessBatch(passportBatch, fields)

	fmt.Printf("Day 04 - Passport Processing\nPart 1:\t%d\nPart 2:\t%d\n\n", insecureValidPassports, secureValidPassports)
}
