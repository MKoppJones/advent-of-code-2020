package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type Point struct {
	x int
	y int
}

var floor = "."
var emptySeat = "L"
var occupiedSeat = "#"
var width int = 0
var height int = 0

func readInput(path string) (string, error) {
	bytesRead, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		println(readFileErr)
		return "", readFileErr
	}
	fileContent := string(bytesRead)
	return fileContent, nil
}

func contains(arr []([]int), str []int) bool {
	for _, a := range arr {
		if reflect.DeepEqual(a, str) {
			return true
		}
	}
	return false
}

func index(x int, y int) int {
	return x + (y * width)
}

func step(state string) string {
	prevState := state
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currIndex := index(x, y)
			curr := string(state[currIndex])
			if curr == floor {
				continue
			}
			occupied := 0
			if y-1 >= 0 { // Top
				if string(prevState[index(x, y-1)]) == occupiedSeat {
					occupied++
				}
			}

			if y+1 < height { // Bottom
				if string(prevState[index(x, y+1)]) == occupiedSeat {
					occupied++
				}
			}

			if x-1 >= 0 { // Left
				if string(prevState[index(x-1, y)]) == occupiedSeat {
					occupied++
				}
			}

			if x+1 < width { // Right
				if string(prevState[index(x+1, y)]) == occupiedSeat {
					occupied++
				}
			}

			if y-1 >= 0 && x-1 >= 0 { // Top Left
				if string(prevState[index(x-1, y-1)]) == occupiedSeat {
					occupied++
				}
			}

			if y+1 < height && x-1 >= 0 { // Bottom Left
				if string(prevState[index(x-1, y+1)]) == occupiedSeat {
					occupied++
				}
			}

			if y-1 >= 0 && x+1 < width { // Top Right
				if string(prevState[index(x+1, y-1)]) == occupiedSeat {
					occupied++
				}
			}

			if y+1 < height && x+1 < width { // Bottom Right
				if string(prevState[index(x+1, y+1)]) == occupiedSeat {
					occupied++
				}
			}

			if curr == emptySeat {
				if occupied == 0 {
					state = replace(state, occupiedSeat, index(x, y))
				}
			}

			if curr == occupiedSeat {
				if occupied >= 4 {
					state = replace(state, emptySeat, index(x, y))
				}
			}
		}
	}
	return state
}

func firstSeatedOnSlope(startX int, startY int, state string, slopeX int, slopeY int) bool {
	x := startX + slopeX
	y := startY + slopeY
	for {
		if x < 0 || x >= width || y < 0 || y >= height {
			return false
		}
		seat := string(state[index(x, y)])
		if seat == occupiedSeat {
			return true
		} else if seat == emptySeat {
			return false
		}
		x += slopeX
		y += slopeY
	}
}

func stepSight(state string) string {
	prevState := state
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currIndex := index(x, y)
			curr := string(state[currIndex])
			if curr == floor {
				continue
			}

			occupied := 0
			if firstSeatedOnSlope(x, y, prevState, 0, -1) {
				occupied++
			}

			if firstSeatedOnSlope(x, y, prevState, 0, 1) {
				occupied++
			}

			if firstSeatedOnSlope(x, y, prevState, -1, 0) {
				occupied++
			}

			if firstSeatedOnSlope(x, y, prevState, 1, 0) {
				occupied++
			}

			if firstSeatedOnSlope(x, y, prevState, -1, -1) {
				occupied++
			}

			if firstSeatedOnSlope(x, y, prevState, -1, 1) {
				occupied++
			}

			if firstSeatedOnSlope(x, y, prevState, 1, -1) {
				occupied++
			}

			if firstSeatedOnSlope(x, y, prevState, 1, 1) {
				occupied++
			}

			if curr == emptySeat {
				if occupied == 0 {
					state = replace(state, occupiedSeat, index(x, y))
				}
			}

			if curr == occupiedSeat {
				if occupied >= 5 {
					state = replace(state, emptySeat, index(x, y))
				}
			}
		}
	}
	return state
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
	flatMap := strings.ReplaceAll(input, "\n", "")
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	width = len(lines[0])
	height = len(lines)
	state := flatMap
	for {
		newState := step(state)
		if state == newState {
			break
		}
		state = newState
	}

	occupied := 0
	for i := 0; i < len(state); i++ {
		if string(state[i]) == occupiedSeat {
			occupied++
		}
	}

	state = flatMap
	for {
		newState := stepSight(state)
		if state == newState {
			break
		}
		state = newState
	}
	occupied2 := 0
	for i := 0; i < len(state); i++ {
		if string(state[i]) == occupiedSeat {
			occupied2++
		}
	}
	fmt.Printf("Day 11 - Seating System\nPart 1:\t%d\nPart 2:\t%d\n\n", occupied, occupied2)
}
