package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input_03.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create scanner to read file
	scanner := bufio.NewScanner(file)

	joltageTotal := 0

	for scanner.Scan() {
		line := scanner.Text()

		joltage := getJoltage(line)
		joltageTotal += joltage
	}

	fmt.Println(joltageTotal)
}

func getJoltage(bank string) int {
	// Find max first digit
	firstMax := 0
	firstMaxIndex := 0
	for i := 0; i < len(bank)-1; i++ {
		value, _ := strconv.Atoi(string(bank[i]))
		if value > firstMax {
			firstMax = value
			firstMaxIndex = i
		}
	}

	secondMax := 0
	for i := firstMaxIndex + 1; i < len(bank); i++ {
		value, _ := strconv.Atoi(string(bank[i]))
		if value > secondMax {
			secondMax = value
		}
	}

	return firstMax*10 + secondMax
}
