package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) ([]int, error) {
	bytesRead, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		println(readFileErr)
		return nil, readFileErr
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]
	nums := []int{}
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	return nums, nil
}

func contains(arr []int, str int) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}
	preamble := 25
	num := 0
	for i := preamble; i < len(input); i++ {
		prevNumbers := input[i-preamble : i]
		num = input[i]
		found := false
		for _, x := range prevNumbers {
			for _, y := range prevNumbers {
				if y != x {
					if x+y == num {
						found = true
						break
					}
				}
			}
			if found {
				break
			}
		}
		if !found {
			break
		}
	}

	nums := []int{}
	found := false
	counter := 0
	min := math.MaxInt64
	max := 0
	for !found {
		total := 0
		min = math.MaxInt64
		max = 0
		nums = []int{}
		for i := counter; i < len(input); i++ {
			nums = append(nums, input[i])
			total += input[i]
			if min > input[i] {
				min = input[i]
			}
			if max < input[i] {
				max = input[i]
			}
			found = total == num
			if found {
				break
			} else if total > num {
				break
			}
		}
		counter++
	}
	weakness := min + max

	fmt.Printf("Day 09 - Encoding Error\nPart 1:\t%d\nPart 2:\t%d\n\n", num, weakness)
}
