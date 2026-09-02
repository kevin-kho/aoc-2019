package main

import "fmt"

func SolvePartOne(lower int, upper int) int {

	var valid []int

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
				valid = append(valid, curr)
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

	var count int
	for _, num := range valid {
		if lower <= num && num <= upper {
			count++
		}
	}

	return count

}

func main() {

	res := SolvePartOne(359282, 820401)
	fmt.Println(res)

}
