package main

import (
	aoc "aoc24/library"
	"fmt"
)

type Chalange struct {
	grid     [][]int
	isOnGrid func(p aoc.Point) bool
}

func main() {
	lines := aoc.ReadFileLinesAsInts("./input")
	ch := Chalange{
		grid: lines,
		isOnGrid: func(p aoc.Point) bool {
			return p.IsOnGrid(len(lines[0]), len(lines))
		},
	}

	fmt.Println(ch.part1())
	fmt.Println(ch.part2())
}

func (ch *Chalange) backtrackPath(points map[aoc.Point]int, pos aoc.Point, height int) {
	if !ch.isOnGrid(pos) || height != ch.grid[pos.Y][pos.X] {
		return
	}

	if height == 9 {
		points[pos] += 1
		return
	}

	for _, dir := range aoc.CardinalDirections {
		ch.backtrackPath(points, pos.JoinPoints(dir), height+1)
	}
}

func (ch *Chalange) part1() string {
	score := 0
	for i, line := range ch.grid {
		for j := range line {
			points := make(map[aoc.Point]int)
			ch.backtrackPath(points, aoc.Point{X: j, Y: i}, 0)
			score += len(points)
		}
	}

	return fmt.Sprintf("the score is: %d;", score)
}

func (ch *Chalange) part2() string {
	rating := 0
	for i, line := range ch.grid {
		for j := range line {
			points := make(map[aoc.Point]int)
			ch.backtrackPath(points, aoc.Point{X: j, Y: i}, 0)
			for _, v := range points {
				rating += v
			}
		}
	}

	return fmt.Sprintf("the rating is: %d;", rating)
}
