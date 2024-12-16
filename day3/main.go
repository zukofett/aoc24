package main

import (
	aoc "aoc24/library"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var MultRX = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var NewMultRX = regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\))(?:[^d]|d[^o]|[do[^n]|[do[^(]|[do\([^)]|don[^']|don'[^t]|[don't[^(]|don't\([^)])*?(do\(\)|don't\(\))`)

func main() {
	lines := aoc.ReadFileLines("./input")

	fmt.Println(first(lines))
	fmt.Println(second(lines))

}

func first(lines []string) string {
	target := 0
	for _, l := range lines {
		lineOps := MultRX.FindAllStringSubmatch(l, -1)
		for _, x := range lineOps {
			first, err := strconv.Atoi(x[1])
			if err != nil {
				log.Fatal(err)
			}
			second, err := strconv.Atoi(x[2])
			if err != nil {
				log.Fatal(err)
			}
			target += first * second
		}
	}
	return fmt.Sprintf("target is: %d", target)
}

func second(lines []string) string {
	target := 0
	mult := true
	line := strings.Join(lines, "")
	lineOps := NewMultRX.FindAllStringSubmatch(line, -1)
	for _, x := range lineOps {
		if mult {
			expression := MultRX.FindAllStringSubmatch(x[0], -1)
			for _, e := range expression {
				first, err := strconv.Atoi(e[1])
				if err != nil {
					log.Fatal(err)
				}
				second, err := strconv.Atoi(e[2])
				if err != nil {
					log.Fatal(err)
				}
				target += first * second
			}
		}
		mult = x[3] == "do()"
	}

	return fmt.Sprintf("target is: %d", target)
}
