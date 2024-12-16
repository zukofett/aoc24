package main

import (
	aoc "aoc24/library"
	"fmt"
)

type Disk struct {
	disk        []int
	blocks      [][]int
	emptyBlocks [][]int
}

func main() {
	disk := aoc.ReadFileLinesAsRunes("./input")[0]

	fmt.Println(part1(disk))
	fmt.Println(part2(disk))

}

func part1(line []rune) string {
	disk := createDisk(line)
	defraged := defragment(disk)
	checksum := checkSum(defraged)

	return fmt.Sprintf("the checksum for part1 id: %d", checksum)
}

func part2(line []rune) string {
	disk := createDiskBlocks(line)
	defraged := defragmentByBlock(disk)
	checksum := checkSum(defraged.disk)

	return fmt.Sprintf("the checksum for part2 id: %d", checksum)
}

func createDisk(disk []rune) []int {
	blocks := make([]int, 0)

	for i := 0; i < len(disk); i++ {
		id := i / 2
		if i&1 != 0 {
			id = -1
		}

		ammount := disk[i] - '0'

		for ammount > 0 {
			blocks = append(blocks, id)
			ammount--
		}
	}
	return blocks
}

func createDiskBlocks(disk []rune) Disk {
	formattedDisk := Disk{
		disk:        createDisk(disk),
		blocks:      make([][]int, 0, len(disk)/2),
		emptyBlocks: make([][]int, 0, len(disk)/2),
	}
	for i, disk_segment := 0, 0; i < len(disk); i++ {
		ammount := int(disk[i] - '0')
		slice := formattedDisk.disk[disk_segment : disk_segment+ammount]

		if i&1 != 0 {
			formattedDisk.emptyBlocks = append(formattedDisk.emptyBlocks, slice)
		} else {
			formattedDisk.blocks = append(formattedDisk.blocks, slice)
		}

		disk_segment += ammount
	}
	return formattedDisk
}

func defragment(disk []int) []int {
	for start, end := 0, len(disk)-1; start < end; {
		for ; start < end && disk[start] != -1; start++ {
		}
		for ; start < end && disk[end] == -1; end-- {
		}
		if start < end {
			SwapInts(&disk[start], &disk[end])
		}
	}
	return disk
}

func defragmentByBlock(disk Disk) Disk {
	first_empty, last_empty := 0, len(disk.emptyBlocks)-1
	last_full := len(disk.blocks) - 1

	for ; 0 < last_full; last_full-- {
		for ; first_empty < last_empty && len(disk.emptyBlocks[first_empty]) == 0; first_empty++ {
		}
		for ; 0 < last_full && len(disk.blocks[last_full]) == 0; last_full-- {
		}

		write_point := first_empty
		read_point := last_full

		for ; write_point < last_empty && len(disk.emptyBlocks[write_point]) < len(disk.blocks[read_point]); write_point++ {
		}

		if read_point > write_point && len(disk.emptyBlocks[write_point]) >= len(disk.blocks[read_point]) {
			item_size := len(disk.blocks[read_point])
			for i := 0; i < item_size; i++ {
				SwapInts(&disk.emptyBlocks[write_point][0], &disk.blocks[read_point][0])
				disk.emptyBlocks[write_point] = disk.emptyBlocks[write_point][1:]
				disk.blocks[read_point] = disk.blocks[read_point][1:]
			}
		}
	}
	return disk
}

func checkSum(disk []int) int {
	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			sum += i * disk[i]
		}
	}

	return sum
}

func SwapInts(lhs, rhs *int) {
	*lhs, *rhs = *rhs, *lhs
}
