package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

func BuildOrbitMap(data []byte) map[string][]string {
	mp := make(map[string][]string)

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {

		entryStrArr := strings.Split(string(entry), ")")
		planet := entryStrArr[0]
		orbiter := entryStrArr[1]

		mp[planet] = append(mp[planet], orbiter)
	}

	return mp
}

func BuildOrbitMapBidirectional(data []byte) map[string][]string {
	mp := make(map[string][]string)

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {

		entryStrArr := strings.Split(string(entry), ")")
		planet := entryStrArr[0]
		orbiter := entryStrArr[1]

		mp[planet] = append(mp[planet], orbiter)
		mp[orbiter] = append(mp[orbiter], planet)
	}

	return mp
}

func SolvePartOne(orbitMap map[string][]string) int {
	var orbits int

	var dfs func(curr string, steps int)
	dfs = func(curr string, steps int) {

		// No exit condition
		// Input guaranteed to not cycle

		// add to orbits
		orbits += steps

		// DFS to next one
		for _, nxt := range orbitMap[curr] {
			dfs(nxt, steps+1)
		}

	}

	dfs("COM", 0)

	return orbits

}

func SolvePartTwo(orbitMapBidirectional map[string][]string) int {

	seen := make(map[string]bool)
	res := len(orbitMapBidirectional)
	fmt.Println(res)

	var dfs func(curr string, steps int)
	dfs = func(curr string, steps int) {
		// Case: already seen
		if seen[curr] {
			return
		}

		// Exit condition
		if curr == "SAN" {
			res = min(res, steps)
			return
		}

		seen[curr] = true
		for _, nxt := range orbitMapBidirectional[curr] {
			dfs(nxt, steps+1)
		}
		seen[curr] = false

	}

	dfs("YOU", 0)

	// Going from the planet YOU is orbiting to planet SAN is orbiting
	return res - 2
}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	mp := BuildOrbitMap(data)

	res := SolvePartOne(mp)
	fmt.Println(res)

	mpBidirectional := BuildOrbitMapBidirectional(data)
	res2 := SolvePartTwo(mpBidirectional)
	fmt.Println(res2)

}
