package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input_07.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var beams map[int]bool = make(map[int]bool)
	count := 0

	// Create scanner to read file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// If no beams found find starting beam
		printLine := ""
		if len(beams) == 0 {
			index := strings.Index(line, "S")
			beams[index] = true
			printLine = line
		} else {
			for j, char := range line {
				if _, exists := beams[j]; exists {
					if char == '^' {
						count++
						beams[j-1] = true
						beams[j+1] = true
						delete(beams, j)
						printLine = printLine[:len(printLine)-1] + "|^"
					} else {
						printLine += "|"
					}
				} else {
					printLine += string(char)
				}
			}
		}
		fmt.Println(printLine)
	}
	fmt.Println("Beam Count: ", count)
}
