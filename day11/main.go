package main

import (
	aoc "aoc24/library"
	"fmt"
)

type Key struct {
	value  int
	blinks int
}

func main() {
	input := aoc.ReadFileLines("./input")[0]
	rocks := aoc.StrsToInts(input)

	fmt.Println(part1(rocks))
	fmt.Println(part2(rocks))
}

func part1(rocks []int) string {
	return fmt.Sprintf("the num of rocks is: %d;", run(rocks, 25))
}

func part2(rocks []int) string {
	return fmt.Sprintf("the num of rocks is: %d;", run(rocks, 75))
}

func run(rocks []int, blinks int) int {
	cache := make(map[Key]int)

	res := 0
	for _, rock := range rocks {
		res += loop(rock, blinks, cache)
	}
	return res
}

func loop(rock, blinks int, cache map[Key]int) int {
	if blinks == 0 {
		return 1
	}

	if r, exists := cache[Key{value: rock, blinks: blinks}]; exists {
		return r
	}

	if rock == 0 {
		res := loop(1, blinks-1, cache)
		cache[Key{value: rock, blinks: blinks}] = res
		return res
	}

	if digits := aoc.CountDigits(rock); digits%2 == 0 {
		a, b := aoc.SplitDigits(rock, digits)
		res := loop(a, blinks-1, cache) + loop(b, blinks-1, cache)
		cache[Key{value: rock, blinks: blinks}] = res

		return res
	}

	res := loop(rock*2024, blinks-1, cache)
	cache[Key{value: rock, blinks: blinks}] = res
	return res
}
