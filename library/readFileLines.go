package aoc

import (
	"bufio"
	"log"
	"os"
)

func ReadFileLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ReadFileLinesAsRunes(path string) [][]rune {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]rune, 0)

	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ReadFileLinesAsBytes(path string) [][]byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]byte, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Bytes())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ReadFileLinesAsInts(path string) [][]int {
	lines := ReadFileLines(path)

	nums := make([][]int, len(lines))

	for i, line := range lines {
		for _, digit := range line {
			nums[i] = append(nums[i], int(digit-'0'))
		}
	}

	return nums
}

func ReadFileLinesInParts(path string) [][]string {
	lines := ReadFileLines(path)

	parts := make([][]string, 0)

	start := 0

	for i, l := range lines {
		if l == "" {
			parts = append(parts, lines[start:i])
			start = i + 1
		}
	}

	if start < len(lines) {
		parts = append(parts, lines[start:])
	}

	return parts
}
