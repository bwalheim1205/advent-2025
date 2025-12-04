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

		joltage := 0

		bank := line
		maxValue := 0
		for i := 12; i > 0; i-- {
			maxValue, bank = getMaxBattery(bank, i)
			joltage = joltage*10 + maxValue
		}

		joltageTotal += joltage
	}

	fmt.Println(joltageTotal)
}

func getMaxBattery(bank string, batteriesLeft int) (int, string) {

	max := 0
	maxIndex := 0
	for i := 0; i < len(bank)-batteriesLeft+1; i++ {
		value, _ := strconv.Atoi(string(bank[i]))
		if value > max {
			max = value
			maxIndex = i
		}
	}
	return max, bank[maxIndex+1:]
}
