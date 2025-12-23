package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func parseFile(r io.Reader) map[string][]string {
	// Create scanner to read file
	scanner := bufio.NewScanner(r)

	var graph map[string][]string = make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()

		splitStirings := strings.Split(line, ":")

		node := splitStirings[0]
		nextNodeList := splitStirings[1]
		nextNodes := strings.Fields(nextNodeList)

		graph[node] = nextNodes
	}

	return graph
}

func findNumberOfPaths(graph map[string][]string) int {

	start := "you"
	end := "out"

	numberOfPaths := 0

	queue := []string{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == end {
			numberOfPaths++
		} else {
			queue = append(queue, graph[node]...)
		}
	}

	return numberOfPaths
}

func main() {
	file, err := os.Open("input_11.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	graph := parseFile(file)
	paths := findNumberOfPaths(graph)

	fmt.Println("Number of paths:", paths)
}
