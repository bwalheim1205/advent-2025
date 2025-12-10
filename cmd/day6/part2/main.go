package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func parseFile(r io.Reader) ([]string, string) {
	// Create scanner to read file
	scanner := bufio.NewScanner(r)

	numberLines := []string{}
	var operations string

	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == '*' || line[0] == '+' {
			operations = line
		} else {
			numberLines = append(numberLines, line)
		}

	}
	return numberLines, operations
}

func columnOperation(start int, end int, numbers []string, operation string) int {
	// Set initial count
	var total int
	if operation == "+" {
		total = 0
	} else if operation == "*" {
		total = 1
	}

	for i := start; i <= end; i++ {
		numStr := ""
		for _, number := range numbers {
			if number[i] != ' ' {
				numStr += string(number[i])
			}
		}
		num, err := strconv.Atoi(numStr)
		fmt.Print(numStr, " ")

		if err == nil {
			if operation == "+" {
				total += num
			} else if operation == "*" {
				total *= num
			}
		}
	}
	fmt.Println(operation)
	return total
}

func columnOperationSum(numbers []string, operations string) int {
	sum := 0
	start := 0
	end := 0
	lastOp := ""
	for i, char := range operations {
		if char == '*' || char == '+' {
			start = end
			end = i
			// Not first instance
			if start != end {
				sum += columnOperation(start, end-1, numbers, lastOp)
			}
			lastOp = string(char)
		}
	}
	sum += columnOperation(end, len(operations)-1, numbers, lastOp)

	return sum
}

func main() {
	file, err := os.Open("input_06.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	numbers, operations := parseFile(file)

	sum := columnOperationSum(numbers, operations)

	fmt.Println(sum)
}
