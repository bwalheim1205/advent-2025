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
	for i, line := range lines {
		for j, char := range line {
			// If character scan number of them around
			if char == '@' {
				if numberAdjacentOfPaper(lines, i, j) < 4 {
					count++
				}
			}
		}
	}
	fmt.Println(count)
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
