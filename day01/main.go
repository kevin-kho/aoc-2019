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
}
