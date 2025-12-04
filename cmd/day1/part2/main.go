package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input_01.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create scanner to read file
	scanner := bufio.NewScanner(file)

	dialValue := 50
	zeroCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		var direction string = line[0:1]
		value, _ := strconv.Atoi(line[1:])

		switch direction {
		case "R":
			dialValue += value
		case "L":
			dialValue -= value
		}

		// If negative, increase by 100 means pased 0
		for dialValue < 0 {
			dialValue += 100
			zeroCount++
		}
		// If greater than 100 module means passed 0
		for dialValue > 99 {
			dialValue -= 100
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
}
