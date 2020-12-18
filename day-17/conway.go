package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	return lines, nil
}

func replace(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func part1() int {
	cycle := 0

	for cycle < 1 {
		deepCopyMap(fullState, prevFullState)
		for point := range prevFullState {
			modifyState(point)
		}
		cycle++
	}
	count := 0
	for _, v := range fullState {
		if v == active {
			count++
		}
	}

	return count
}

func deepCopyMap(src map[Point]string, dst map[Point]string) {
	for k, v := range src {
		dst[k] = v
	}
}

func modifyState(point Point) {
	dx := -1
	neigbours := 0
	for dx < 2 {
		dy := -1
		nx := point.x + dx
		for dy < 2 {
			dz := -1
			ny := point.y + dy
			for dz < 2 {
				nz := point.z + dz
				newPoint := Point{nx, ny, nz}
				if dx == 0 && dy == 0 && dz == 0 {
					dz++
					continue
				}
				_, exists := prevFullState[newPoint]
				if !exists {
					prevFullState[newPoint] = inactive
					fullState[newPoint] = inactive
				}

				if prevFullState[newPoint] == active {
					neigbours++
				}
				dz++
			}
			dy++
		}
		dx++
	}

	fmt.Println(point, prevFullState[point], neigbours)
	if prevFullState[point] == inactive && neigbours == 3 {
		fmt.Println("Made active")
		fullState[point] = active
	} else if prevFullState[point] == active && !(neigbours == 3 || neigbours == 2) {
		fmt.Println("Made inactive")
		fullState[point] = inactive
	} else {
		fmt.Println("No Change")
	}
}

func processInitialState(input []string) {
	z := 0
	for y, row := range input {
		for x, col := range row {
			point := Point{x, y, z}
			fullState[point] = string(col)
		}
	}
}

func part2() int {
	return 0
}

// Point ...
type Point struct {
	x int
	y int
	z int
}

var fullState map[Point]string
var prevFullState map[Point]string
var active string = "#"
var inactive string = "."

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}

	prevFullState = map[Point]string{}
	fullState = map[Point]string{}

	processInitialState(input)
	partOne := part1()

	partTwo := part2()

	fmt.Printf("Day 17 - Conway Cube\nPart 1:\t%d\nPart 2:\t%d\n\n", partOne, partTwo)
}
