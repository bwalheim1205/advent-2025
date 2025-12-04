package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input_02.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create scanner to read file
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		range_strings := strings.Split(line, ",")
		for _, range_string := range range_strings {
			numbers := strings.Split(range_string, "-")
			start, _ := strconv.Atoi(numbers[0])
			end, _ := strconv.Atoi(numbers[1])

			for i := start; i <= end; i++ {
				if isInvalidId(i) {
					count += i
				}
			}
		}
	}
	fmt.Println(count)
}

func isInvalidId(id int) bool {
	len_id := len(strconv.Itoa(id))

	for i := len_id / 2; i > 0; i -= 1 {
		if len_id%i == 0 {
			if isInvalidIdSize(id, i) {
				return true
			}
		}
	}
	return false
}

func isInvalidIdSize(id int, len int) bool {

	mask := int(math.Pow10(len))

	value := id
	match := id % mask
	for value > 0 {
		if value%mask != match {
			return false
		}
		value /= mask
	}
	return true
}
