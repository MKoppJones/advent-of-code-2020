package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readExpenses(path string) ([]int, error) {
	expenses := []int{}

	bytesRead, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		println(readFileErr)
		return nil, readFileErr
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]

	for _, i := range lines {
		j, atoiError := strconv.Atoi(i)
		if atoiError != nil {
			return nil, atoiError
		}
		expenses = append(expenses, j)
	}
	return expenses, nil
}

func findValues(count int, goal int, items []int) []int {
	arr := make([]int, count)
	return iterateValues(0, count, goal, items, arr)
}

func iterateValues(curr int, max int, goal int, items []int, result []int) []int {
	for _, element := range items {
		result[curr] = element
		if curr < max-1 {
			result = iterateValues(curr+1, max, goal, items, result)
		}
		total := 0
		for _, i := range result {
			total += i
		}
		if total == goal {
			break
		}
	}
	return result
}

func multi(values []int) int {
	total := 1
	for _, i := range values {
		total = total * i
	}
	return total
}

func main() {
	args := os.Args[1:]
	expenses, err := readExpenses(args[0])
	if err != nil {
		println(err)
		return
	}
	part1 := multi(findValues(2, 2020, expenses))
	part2 := multi(findValues(3, 2020, expenses))
	fmt.Printf("Day 01 - Report Repair\nPart 1:\t%d\nPart 2:\t%d\n\n", part1, part2)
}
