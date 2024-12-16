package main

import (
	aoc "aoc24/library"
	"fmt"
)

type Chalange struct {
	lines   [][]rune
	guard   aoc.Coordinate
	visited map[aoc.Point]struct{}
}

func (ch *Chalange) isInBounds(p aoc.Point) bool {
	return aoc.IsInRange(p.Y, 0, len(ch.lines)) &&
		aoc.IsInRange(p.X, 0, len(ch.lines[p.Y]))
}

func (ch *Chalange) nextIsBlocked(g aoc.Coordinate) bool {
	next := g.GetNext()

	if !ch.isInBounds(next) {
		return false
	}

	return ch.lines[next.Y][next.X] == '#'
}

func findGuard(lines [][]rune) aoc.Coordinate {
	guard := aoc.Coordinate{Direction: aoc.Up}
	for i, line := range lines {
		for j, ch := range line {
			if ch == '^' {
				guard.Position.X = j
				guard.Position.Y = i
				return guard
			}
		}
	}
	return guard
}

func main() {
	ch := Chalange{
		lines: aoc.ReadFileLinesAsRunes("./input"),
	}

	ch.guard = findGuard(ch.lines)
	ch.visited = make(map[aoc.Point]struct{})

	fmt.Println(ch.part1())
	fmt.Println(ch.part2())
}

func (ch *Chalange) part1() string {
	tiles := 0

	for {
		ch.guard.Walk()

		if !ch.isInBounds(ch.guard.Position) {
			break
		}

		if _, exists := ch.visited[ch.guard.Position]; !exists {
			ch.visited[ch.guard.Position] = struct{}{}
			tiles++
		}

		for ch.nextIsBlocked(ch.guard) {
			ch.guard.TurnRight()
		}
	}

	return fmt.Sprintf("num of visited tiles is: %d\n", tiles)
}

func (ch *Chalange) nextIsLoop(start aoc.Coordinate) bool {
	coords := make(map[aoc.Coordinate]struct{})

	for {
		if _, exists := coords[start]; !exists {
			coords[start] = struct{}{}
		} else {
			return true
		}

		if !ch.isInBounds(start.Position) {
			break
		}

		start.Walk()

		for ch.nextIsBlocked(start) {
			start.TurnRight()
		}
	}

	return false
}

func (ch *Chalange) part2() string {
	start := findGuard(ch.lines)

	points := make(map[aoc.Point]struct{})

	for step := range ch.visited {
		if ch.lines[step.Y][step.X] == '.' {
			ch.lines[step.Y][step.X] = '#'
			if ch.nextIsLoop(start) {
				if _, ok := points[step]; !ok {
					points[step] = struct{}{}
				}
			}
			ch.lines[step.Y][step.X] = '.'
		}

	}
	return fmt.Sprintf("num of visited tiles is: %d\n", len(points))
}
