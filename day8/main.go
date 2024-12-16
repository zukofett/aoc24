package main

import (
	aoc "aoc24/library"
	"fmt"
)

type Chalange struct {
	lines         []string
	signalTypes   map[rune][]aoc.Point
	antinodes     map[aoc.Point]struct{}
	isPointOnGrid func(p aoc.Point) bool
}

func (ch *Chalange) addAntenaIfShould(should bool, location aoc.Point) {
	if should {
		ch.antinodes[location] = struct{}{}
	}
}

func findBeforeAndAfter(first, second aoc.Point) (aoc.Point, aoc.Point) {
	distance := first.FindDistance(second)
	after := distance.JoinPoints(second)

	revDistance := distance.ReversePoint()
	before := revDistance.JoinPoints(first)

	return before, after
}

func (ch *Chalange) findInLine(first, second aoc.Point) []aoc.Point {
	locations := make([]aoc.Point, 0)
	distance := first.FindDistance(second)

	for runner := first; ch.isPointOnGrid(runner); runner.AddPoint(distance) {
		locations = append(locations, runner)
	}

	revDistance := distance.ReversePoint()
	for runner := second; ch.isPointOnGrid(runner); runner.AddPoint(revDistance) {
		locations = append(locations, runner)
	}

	return locations
}

func (ch *Chalange) part1() string {
	for _, antenas := range ch.signalTypes {
		for i, antena := range antenas {
			for j := i + 1; j < len(antenas); j++ {
				before, after := findBeforeAndAfter(antena, antenas[j])
				ch.addAntenaIfShould(ch.isPointOnGrid(before), before)
				ch.addAntenaIfShould(ch.isPointOnGrid(after), after)
			}
		}
	}
	return fmt.Sprintf("the number of antinodes in part 1 is: %d", len(ch.antinodes))
}

func (ch *Chalange) part2() string {
	for _, antenas := range ch.signalTypes {
		for i, antena := range antenas {
			for j := i + 1; j < len(antenas); j++ {
				locations := ch.findInLine(antena, antenas[j])
				for _, location := range locations {
					ch.addAntenaIfShould(ch.isPointOnGrid(location), location)
				}
			}
		}
	}
	return fmt.Sprintf("the number of antinodes in part 2 is: %d", len(ch.antinodes))
}

func main() {
	ch := createChalange("./input")

	fmt.Println(ch.part1())
	fmt.Println(ch.part2())
}

func createChalange(path string) *Chalange {
	lines := aoc.ReadFileLines(path)
	signalTypes := make(map[rune][]aoc.Point)

	for y, line := range lines {
		for x, r := range line {
			if aoc.IsAlpha(r) {
				signalTypes[r] = append(signalTypes[r], aoc.Point{X: x, Y: y})
			}
		}
	}

	return &Chalange{
		lines:         lines,
		signalTypes:   signalTypes,
		antinodes:     make(map[aoc.Point]struct{}),
		isPointOnGrid: aoc.OnGrid(len(lines), len(lines[0])),
	}
}
