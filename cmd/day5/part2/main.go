package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// Map of valid fresh numbers
	freshRanges := []Range{}

	for scanner.Scan() {
		line := scanner.Text()
		// No longer scan ranges
		if line == "" {
			break
		} else {
			// Parse String
			splitString := strings.Split(line, "-")
			start, _ := strconv.Atoi(splitString[0])
			end, _ := strconv.Atoi(splitString[1])

			// Append new range to list
			freshRanges = append(freshRanges, Range{
				Start: start,
				End:   end,
			})

		}
	}

	// Removes overlapping ranges from the slice
	noOverlap := removeOverlap(freshRanges)

	fmt.Println(numbersInRanges(noOverlap))

}

func removeOverlap(freshRanges []Range) []Range {
	// Sort the existing slice of ranges by start
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i].Start < freshRanges[j].Start
	})

	// Create new range to return
	newRanges := []Range{}

	for i := 0; i < len(freshRanges); i++ {
		// Get recent range
		freshRange := freshRanges[i]

		// If range doesn't overlap or is last enetry add to new ranges
		if i+1 >= len(freshRanges) || freshRange.End < freshRanges[i+1].Start {
			newRanges = append(newRanges, freshRange)
		} else {
			// If overlap iterate through overlapping ranges
			for i < len(freshRanges)-1 && freshRange.End >= freshRanges[i+1].Start {
				// Update end to be largest of two ranges
				if freshRange.End < freshRanges[i+1].End {
					freshRange.End = freshRanges[i+1].End
				}

				i++
			}
			newRanges = append(newRanges, freshRange)
		}
	}

	return newRanges
}

func numbersInRanges(ranges []Range) int {
	count := 0
	for _, value := range ranges {
		count += value.End - value.Start + 1
	}
	return count
}
