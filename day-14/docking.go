package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

func setBit(n uint, pos uint) uint {
	n |= (1 << pos)
	return uint(n)
}

func clearBit(n uint, pos uint) uint {
	return uint(n &^ (1 << pos))
}

func setMemoryV1(index int, value int, mask string) {
	max := 36
	maskedValue := uint(value)
	for i := 0; i < len(mask); i++ {
		switch string(mask[i]) {
		case "1":
			maskedValue = uint(setBit(maskedValue, uint(max-i-1)))
		case "0":
			maskedValue = uint(clearBit(maskedValue, uint(max-i-1)))
		}
	}
	memory[index] = maskedValue
}

func replace(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func setMemoryV2(index int, value int, mask string) {
	bitString := fmt.Sprintf("%036b", index)
	xes := 0
	for i := 0; i < len(mask); i++ {
		switch string(mask[i]) {
		case "X":
			xes++
			fallthrough
		case "1":
			bitString = replace(bitString, string(mask[i]), i)
		}
	}

	max := int(math.Pow(2, float64(xes)))
	count := 0
	for count < max {
		index := 0
		newAddress := fmt.Sprintf("%036b", 0)
		for i := 0; i < len(bitString); i++ {
			if string(bitString[i]) == "X" {
				bitSet := count&(1<<index) != 0
				newBit := "0"
				if bitSet {
					newBit = "1"
				}
				newAddress = replace(newAddress, newBit, i)
				index++
			} else {
				newAddress = replace(newAddress, string(bitString[i]), i)
			}
		}

		v, _ := strconv.ParseInt(newAddress, 2, 64)

		memory[int(v)] = uint(value)

		count++
	}
}

var memory map[int]uint

func runDecorder1(commandLines []string) map[int]uint {
	memory = map[int]uint{}

	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	for _, line := range commandLines {
		lineParts := strings.Split(line, " ")
		if lineParts[0] == "mask" {
			mask = lineParts[2]
		} else {
			strMemIndex := strings.Split(lineParts[0], "[")[1]
			strMemIndex = strMemIndex[:len(strMemIndex)-1]
			memIndex, _ := strconv.Atoi(strMemIndex)
			value, _ := strconv.Atoi(lineParts[2])
			setMemoryV1(memIndex, value, mask)
		}
	}
	return memory
}

func runDecorder2(commandLines []string) map[int]uint {
	memory = map[int]uint{}

	mask := ""
	for _, line := range commandLines {
		lineParts := strings.Split(line, " ")
		if lineParts[0] == "mask" {
			mask = lineParts[2]
		} else {
			strMemIndex := strings.Split(lineParts[0], "[")[1]
			strMemIndex = strMemIndex[:len(strMemIndex)-1]
			memIndex, _ := strconv.Atoi(strMemIndex)
			value, _ := strconv.Atoi(lineParts[2])
			setMemoryV2(memIndex, value, mask)
		}
	}
	return memory
}

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}

	mem1 := runDecorder1(input)

	mem2 := runDecorder2(input)

	fmt.Printf("Day 14 - Docking Data\nPart 1:\t%d\nPart 2:\t%d\n\n", sum(mem1), sum(mem2))
}

func sum(mem map[int]uint) int {
	memSum := 0
	for _, value := range mem {
		memSum += int(value)
	}
	return memSum
}
