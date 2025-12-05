package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func main() {
	file, err := os.Open("input_05.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create scanner to read file
	scanner := bufio.NewScanner(file)

	ranges := []Range{}
	rangeMode := true
	numberIngredients := 0
	numberFresh := 0

	for scanner.Scan() {
		line := scanner.Text()
		// No longer scan ranges
		if line == "" {
			rangeMode = false
			fmt.Println(len(ranges))
		} else if rangeMode {
			// Parse String
			splitString := strings.Split(line, "-")
			start, _ := strconv.Atoi(splitString[0])
			end, _ := strconv.Atoi(splitString[1])

			// Append new range to list
			freshRange := Range{
				Start: start,
				End:   end,
			}
			ranges = append(ranges, freshRange)
		} else {
			number, _ := strconv.Atoi(line)
			numberIngredients++
			for _, freshRange := range ranges {
				if checkFresh(number, freshRange.Start, freshRange.End) {
					numberFresh++
					break
				}
			}
		}
	}

	fmt.Println("Ingredients", numberIngredients)
	fmt.Println("Fresh", numberFresh)
}

func checkFresh(day int, start int, end int) bool {
	if day >= start && day <= end {
		return true
	}
	return false
}
