package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

func ApplyPattern(i int, nums []int) int {
	j := 0
	n := 0
	basePattern := []int{0, 1, 0, -1}
	skipped := false
	ct := 0
	var resArr []int

	for n < len(nums) {
		if !skipped {
			ct++
			skipped = true
			continue
		}

		if ct == i+1 {
			ct = 0
			j++
		}
		j = j % len(basePattern)
		mul := basePattern[j]
		// fmt.Println(nums[n], mul)
		resArr = append(resArr, nums[n]*mul)
		ct++
		n++
	}

	var res int
	for _, n := range resArr {
		res += n
	}
	res = common.IntAbs(res % 10)
	return res

}

func GetIntArr(i int) []int {
	var res []int
	for i > 0 {
		digit := i % 10
		res = append(res, digit)

		i = i / 10
	}

	slices.Reverse(res)
	return res
}

func GetIntArrFromStr(i string) ([]int, error) {
	var res []int
	for j := range i {
		k, err := strconv.Atoi(i[j : j+1])
		if err != nil {
			return res, err
		}
		res = append(res, k)
	}

	return res, nil

}

func SolvePartOne(intArr []int, phases int) {
	curr := slices.Clone(intArr)
	nxt := slices.Clone(intArr)
	for range phases {

		for i := range curr {
			n := ApplyPattern(i, curr)
			nxt[i] = n
		}
		curr = nxt

	}

	fmt.Println(curr[:8])

}

func main() {
	// arr := GetIntArr(12345678)
	// arr, err := GetIntArrFromStr("12345678")
	// arr, err := GetIntArrFromStr("80871224585914546619083218645595")
	// arr, err := GetIntArrFromStr("19617804207202209144916044189917")
	// arr, err := GetIntArrFromStr("69317163492948606335995924319873")
	arr, err := GetIntArrFromStr("59791871295565763701016897619826042828489762561088671462844257824181773959378451545496856546977738269316476252007337723213764111739273853838263490797537518598068506295920453784323102711076199873965167380615581655722603274071905196479183784242751952907811639233611953974790911995969892452680719302157414006993581489851373437232026983879051072177169134936382717591977532100847960279215345839529957631823999672462823375150436036034669895698554251454360619461187935247975515899240563842707592332912229870540467459067349550810656761293464130493621641378182308112022182608407992098591711589507803865093164025433086372658152474941776320203179747991102193608")

	if err != nil {
		log.Fatal(err)
	}
	SolvePartOne(arr, 100)
	// ApplyPattern(1, []int{1, 2, 3, 4, 5, 6, 7, 8})

}
