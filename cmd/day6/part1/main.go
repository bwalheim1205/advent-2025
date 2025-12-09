package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parseFile(r io.Reader) ([][]int, []string) {
	// Create scanner to read file
	scanner := bufio.NewScanner(r)

	numbers := [][]int{}
	var operations []string

	for scanner.Scan() {
		line := scanner.Text()

		// Split by white space
		fields := strings.Fields(line)

		// If not an interger, it is an operation
		_, err := strconv.Atoi(fields[0])
		if err == nil {
			line_numbers := []int{}
			for _, field := range fields {
				number, err := strconv.Atoi(field)
				if err != nil {

				}
				line_numbers = append(line_numbers, number)
			}
			numbers = append(numbers, line_numbers)
		} else {
			operations = fields
		}

	}
	return numbers, operations
}

func columnOperationSum(numbers [][]int, operations []string) int {
	count := 0
	for i, op := range operations {
		if op == "+" {
			for _, rows := range numbers {
				count += rows[i]
			}
		} else if op == "*" {
			multTotal := 1
			for _, rows := range numbers {
				multTotal *= rows[i]
			}
			count += multTotal
		}
	}
	return count
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
