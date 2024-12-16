package aoc

import (
	"log"
	"strconv"
	"strings"
)

func StrsToInts(line string) []int {
	strs := strings.Fields(line)

	ints := make([]int, len(strs))

	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		ints[i] = num
	}

	return ints
}

func StrsToIntsTok(line string, sep string) []int {
	strs := strings.Split(line, sep)

	ints := make([]int, len(strs))

	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		ints[i] = num
	}

	return ints
}

func IsAlpha(r rune) bool {
	return r <= 'Z' && r >= 'A' || r <= 'z' && r >= 'a' || r <= '9' && r >= '0'
}
