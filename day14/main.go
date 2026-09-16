package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Element struct {
	Name     string
	Quantity int
}

type Product struct {
	Element
	Reagent []Element
}

func GetElement(ele string) (Element, error) {
	var res Element
	ele = strings.TrimSpace(ele)

	eleArr := strings.Split(ele, " ")
	qty, err := strconv.Atoi(eleArr[0])
	if err != nil {
		return res, err
	}
	name := eleArr[1]

	res.Name = name
	res.Quantity = qty
	return res, nil
}

func GetProducts(data []byte) ([]Product, error) {
	var res []Product
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryStrArr := strings.Split(string(entry), "=>")
		var reagents []Element
		for _, e := range strings.Split(entryStrArr[0], ",") {
			r, err := GetElement(e)
			if err != nil {
				return res, err
			}
			reagents = append(reagents, r)
		}

		productEle, err := GetElement(entryStrArr[1])
		if err != nil {
			return res, err
		}

		res = append(res, Product{
			Element: productEle,
			Reagent: reagents,
		})

	}

	return res, nil
}

func SolvePartOne(products []Product) {
	// map for quick lookup of product
	mp := make(map[string]Product)
	for _, p := range products {
		mp[p.Name] = p
	}

	fmt.Println(mp)

}

func main() {
	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	products, err := GetProducts(data)
	if err != nil {
		log.Fatal(err)
	}

	SolvePartOne(products)

}
