package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Layer struct {
	Values      []int
	ValuesCount map[int]int
}

func CreateLayer(values []int) Layer {

	mp := make(map[int]int)
	for _, v := range values {
		mp[v]++
	}

	return Layer{
		Values:      values,
		ValuesCount: mp,
	}
}

func CreateIntArr(data []byte) ([]int, error) {
	var res []int
	for _, b := range data {
		i, err := strconv.Atoi(string(b))
		if err != nil {
			return res, err
		}
		res = append(res, i)
	}

	return res, nil

}

func CreateLayers(arr []int, width int, length int) []Layer {
	var res []Layer

	count := width * length

	l := 0
	r := count
	for r < len(arr) {
		values := arr[l:r]
		res = append(res, CreateLayer(values))

		l = r
		r += count
	}

	return res

}

func SolvePartOne(layers []Layer, width int, height int) int {

	var res int
	zeroCount := width * height
	for _, l := range layers {
		if l.ValuesCount[0] < zeroCount {
			zeroCount = l.ValuesCount[0]
			res = l.ValuesCount[1] * l.ValuesCount[2]
		}
	}

	return res

}

func main() {
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	intArr, err := CreateIntArr(data)
	if err != nil {
		log.Fatal(err)
	}
	layers := CreateLayers(intArr, 25, 6)

	res := SolvePartOne(layers, 25, 6)
	fmt.Println(res)

}
