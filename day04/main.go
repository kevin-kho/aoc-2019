package main

import "fmt"

func GetDigitMap(num int) map[int]int {
	mp := make(map[int]int)
	for num > 0 {
		digit := num % 10
		num = num / 10
		mp[digit]++
	}

	return mp

}

func DigitMapHasTwo(mp map[int]int) bool {

	for _, ct := range mp {
		if ct == 2 {
			return true
		}
	}
	return false
}

func GetIncreasingAdjacents() []int {
	var res []int
	var recurse func(curr int, adjacent bool, digit int)
	recurse = func(curr int, adjacent bool, digit int) {
		// Add the digit
		curr = curr*10 + digit

		// exit condition: out of bounds
		if curr >= 1_000_000 {
			return
		}

		// exit condition: curr is out of bounds (six digit number with adj digit)
		if 99999 < curr && curr < 1_000_000 {
			if adjacent {
				res = append(res, curr)
			}
			return
		}

		// Increment the digit
		if digit < 9 {
			for i := 1; i < 10-digit; i++ {
				recurse(curr, adjacent, digit+i)
			}
		}

		// Keep the digit
		if curr <= 99999 {
			adjacent = true
		}
		recurse(curr, adjacent, digit)
	}
	for i := 1; i < 10; i++ {
		recurse(0, false, i)
	}

	return res
}

func SolvePartOne(lower int, upper int) int {

	var valid []int

	valid = GetIncreasingAdjacents()

	var count int
	for _, num := range valid {
		if lower <= num && num <= upper {
			count++
		}
	}

	return count

}

func SolvePartTwo(lower int, upper int) int {

	valid := GetIncreasingAdjacents()

	var count int
	for _, num := range valid {
		if lower <= num && num <= upper && DigitMapHasTwo(GetDigitMap(num)) {
			count++
		}
	}

	return count

}

func main() {

	res := SolvePartOne(359282, 820401)
	fmt.Println(res)

	res2 := SolvePartTwo(359282, 820401)
	fmt.Println(res2)

}
