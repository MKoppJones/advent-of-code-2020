package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var countedChain map[int]int

func readInput(path string) ([]int, error) {
	bytesRead, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		println(readFileErr)
		return nil, readFileErr
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]
	jolts := []int{}
	for _, line := range lines {
		jolt, _ := strconv.Atoi(line)
		jolts = append(jolts, jolt)
	}
	return jolts, nil
}

func contains(arr []([]int), str []int) bool {
	for _, a := range arr {
		if reflect.DeepEqual(a, str) {
			return true
		}
	}
	return false
}

func walkChain(input []int) bool {
	for i := 1; i < len(input); i++ {
		adapterJoltage := input[i]
		prevJoltage := input[i-1]
		diff := adapterJoltage - prevJoltage
		if !(diff <= 3 && diff >= 1) {
			return false
		}
	}
	return true
}

// func nextChain(chain []int, checked []([]int), count int) (int, []([]int)) {
// 	for i := 1; i < len(chain)-1; i++ {
// 		newChain := remove(append([]int(nil), chain...), i)
// 		if !contains(checked, newChain) {
// 			checked = append(checked, newChain)
// 			if walkChain(newChain) {
// 				count++
// 				count, checked = nextChain(newChain, checked, count)
// 			}
// 		}
// 	}
// 	return count, checked
// }

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}

	prevJoltage := 0
	chain := []int{0}
	sort.Ints(input)
	highDiff := 1
	lowDiff := 0
	for _, adapterJoltage := range input {
		diff := adapterJoltage - prevJoltage
		if diff <= 3 && diff >= 1 {
			if diff == 3 {
				highDiff++
			}
			if diff == 1 {
				lowDiff++
			}
			chain = append(chain, adapterJoltage)
			prevJoltage = adapterJoltage
		}
	}
	multi := highDiff * lowDiff
	chain = append(chain, chain[len(chain)-1]+3)
	ans := 0
	ans = countChain(chain)
	fmt.Printf("Day 10 - Adapter Array\nPart 1:\t%d\nPart 2:\t%d\n\n", multi, ans)
}

func countChain(input []int) int {
	countedChain = make(map[int]int)

	chain := make(map[int]bool)

	for _, joltage := range input {
		chain[joltage] = true
	}

	chain[input[len(input)-1]+3] = true

	return countPathsFrom(0, chain)
}

func countPathsFrom(start int, chain map[int]bool) int {
	count, visited := countedChain[start]

	if visited {
		return count
	}

	counted := 0
	noCandidate := true

	for i := 1; i <= 3; i++ {
		candidate := start + i
		_, ok := chain[candidate]
		if ok {
			counted += countPathsFrom(candidate, chain)
			noCandidate = false
		}
	}

	if noCandidate {
		counted++
	}

	countedChain[start] = counted

	return counted
}
