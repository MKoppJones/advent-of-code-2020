package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) ([]int, error) {
	nums := []int{}
	bytesRead, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		println(readFileErr)
		return nil, readFileErr
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	lines = strings.Split(lines[0], ",")
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	return nums, nil
}

func replace(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}
	max := 2020
	count := 0
	numbers := map[int][]int{}
	previousNum := 0

	for index, num := range input {
		numbers[num] = append(numbers[num], index+1)
		count = index + 1
		previousNum = num
	}

	for count < max {
		if len(numbers[previousNum]) == 1 {
			numbers[0] = append(numbers[0], count+1)
			previousNum = 0
		} else {
			previousTimes := numbers[previousNum][len(numbers[previousNum])-2:]
			previousNum = previousTimes[1] - previousTimes[0]
			_, exists := numbers[previousNum]
			if !exists {
				numbers[previousNum] = []int{}
			}
			numbers[previousNum] = append(numbers[previousNum], count+1)
		}
		count++
	}

	part1 := previousNum

	max = 30000000

	for count < max {
		if len(numbers[previousNum]) == 1 {
			numbers[0] = append(numbers[0], count+1)
			previousNum = 0
		} else {
			previousTimes := numbers[previousNum][len(numbers[previousNum])-2:]
			previousNum = previousTimes[1] - previousTimes[0]
			_, exists := numbers[previousNum]
			if !exists {
				numbers[previousNum] = []int{}
			}
			numbers[previousNum] = append(numbers[previousNum], count+1)
		}
		count++
	}

	fmt.Printf("Day 15 - Rambunctious Recitation\nPart 1:\t%d\nPart 2:\t%d\n\n", part1, previousNum)
}
