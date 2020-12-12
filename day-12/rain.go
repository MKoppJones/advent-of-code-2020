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

func moveDir(facing string, count int, x int, y int) (int, int) {
	switch facing {
	case "N":
		y += count
	case "S":
		y -= count
	case "E":
		x += count
	case "W":
		x -= count
	}
	return x, y
}

func rotate(starting string, steps int) string {
	for steps > 0 {
		switch starting {
		case "N":
			starting = "W"
		case "W":
			starting = "S"
		case "S":
			starting = "E"
		case "E":
			starting = "N"
		}
		steps--
	}
	return starting
}

func rotateWaypoint(steps int, x int, y int) (int, int) {
	for steps > 0 {
		x, y = y, -x
		steps--
	}
	return x, y
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func moveDirByWapoint(count int, x int, y int, waypointX int, waypointY int) (int, int) {
	return x + (waypointX * count), y + (waypointY * count)
}

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}

	facing := "E"
	x, y := 0, 0
	for _, rule := range input {
		dir := string(rule[0])
		count, _ := strconv.Atoi(rule[1:])
		switch dir {
		case "F":
			x, y = moveDir(facing, count, x, y)
		case "N":
			x, y = moveDir("N", count, x, y)
		case "S":
			x, y = moveDir("S", count, x, y)
		case "E":
			x, y = moveDir("E", count, x, y)
		case "W":
			x, y = moveDir("W", count, x, y)
		case "L":
			steps := count / 90
			facing = rotate(facing, steps)
		case "R":
			steps := count / 90
			facing = rotate(facing, steps*3)
		}
	}

	manhattan := abs(x) + abs(y)

	x, y = 0, 0
	waypointX, waypointY := 10, 1
	for _, rule := range input {
		dir := string(rule[0])
		count, _ := strconv.Atoi(rule[1:])
		switch dir {
		case "F":
			x, y = moveDirByWapoint(count, x, y, waypointX, waypointY)
		case "N":
			waypointX, waypointY = moveDir("N", count, waypointX, waypointY)
		case "S":
			waypointX, waypointY = moveDir("S", count, waypointX, waypointY)
		case "E":
			waypointX, waypointY = moveDir("E", count, waypointX, waypointY)
		case "W":
			waypointX, waypointY = moveDir("W", count, waypointX, waypointY)
		case "L":
			steps := count / 90
			waypointX, waypointY = rotateWaypoint(steps*3, waypointX, waypointY)
		case "R":
			steps := count / 90
			waypointX, waypointY = rotateWaypoint(steps, waypointX, waypointY)
		}
		// println(dir, count, x, y, waypointX, waypointY)
	}
	manhattan2 := abs(x) + abs(y)

	fmt.Printf("Day 12 - Rain Risk\nPart 1:\t%d\nPart 2:\t%d\n\n", manhattan, manhattan2)
}
