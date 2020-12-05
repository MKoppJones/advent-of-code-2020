package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
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

func binaryToDecimal(binary string) int64 {
	var result int64 = 0
	var pow float64 = 0

	for i := len(binary) - 1; i >= 0; i-- {
		r := math.Pow(2, pow)

		i, _ := strconv.ParseInt(string(binary[i]), 10, 64)
		result += int64(r) * i

		pow++
	}

	return result
}

func calculateSeatID(seatDefinition string) int {
	output := "0"
	rowBin := seatDefinition[:7]
	colBin := seatDefinition[len(seatDefinition)-3:]

	output += strings.ReplaceAll(strings.ReplaceAll(rowBin, "F", "0"), "B", "1")
	row := binaryToDecimal(output)

	output = "00000"
	output += strings.ReplaceAll(strings.ReplaceAll(colBin, "L", "0"), "R", "1")
	col := binaryToDecimal(output)

	seatID := row*8 + col
	return int(seatID)
}

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}

	maxSeatID := 0
	seats := []int{}
	for _, seatDefinition := range input {
		seatID := calculateSeatID(seatDefinition)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
		seats = append(seats, seatID)
	}

	sort.Ints(seats)
	currentSeat := 0
	for i := 0; i < len(seats)-1; i++ {
		currentSeat = seats[i]
		nextSeat := seats[i+1]

		if currentSeat+2 == nextSeat {
			currentSeat++
			break
		}
	}

	fmt.Printf("Day 05 - Binary Boarding\nPart 1:\t%d\nPart 2:\t%d\n\n", maxSeatID, currentSeat)
}
