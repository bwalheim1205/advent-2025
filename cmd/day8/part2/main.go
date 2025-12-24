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

func connectJunctionBoxes(points []Point, wires []Wire) (Point, Point) {

	sort.Slice(wires, func(i, j int) bool {
		return wires[i].Distance < wires[j].Distance
	})

	var groups map[int]int = make(map[int]int)
	groupCount := 0

	// Iterate through deleting visited nodes
	for _, wire := range wires {

		// Check if nodes are already in group
		group1, ok1 := groups[wire.Point1.I]
		group2, ok2 := groups[wire.Point2.I]

		// Create a new group
		if !ok1 && !ok2 {
			groups[wire.Point1.I] = groupCount
			groups[wire.Point2.I] = groupCount
			groupCount++
		} else {

			// Else merge groups
			if !ok1 {
				group1 = groupCount
			}
			if !ok2 {
				group2 = groupCount
			}

			groups[wire.Point1.I] = group1
			groups[wire.Point2.I] = group2

			// Merge groups
			if group1 != group2 {
				if group1 > group2 {
					for k, v := range groups {
						if v == group1 {
							groups[k] = group2
						}
					}
				} else {
					for k, v := range groups {
						if v == group2 {
							groups[k] = group1
						}
					}
				}
			}
		}

		groupCounts := make(map[int]int)
		for _, v := range groups {
			groupCounts[v]++
		}
		fmt.Println(groupCounts)

		if len(points) == len(groups) {
			sameGroup := true
			for _, v := range groups {
				if v != 0 {
					sameGroup = false
					break
				}
			}
			if sameGroup {
				return wire.Point1, wire.Point2
			}
		}
	}
	return Point{}, Point{}
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
	point1, point2 := connectJunctionBoxes(points, wires)
	fmt.Println("Point 1:", point1.I, "(", point1.X, "-", point1.Y, "-", point1.Z, ")")
	fmt.Println("Point 2:", point2.I, "(", point2.X, "-", point2.Y, "-", point2.Z, ")")
	fmt.Println("key:", point1.X*point2.X)

}
