package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Instruction ...
type Instruction struct {
	instruction string
	value       int
}

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

func contains(arr []int, str int) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func run(instructions []Instruction) (accumulator int, halted bool, instructionsVisited []int) {
	pointer := 0
	accumulator = 0
	for !contains(instructionsVisited, pointer) && pointer < len(instructions) {
		instructionsVisited = append(instructionsVisited, pointer)
		switch instructions[pointer].instruction {
		case "nop":
			pointer++
		case "acc":
			accumulator += instructions[pointer].value
			pointer++
		case "jmp":
			pointer += instructions[pointer].value
		}
	}

	if contains(instructionsVisited, pointer) {
		halted = false
	} else {
		halted = true
	}
	return
}

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}

	is := []Instruction{}
	for _, value := range input {
		split := strings.Split(value, " ")
		number, _ := strconv.Atoi(split[1])
		is = append(is, Instruction{split[0], number})
	}

	accumulator1, _, visited := run(is)

	accumulator2 := 0
	halted := false
	for i := len(visited) - 1; i >= 0; i-- {
		pointer := visited[i]

		newInstructions := make([]Instruction, len(is))
		copy(newInstructions, is)
		if newInstructions[pointer].instruction != "acc" {

			if newInstructions[pointer].instruction == "nop" {
				newInstructions[pointer] = Instruction{"jmp", newInstructions[pointer].value}
			} else if newInstructions[pointer].instruction == "jmp" {
				newInstructions[pointer] = Instruction{"nop", newInstructions[pointer].value}
			}
			accumulator2, halted, _ = run(newInstructions)

			if halted {
				break
			}
		}
	}

	fmt.Printf("Day 08 - Handheld Halting\nPart 1:\t%d\nPart 2:\t%d\n\n", accumulator1, accumulator2)
}
