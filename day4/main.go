package main

import (
	aoc "aoc24/library"
	"fmt"
)

type Point struct {
	x, y int
}

var directions = []Point{
	{x: 0, y: 1},
	{x: 1, y: 1},
	{x: 1, y: 0},
	{x: 1, y: -1},
	{x: 0, y: -1},
	{x: -1, y: -1},
	{x: -1, y: 0},
	{x: -1, y: 1},
}

type Diagonals struct {
	dir    Point
	origin Point
}

var diagonals = []Diagonals{
	{
		dir:    Point{x: 1, y: 1},
		origin: Point{x: 0, y: 0},
	}, {
		dir:    Point{x: -1, y: 1},
		origin: Point{x: 2, y: 0},
	}, {
		dir:    Point{x: -1, y: -1},
		origin: Point{x: 2, y: 2},
	}, {
		dir:    Point{x: 1, y: -1},
		origin: Point{x: 0, y: 2},
	},
}

func main() {
	matrix := aoc.ReadFileLines("./input")

	fmt.Println(part1(matrix))
	fmt.Println(part2(matrix))
}

func part1(matrix []string) string {
	xmases := 0

	for y, line := range matrix {
		for x := range line {
			for _, dir := range directions {
				if isStringCorrect(matrix, "XMAS", x, y, dir) {
					xmases++
				}
			}
		}
	}
	return fmt.Sprintf("There are %d XMASs", xmases)
}

func part2(matrix []string) string {
	xmases := 0

	for y, line := range matrix {
		for x := range line {
			num_of_mases := 0
			for _, d := range diagonals {
				if isStringCorrect(matrix, "MAS", x+d.origin.x, y+d.origin.y, d.dir) {
					num_of_mases++
				}
				if isStringCorrect(matrix, "SAM", x+d.origin.x, y+d.origin.y, d.dir) {
					num_of_mases++
				}
			}
            if num_of_mases == 4 {
                xmases++
            }
		}
	}
	return fmt.Sprintf("There are %d XMASs", xmases)
}

func isStringCorrect(matrix []string, target string, x int, y int, dir Point) bool {
	if len(target) == 0 {
		return true
	}

	if x < 0 || y < 0 || y >= len(matrix) || x >= len(matrix[y]) {
		return false
	}

	if matrix[y][x] == target[0] {
		return isStringCorrect(matrix, target[1:], x+dir.x, y+dir.y, dir)
	}

	return false
}
