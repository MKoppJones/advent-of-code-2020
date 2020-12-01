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

func main() {
	expenses := readExpenses()
	value1 := 0
	value2 := 0
	found := false

	for _, element := range expenses {
		value1 = element

		for _, element2 := range expenses {
			value2 = element2
			found = (value1+value2 == 2020)

			if found {
				break
			}
		}

		if found {
			break
		}
	}
	println(value1 * value2)
}
