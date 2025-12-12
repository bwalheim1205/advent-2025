package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	I int
	X int
	Y int
	Z int
}

type Wire struct {
	Point1   Point
	Point2   Point
	Distance float64
}

func parseFile(r io.Reader) []Point {
	// Create scanner to read file
	scanner := bufio.NewScanner(r)

	points := []Point{}
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		splitStirings := strings.Split(line, ",")

		x, _ := strconv.Atoi(splitStirings[0])
		y, _ := strconv.Atoi(splitStirings[1])
		z, _ := strconv.Atoi(splitStirings[2])
		points = append(points, Point{
			I: count,
			X: x,
			Y: y,
			Z: z,
		})

		count++
	}
	return points
}

func calculateDistance(point1 Point, point2 Point) float64 {

	// short disatnce formula
	x := point1.X - point2.X
	y := point1.Y - point2.Y
	z := point1.Z - point2.Z

	return math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2) + math.Pow(float64(z), 2))
}

func generateWires(points []Point) []Wire {

	wires := []Wire{}

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			point1 := points[i]
			points2 := points[j]
			wire := Wire{
				Point1:   point1,
				Point2:   points2,
				Distance: calculateDistance(point1, points2),
			}
			wires = append(wires, wire)
		}
	}

	return wires
}

func connectJunctionBoxes(wires []Wire) map[int]int {
	sort.Slice(wires, func(i, j int) bool {
		return wires[i].Distance < wires[j].Distance
	})

	// generate graph of first 1000 wires
	var adjacencyGraph map[int][]int = make(map[int][]int)
	for i := 0; i < 1000; i++ {
		if _, exist := adjacencyGraph[wires[i].Point1.I]; !exist {
			adjacencyGraph[wires[i].Point1.I] = []int{}
		}
		if _, exist := adjacencyGraph[wires[i].Point2.I]; !exist {
			adjacencyGraph[wires[i].Point2.I] = []int{}
		}
		adjacencyGraph[wires[i].Point1.I] = append(adjacencyGraph[wires[i].Point1.I], wires[i].Point2.I)
		adjacencyGraph[wires[i].Point2.I] = append(adjacencyGraph[wires[i].Point2.I], wires[i].Point1.I)
	}

	// size of graphs
	var graphSize map[int]int = make(map[int]int)
	var visited map[int]bool = make(map[int]bool)
	for start := range adjacencyGraph {
		if visited[start] {
			continue
		}

		queue := []int{start}
		visited[start] = true
		size := 1

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, neigh := range adjacencyGraph[node] {
				if !visited[neigh] {
					visited[neigh] = true
					queue = append(queue, neigh)
					size++
				}
			}
		}

		graphSize[start] = size
	}

	return graphSize
}

func sortMapByValue(m map[int]int) []int {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys
}

func main() {
	file, err := os.Open("input_08.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	points := parseFile(file)
	wires := generateWires(points)
	graphSize := connectJunctionBoxes(wires)
	graphKeysSorted := sortMapByValue(graphSize)
	total := 1
	for i := 0; i < 3; i++ {
		fmt.Println("Circuit", i+1, "-", graphSize[graphKeysSorted[i]])
		total *= graphSize[graphKeysSorted[i]]
	}
	fmt.Println(total)

}
