package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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
	lines = lines[:len(lines)-1]
	return lines, nil
}

func part1(input []string) int {
	sum := 0
	for _, formula := range input {
		formula = strings.ReplaceAll(formula, " ", "")
		sum += parseFormula(formula)
	}
	return sum
}

func walk(valStack []int, opStack []string, close bool) ([]int, []string) {
	op := ""
	for len(opStack) > 0 && len(valStack) > 0 {
		if opStack[len(opStack)-1] == "(" {
			if close {
				op, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
			}
			break
		}
		op, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
		rhs := 0
		lhs := 0

		rhs, valStack = valStack[len(valStack)-1], valStack[:len(valStack)-1]
		lhs, valStack = valStack[len(valStack)-1], valStack[:len(valStack)-1]
		res := 0
		switch op {
		case "+":
			res = lhs + rhs
		case "*":
			res = lhs * rhs
		}
		valStack = append(valStack, res)
	}
	return valStack, opStack
}

func parseFormula(formula string) int {
	formulaParts := strings.Split(formula, "")

	valStack := []int{}
	opStack := []string{}

	doFetch := true
	for i := 0; i < len(formulaParts); i++ {
		part := formulaParts[i]
		if part == "(" {
			opStack = append(opStack, part)
		} else if part == ")" {
			valStack, opStack = walk(valStack, opStack, true)
		} else {
			if doFetch {
				v, _ := strconv.Atoi(part)
				valStack = append(valStack, v)
			} else {
				if part == "+" || part == "*" {
					valStack, opStack = walk(valStack, opStack, false)
				}
				opStack = append(opStack, part)
			}
			doFetch = !doFetch
		}
	}

	valStack, opStack = walk(valStack, opStack, false)
	return valStack[len(valStack)-1]
}

func part2(input []string) int {
	regexp, _ := regexp.Compile(`\d\*\d`)
	sum := 0
	for _, formula := range input {
		formula = strings.ReplaceAll(formula, " ", "")
		formula = regexp.ReplaceAllStringFunc(formula, func(m string) string {
			parts := regexp.FindStringSubmatch(m)
			return "(" + parts[0] + ")"
		})
		sum += parseFormula(formula)
		fmt.Println(formula, parseFormula(formula))
	}
	return sum
}

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}

	partOne := part1(input)
	partTwo := part2(input)

	fmt.Printf("Day 18 - Operation Order\nPart 1:\t%d\nPart 2:\t%d\n\n", partOne, partTwo)
}
