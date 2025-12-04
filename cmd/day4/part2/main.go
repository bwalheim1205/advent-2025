package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input_04.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create scanner to read file
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	count := 0
	removed, countRemoved := removeAvailablePaper(lines)
	for countRemoved > 0 {
		count += countRemoved
		lines = removed
		removed, countRemoved = removeAvailablePaper(lines)
	}

	fmt.Println(count)
}

func removeAvailablePaper(lines []string) ([]string, int) {
	countRemoved := 0
	newLines := make([]string, len(lines))
	for i, line := range lines {
		for j, char := range line {
			// If character scan number of them around
			if char == '@' && numberAdjacentOfPaper(lines, i, j) < 4 {
				countRemoved++
				newLines[i] += "x"

			} else {
				newLines[i] += string(char)
			}
		}
	}

	return newLines, countRemoved
}

func numberAdjacentOfPaper(lines []string, x int, y int) int {

	count := 0

	// Iterate through neighbors
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			// Check in bounds
			if i >= 0 && i < len(lines) && j >= 0 && j < len(lines[i]) {
				// Not itself and is paper
				if !(i == x && j == y) && string(lines[i][j]) == "@" {
					count++
				}
			}
		}
	}

	return count
}
