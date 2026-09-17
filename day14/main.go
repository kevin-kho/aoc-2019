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

func (p *Product) ReduceReaction() {

	var coeff []int
	coeff = append(coeff, p.Quantity)
	for _, r := range p.Reagent {
		coeff = append(coeff, r.Quantity)
	}

	gcd := 0
	for _, c := range coeff {
		for c != 0 {
			gcd, c = c, gcd%c
		}
	}

	p.Quantity /= gcd
	for i, r := range p.Reagent {
		r.Quantity /= gcd
		p.Reagent[i] = r
	}

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

		p := Product{
			Element: productEle,
			Reagent: reagents,
		}
		// p.ReduceReaction()

		res = append(res, p)

	}

	return res, nil
}

func SolvePartOne(products []Product) {
	// map for quick lookup of product
	mp := make(map[string]Product)
	for _, p := range products {
		mp[p.Name] = p
	}
	var oreCount int

	var dfs func(curr Product, need int)
	dfs = func(curr Product, need int) {
		mul := 1
		if need > curr.Quantity {
			mul = need / curr.Quantity
		}
		fmt.Println(curr, curr.Quantity, need, mul)

		for _, r := range curr.Reagent {
			if r.Name == "ORE" {
				fmt.Println("adding ore: ", mul*r.Quantity)
				oreCount += mul * r.Quantity
			} else {
				dfs(mp[r.Name], r.Quantity*mul)

			}
		}
	}

	reactCt := make(map[string]int)
	var dfs2 func(curr Product, need int)
	dfs2 = func(curr Product, need int) {
		mul := 1
		if need > curr.Quantity {
			mul = need / curr.Quantity
			if need%curr.Quantity != 0 {
				mul++
			}
		}
		for _, r := range curr.Reagent {
			if r.Name == "ORE" {
				reactCt[curr.Name] += mul * curr.Quantity
			} else {
				dfs2(mp[r.Name], r.Quantity*mul)

			}
		}
	}

	isBaseEle := make(map[string]bool)
	for _, p := range products {
		if len(p.Reagent) == 1 && p.Reagent[0].Name == "ORE" {
			isBaseEle[p.Name] = true
		}
	}

	var dfs3 func(curr Product, need int)
	dfs3 = func(curr Product, need int) {

		// mul := 1
		// if need > curr.Quantity {
		// 	mul = need / curr.Quantity
		// 	if need%curr.Quantity != 0 {
		// 		mul++
		// 	}
		// }

		mul := need / curr.Quantity
		if need%curr.Quantity != 0 {
			mul++
		}

		for _, r := range curr.Reagent {
			if isBaseEle[r.Name] {
				reactCt[r.Name] += mul * r.Quantity
			} else {
				dfs3(mp[r.Name], r.Quantity*mul)
			}

		}
	}

	// dfs(mp["FUEL"], 1)
	// dfs2(mp["FUEL"], 1)
	dfs3(mp["FUEL"], 1)

	fmt.Println(reactCt)
	for ele := range reactCt {
		fmt.Println(mp[ele])
		for reactCt[ele] > 0 {
			reactCt[ele] -= mp[ele].Quantity
			oreCount += mp[ele].Reagent[0].Quantity
		}
	}
	fmt.Println(oreCount)

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	// data, err := common.ReadInput("inputExample2.txt")
	// data, err := common.ReadInput("inputExample3.txt")
	// data, err := common.ReadInput("inputExample4.txt")
	// data, err := common.ReadInput("inputExample5.txt")
	data, err := common.ReadInput("input.txt")
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
