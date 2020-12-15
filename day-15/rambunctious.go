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
	fmt.Println(input)
	fmt.Printf("Day 15 - Rambunctious Recitation\nPart 1:\t%d\nPart 2:\t%d\n\n", 0, 0)
}
