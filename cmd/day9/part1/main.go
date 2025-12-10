package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func parseFile(r io.Reader) []Point {
	// Create scanner to read file
	scanner := bufio.NewScanner(r)

	points := []Point{}

	for scanner.Scan() {
		line := scanner.Text()

		splitStirings := strings.Split(line, ",")

		x, _ := strconv.Atoi(splitStirings[0])
		y, _ := strconv.Atoi(splitStirings[1])
		points = append(points, Point{
			x: x,
			y: y,
		})

	}
	return points
}

func maxArea(points []Point) int {
	max := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			width := (points[i].x - points[j].x)
			height := (points[i].y - points[j].y)
			if width < 0 {
				width *= -1
			}
			if height < 0 {
				height *= -1
			}
			area := (width + 1) * (height + 1)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func main() {
	file, err := os.Open("input_09.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	points := parseFile(file)
	fmt.Println(points[0].x, points[0].y)
	area := maxArea(points)
	fmt.Println(area)
}
