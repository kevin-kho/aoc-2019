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
	for r <= len(arr) {
		values := arr[l:r]
		res = append(res, CreateLayer(values))

		l = r
		r += count
	}

	return res

}

func CombineLayers(layers []Layer, width int, length int) Layer {

	count := width * length

	var values []int

	// Loop over row
	for i := 0; i < count; i++ {
		color := 2
		j := 0
		for color == 2 && j < len(layers) {
			color = layers[j].Values[i]
			j++
		}
		values = append(values, color)

	}

	return CreateLayer(values)
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

func SolvePartTwo(layer Layer, width int, height int) {
	count := width
	l := 0
	r := width
	var rows [][]int
	for r <= len(layer.Values) {
		row := layer.Values[l:r]
		rows = append(rows, row)

		l = r
		r += count

	}

	var rowsString [][]string

	for _, row := range rows {
		var rowString []string
		for _, val := range row {
			var code string
			switch val {
			case 0:
				code = "\u25A0"
			case 1:
				code = "\u25A1"
			}
			rowString = append(rowString, code)
		}
		rowsString = append(rowsString, rowString)
	}

	for _, r := range rowsString {
		fmt.Println(r)
	}

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

	combinedLayer := CombineLayers(layers, 25, 6)
	fmt.Println(len(combinedLayer.Values))
	SolvePartTwo(combinedLayer, 25, 6)

}
