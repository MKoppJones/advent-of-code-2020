package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type BagCount struct {
	count int
	bag   string
}

func readInput(path string) ([]string, error) {
	bytesRead, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		println(readFileErr)
		return nil, readFileErr
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, ".\n")
	lines = lines[:len(lines)-1]
	return lines, nil
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func crawlCanHold(bagMap map[string][]string, bag string, validBags map[string]int) map[string]int {
	for key, bags := range bagMap {
		if contains(bags, bag) {
			validBags[key] = 0
			validBags = crawlCanHold(bagMap, key, validBags)
		}
	}
	return validBags
}

func crawlHolds(bagMap map[string][]BagCount, bag string, validBags map[BagCount]int) map[BagCount]int {
	for _, bagCount := range bagMap[bag] {
		_, exists := validBags[bagCount]
		if !exists {
			validBags[bagCount] = 0
		}
		validBags[bagCount]++
		validBags = crawlHolds(bagMap, bagCount.bag)
	}
	return validBags
}

func main() {
	args := os.Args[1:]
	input, err := readInput(args[0])
	if err != nil {
		println(err)
		return
	}

	bagMap := map[string][]string{}
	countedBagMap := map[string][]BagCount{}
	for _, bagRule := range input {
		rule := strings.Split(bagRule, " contain ")
		bag := strings.ReplaceAll(rule[0], " bags", "")
		_, exists := bagMap[bag]
		if !exists {
			bagMap[bag] = []string{}
			countedBagMap[bag] = []BagCount{}
		}
		bags := strings.Split(rule[1], ", ")
		for _, storedBag := range bags {
			parts := strings.Split(storedBag, " ")
			color := parts[1] + " " + parts[2]
			bagMap[bag] = append(bagMap[bag], color)
			count, _ := strconv.Atoi(parts[0])
			countedBagMap[bag] = append(countedBagMap[bag], BagCount{count, color})
		}
	}

	canHoldShinyGold := len(crawlCanHold(bagMap, "shiny gold", map[string]int{}))
	maxHold := len(crawlHolds(countedBagMap, "shiny gold", map[BagCount]int{}))

	fmt.Printf("Day 07 - Handy Haversacks\nPart 1:\t%d\nPart 2:\t%d\n\n", canHoldShinyGold, maxHold)
}
