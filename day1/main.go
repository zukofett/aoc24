package main

import (
	"fmt"
	"sort"

	aoc "aoc24/library"
)

func main() {
	lines := aoc.ReadFileLines("./input")
    size := len(lines)

	leftArr := make([]int, 0, size)
	rightArr := make([]int, 0, size)
	counts := make(map[int]int, size)

	for _, line := range lines {
		parts := aoc.StrsToInts(line)

		leftArr = append(leftArr, parts[0])
		rightArr = append(rightArr, parts[1])

		counts[parts[1]]++
	}

	sort.Ints(leftArr)
	sort.Ints(rightArr)

	distance := 0
	simScore := 0

	for i := 0; i < len(lines); i++ {
		distance += aoc.AbsDiff(leftArr[i], rightArr[i])
		simScore += leftArr[i] * counts[leftArr[i]]
	}

	fmt.Printf("The distance is %d\n", distance)
	fmt.Printf("The sim score is %d\n", simScore)
}
