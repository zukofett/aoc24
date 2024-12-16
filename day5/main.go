package main

import (
	aoc "aoc24/library"

	"fmt"
	"slices"
)

type Chalange struct {
	WrongPagesMap map[int][]int
	UpdateLists   [][]int
}

func main() {
	parts := aoc.ReadFileLinesInParts("./input")

	wrong_pages := make(map[int][]int, 0)
	for _, l := range parts[0] {
		nums := aoc.StrsToIntsTok(l, "|")
		wrong_pages[nums[1]] = append(wrong_pages[nums[1]], nums[0])
	}

	update_lists := make([][]int, 0)
	for _, l := range parts[1] {
		nums := aoc.StrsToIntsTok(l, ",")
		update_lists = append(update_lists, nums)
	}

	chalange := Chalange{
		WrongPagesMap: wrong_pages,
		UpdateLists:   update_lists,
	}

	fmt.Println(chalange.part1())
	fmt.Println(chalange.part2())
}

func (ch *Chalange) part1() string {
	sum := 0
	for _, l := range ch.UpdateLists {
		if ch.isListValid(l) {
			sum += l[len(l)/2]
		}
	}

	return fmt.Sprintf("The sum is: %d", sum)
}

func (ch *Chalange) part2() string {
	sum := 0
	for _, l := range ch.UpdateLists {
		if !ch.isListValid(l) {
			ch.makeValid(l)
			sum += l[len(l)/2]
		}
	}

	return fmt.Sprintf("The sum is: %d", sum)
}

func (ch *Chalange) isListValid(list []int) bool {
	for i, num := range list {
        if aoc.SlicesIntersect(list[i:], ch.WrongPagesMap[num]) {
            return false
        }
	}
	return true
}

func (ch *Chalange) makeValid(list []int) {
	for i := 0; i < len(list); {
        changed := false
		for _, n := range ch.WrongPagesMap[list[i]] {
			if slices.Contains(list[i:], n) {
				wrongIdx := slices.Index(list, n)
				list[i], list[wrongIdx] = list[wrongIdx], list[i]
                changed = true
                break
			}
		}
        if !changed {
            i++
        }
	}
}
