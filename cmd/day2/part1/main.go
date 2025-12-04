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

	if len_id%2 != 0 {
		return false
	} else {
		mask := int(math.Pow10(len_id / 2))
		part1 := id / mask
		part2 := id % mask
		return part1 == part2
	}

}
