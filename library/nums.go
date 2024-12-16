package aoc

func AbsDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func IsInRange(n, low, high int) bool {
	return n >= low && n < high
}

func CountDigits(num int) int {
	if num == 0 {
		return 1
	}

	digits := 0
	for num != 0 {
		num /= 10
		digits++
	}

	return digits
}

func SplitDigits(num, digits int) (int, int) {
	div := 1
	for i := 0; i < digits/2; i++ {
		div *= 10
	}

	a := num / div
	b := num % div

	return a, b
}
