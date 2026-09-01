package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

func GetMasses(data []byte) ([]int, error) {
	var res []int
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {

		i, err := strconv.Atoi(string(entry))
		if err != nil {
			return res, err
		}

		res = append(res, i)
	}

	return res, nil

}

func SolvePartOne(masses []int) int {
	var res int
	for _, m := range masses {
		res += m/3 - 2
	}

	return res
}

func FindFuelRecursive(fuel int) int {

	if fuel <= 0 {
		return 0
	}

	ff := max(fuel/3-2, 0)
	return ff + FindFuelRecursive(ff)

}

func SolvePartTwo(masses []int) int {
	var res int
	for _, m := range masses {
		res += FindFuelRecursive(m)
	}
	return res

}

func main() {
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	masses, err := GetMasses(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(masses)
	fmt.Println(res)

	res2 := SolvePartTwo(masses)
	fmt.Println(res2)
}
