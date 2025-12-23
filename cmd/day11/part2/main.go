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

	numberOfPathsSToFFT := dfs(graph, "svr", map[string]bool{}, "fft")

	return numberOfPathsSToFFT
}

func dfs(graph map[string][]string, node string, visited map[string]bool, end string) int {

	if node == end {
		if visited["fft"] && visited["dac"] {
			return 1
		}
		return 0
	}

	visited[node] = true
	paths := 0

	for _, next := range graph[node] {
		if !visited[next] {
			paths += dfs(graph, next, visited, end)
		}
	}

	delete(visited, node)
	return paths
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
