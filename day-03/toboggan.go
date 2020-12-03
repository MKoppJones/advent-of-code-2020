package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Position is the cursor position
type Position struct {
	x int
	y int
}

// Slope describes the slope
type Slope struct {
	right int
	down  int
}

func readMap(path string) (string, error) {
	treeMap := ""

	bytesRead, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		println(readFileErr)
		return "", readFileErr
	}
	treeMap = string(bytesRead)
	return treeMap, nil
}

func checkTrees(mapString string, slope Slope) int {
	flatMap := strings.ReplaceAll(mapString, "\n", "")
	lines := strings.Split(mapString, "\n")
	width := len(lines[0])
	height := len(lines)
	cursor := Position{0, 0}

	treeCount := 0

	for cursor.y+slope.down < height {
		cursor.x += slope.right
		cursor.y += slope.down
		x := cursor.x % width
		y := cursor.y % height
		index := x + (y * width)
		if index > len(flatMap) {
			break
		}
		if string(flatMap[index]) == "#" {
			treeCount++
		}
	}
	return treeCount
}

func checkSlopes(treeMap string, slopes []Slope) int {
	value := 1
	for _, slope := range slopes {
		treeCount := checkTrees(treeMap, slope)
		value = value * treeCount
		println(treeCount)
	}
	return value
}

func main() {
	args := os.Args[1:]
	treeMapString, err := readMap(args[0])
	if err != nil {
		println(err)
		return
	}
	slopes := []Slope{Slope{1, 1}, Slope{3, 1}, Slope{5, 1}, Slope{7, 1}, Slope{1, 2}}

	fmt.Printf("Day 03 - Toboggan Trajectory\nPart 1:\t%d\nPart 2:\t%d\n\n", checkTrees(treeMapString, slopes[1]), checkSlopes(treeMapString, slopes))
}
