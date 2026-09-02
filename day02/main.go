package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

func GetIntCodes(data []byte) ([]int, error) {
	var res []int
	for b := range bytes.SplitSeq(data, []byte{','}) {
		i, err := strconv.Atoi(string(b))
		if err != nil {
			return res, err
		}
		res = append(res, i)
	}

	return res, nil

}

func SolvePartOne(intCodes []int) int {

	i := 0
	for intCodes[i] != 99 {

		a := intCodes[intCodes[i+1]]
		b := intCodes[intCodes[i+2]]

		var c int
		switch intCodes[i] {
		case 1:
			c = a + b
		case 2:
			c = a * b
		default:
			fmt.Println("boo")
		}

		dst := intCodes[i+3]
		intCodes[dst] = c

		i += 4
	}

	return intCodes[0]
}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	intCodes, err := GetIntCodes(data)
	if err != nil {
		log.Fatal(err)
	}

	intCodes[1] = 12
	intCodes[2] = 2
	res := SolvePartOne(intCodes)
	fmt.Println(res)

}
