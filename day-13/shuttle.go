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
	lines = lines[:len(lines)-1]
	return lines, nil
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// BusTime ...
type BusTime struct {
	val   int
	index int
}

func part1(input []string) int {
	earliest, _ := strconv.Atoi(input[0])
	busses := unique(strings.Split(input[1], ","))
	timestamp := 0
	earliestBuses := map[string]int{}
	for len(earliestBuses) < len(busses)-1 {
		timestamp++
		for _, raw := range busses {
			id, err := strconv.Atoi(raw)
			if err == nil {
				_, exists := earliestBuses[raw]
				if !exists && id*timestamp > earliest {
					earliestBuses[raw] = id * timestamp
				}
			}
		}
	}

	bus := 0
	minVal := earliest * 2
	for id, val := range earliestBuses {
		if val < minVal {
			bus, _ = strconv.Atoi(id)
			minVal = val
		}
	}

	timeDifference := (minVal - earliest)
	return bus * timeDifference
}

func part2(input []string) int {
	timestamp := 0
	busNumbers := strings.Split(input[1], ",")
	increment, _ := strconv.Atoi(busNumbers[0])
	for index := 1; index < len(busNumbers); index++ {
		if string(busNumbers[index]) != "x" {
			nextBus, _ := strconv.Atoi(busNumbers[index])
			for {
				timestamp += increment

				// If we have found the next number that is the previous plus index,
				// we have found the point at which the pattern repeats
				//
				// This means we can multiple our step by this bus number.
				// This brings the time of this to something more reasonable.
				if (timestamp+index)%nextBus == 0 {
					increment *= nextBus
					break
				}
			}
		}
	}
	return timestamp
}

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}

	fmt.Printf("Day 13 - Shuttle Search\nPart 1:\t%d\nPart 2:\t%d\n\n", part1(input), part2(input))
}
