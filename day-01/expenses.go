package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readExpenses() []int {
	expenses := []int{}

	bytesRead, _ := ioutil.ReadFile("expenses.txt")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]

	for _, i := range lines {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		expenses = append(expenses, j)
	}
	return expenses
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
	expenses := readExpenses()
	println(multi(findValues(2, 2020, expenses)))
	println(multi(findValues(3, 2020, expenses)))
}
