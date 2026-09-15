package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pos struct {
	X int
	Y int
}

type Vector struct {
	Pos
}

type Station struct {
	Pos
	Asteroids map[Pos][]Pos
}

func CreateVector(src Pos, dst Pos) Vector {
	dx := dst.X - src.X
	dy := dst.Y - src.Y

	v := Vector{
		X: dx,
		Y: dy,
	}

	return v

}

func (p Vector) DivideVectorByGcd() Pos {
	gcd := p.X
	b := p.Y

	for b != 0 {
		gcd, b = b, gcd%b
	}
	if gcd < 0 {
		gcd = -gcd
	}

	newX := p.X / gcd
	newY := p.Y / gcd

	return Pos{
		X: newX,
		Y: newY,
	}

}

func Bfs(src Pos, grid [][]byte) map[Pos][]Pos {

	Y := len(grid)
	X := len(grid[0])
	Deltas := []Pos{
		{X: 1, Y: 0},
		{X: -1, Y: 0},
		{X: 0, Y: 1},
		{X: 0, Y: -1},
	}

	seenPos := make(map[Pos]bool)        // used to prevent backtracking
	seenGcdVector := make(map[Pos][]Pos) // key: gcdVector, value: slice of all vectors along that gcdVector

	queue := []Pos{src}

	for len(queue) > 0 {

		// Pop left
		pos := queue[0]
		queue = queue[1:]

		// case: already visited
		if seenPos[pos] {
			continue
		}

		// visit
		seenPos[pos] = true

		// case: asteroid; verify it can be seen
		if grid[pos.Y][pos.X] == '#' && src != pos {
			vec := CreateVector(src, pos)
			gcdVec := vec.DivideVectorByGcd()
			seenGcdVector[gcdVec] = append(seenGcdVector[gcdVec], pos)
		}

		// BFS outwards
		for _, d := range Deltas {
			newX := pos.X + d.X
			newY := pos.Y + d.Y
			if !(0 <= newX && newX < X) {
				continue
			}
			if !(0 <= newY && newY < Y) {
				continue
			}
			queue = append(queue, Pos{X: newX, Y: newY})
		}

	}

	return seenGcdVector

}

func CreateGrid(data []byte) [][]byte {
	grid := bytes.Split(data, []byte{'\n'})
	return grid

}

func SolvePartOne(grid [][]byte) Station {
	var res Station

	for y, row := range grid {
		for x := range row {
			if grid[y][x] == '#' {
				asteroidMp := Bfs(Pos{X: x, Y: y}, grid)
				if len(asteroidMp) > len(res.Asteroids) {
					res = Station{X: x, Y: y, Asteroids: asteroidMp}
				}
			}
		}
	}

	return res

}

func SolvePartTwo(station Station, grid [][]byte) {

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	// data, err := common.ReadInput("inputExample2.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	grid := CreateGrid(data)
	res := SolvePartOne(grid)
	fmt.Println(len(res.Asteroids))

	SolvePartTwo(res, grid)

}
