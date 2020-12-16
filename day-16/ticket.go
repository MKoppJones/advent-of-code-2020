package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) ([]string, error) {
	bytesRead, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		println(readFileErr)
		return nil, readFileErr
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	return lines, nil
}

func replace(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

// Range ...
type Range struct {
	low  int
	high int
}

// Rule ...
type Rule struct {
	name   string
	range1 Range
	range2 Range
	index  int
}

func part1() int {
	sum := 0
	for _, ticket := range nearbyTickets[1:] {
		validTicket := true
		for _, num := range ticket {
			valid := false
			for _, rule := range rules {
				if rulecheck(rule, num) {
					valid = true
					break
				}
			}
			if !valid {
				validTicket = false
				sum += num
			}
		}
		if validTicket {
			validTickets = append(validTickets, ticket)
		}
	}
	return sum
}

func part2() int {
	allRuleMatches := map[int]([]int){}

	for _, ticket := range validTickets {
		ruleMatches := []int{}
		for range rules {
			ruleMatches = append(ruleMatches, 0)
		}

		for numIndex, num := range ticket {
			_, exists := allRuleMatches[numIndex]
			if !exists {
				allRuleMatches[numIndex] = append([]int(nil), ruleMatches...)
			}
			for ruleIndex, rule := range rules {
				if rulecheck(rule, num) {
					allRuleMatches[numIndex][ruleIndex]++
				}
			}
		}
	}

	fmt.Println(allRuleMatches)
	for key := range allRuleMatches {
		index := findHighestIndex(allRuleMatches[key])
		rules[index].index = key
		emptyAtIndex(allRuleMatches, index)
	}

	sum := 1
	for _, rule := range rules {
		if strings.HasPrefix(rule.name, "departure") {
			sum *= myTicket[rule.index]
		}
	}
	return sum
}

func emptyAtIndex(ruleMatches map[int]([]int), index int) {
	for key := range ruleMatches {
		ruleMatches[key][index] = 0
	}
}

func findHighestIndex(value []int) int {
	index := 0
	maxVal := 0
	for i, v := range value {
		if v >= maxVal {
			maxVal = v
			index = i
		}
	}
	return index
}

var rules []Rule
var myTicket []int
var nearbyTickets []([]int)
var validTickets []([]int)

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}

	rules = []Rule{}
	myTicket = []int{}
	nearbyTickets = []([]int){}

	x := 0
	for _, line := range input {
		if line == "" {
			x++
			continue
		}

		switch x {
		case 0:
			parts := strings.Split(line, ": ")
			parts = strings.Split(parts[1], " or ")
			range1Parts := strings.Split(parts[0], "-")
			range2Parts := strings.Split(parts[1], "-")

			x1, _ := strconv.Atoi(range1Parts[0])
			y1, _ := strconv.Atoi(range1Parts[1])
			x2, _ := strconv.Atoi(range2Parts[0])
			y2, _ := strconv.Atoi(range2Parts[1])

			range1 := Range{x1, y1}
			range2 := Range{x2, y2}
			rules = append(rules, Rule{strings.Split(line, ": ")[0], range1, range2, 0})
		case 1:
			for _, str := range strings.Split(line, ",") {
				x, _ := strconv.Atoi(str)
				myTicket = append(myTicket, x)
			}
		case 2:
			nums := []int{}
			for _, str := range strings.Split(line, ",") {
				x, _ := strconv.Atoi(str)
				nums = append(nums, x)
			}
			nearbyTickets = append(nearbyTickets, nums)
		}
	}

	partOne := part1()
	partTwo := part2()

	fmt.Printf("Day 16 - Ticket Translations\nPart 1:\t%d\nPart 2:\t%d\n\n", partOne, partTwo)
}

func rulecheck(rule Rule, val int) bool {
	return (rule.range1.low <= val && rule.range1.high >= val) || (rule.range2.low <= val && rule.range2.high >= val)
}
