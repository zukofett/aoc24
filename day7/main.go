package main

import (
	aoc "aoc24/library"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Operation func(x, y int) int

func Add(x, y int) int {
	return x + y
}

func Mul(x, y int) int {
	return x * y
}

func Cat(x, y int) int {
	for rhs := y; rhs != 0; rhs /= 10 {
		x *= 10
	}

	return x + y
}

type Equation struct {
	target   int
	operands []int
}

func main() {
	equations := createEquations("./input")
	fmt.Println(part1(equations))
	fmt.Println(part2(equations))
}

func part1(equations []Equation) string {
	sum := 0
	for _, eq := range equations {
		if isEquationCorrect(eq, []Operation{Mul, Add}) {
			sum += eq.target
		}
	}
	return fmt.Sprintf("The sum is: %d", sum)
}

func part2(equations []Equation) string {
	sum := 0
	for _, eq := range equations {
		if isEquationCorrect(eq, []Operation{Cat, Mul, Add}) {
			sum += eq.target
		}
	}
	return fmt.Sprintf("The sum is: %d", sum)
}

func isEquationCorrect(eq Equation, operators []Operation) bool {
	if len(eq.operands) == 1 {
		return eq.operands[0] == eq.target
	}

	for _, op := range operators {
		curr := op(eq.operands[0], eq.operands[1])
		if curr > eq.target {
			return false
		}

		new_equation := Equation{
			target:   eq.target,
			operands: append([]int{curr}, eq.operands[2:]...),
		}
		if isEquationCorrect(new_equation, operators) {
			return true
		}
	}
	return false
}

func createEquations(path string) []Equation {
	lines := aoc.ReadFileLines(path)
	equations := make([]Equation, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ":")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}

		eq := Equation{
			target:   target,
			operands: aoc.StrsToInts(parts[1]),
		}
		equations = append(equations, eq)
	}
	return equations
}
