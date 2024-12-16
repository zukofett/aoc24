package main

import (
	"fmt"

	aoc "aoc24/library"
)

func main() {
	safeScans := 0
	fixedScans := 0

	lines := aoc.ReadFileLines("./input")

	for _, line := range lines {
		levels := aoc.StrsToInts(line)

		if isScanSafe(levels) {
			safeScans++
		} else if afterProblemDumper(levels) {
			fixedScans++
		}
	}

	fmt.Printf("The number of safe scans is %d\n", safeScans)
	fmt.Printf("The number of safe scans after fix is %d\n", safeScans+fixedScans)
}

func isScanSafe(levels []int) bool {
	dir := getDirection(levels[0], levels[len(levels)-1])

	for i := 1; i < len(levels); i++ {
		diff := (levels[i-1] - levels[i]) * dir
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func afterProblemDumper(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		if isScanSafe(removeProblem(levels, i)) {
			return true
		}
	}
	return false
}

func removeProblem(levels []int, ignore int) []int {
	fixed := make([]int, 0, len(levels)-1)

	for i, lvl := range levels {
		if i == ignore {
			continue
		}
		fixed = append(fixed, lvl)
	}
	return fixed
}

func getDirection(start, end int) int {
	if start > end {
		return 1
	}
	if start < end {
		return -1
	}
	return 0
}
