package main

import (
	"bytes"
	"cmp"
	"fmt"
	"log"
	"math"
	"slices"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pos struct {
	X       int
	Y       int
	Radians float64
}

func (p *Pos) CalculateRadian() {
	rad := math.Atan2(float64(p.Y), float64(p.X))
	if rad < 0 {
		rad += 2 * math.Pi
	}

	p.Radians = rad
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
		{X: 0, Y: 1},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: -1, Y: 0},
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

	// Sort gcdVector by polar coordinates
	// Always positive radian
	var gcdVectors []Pos
	for gcdVector := range station.Asteroids {
		gcdVector.CalculateRadian()
		gcdVectors = append(gcdVectors, gcdVector)
	}
	slices.SortFunc(gcdVectors, func(a, b Pos) int {
		return cmp.Compare(a.Radians, b.Radians)
	})

	// Distribute Pos into cartesian quadrants
	var q1 []Pos
	var q2 []Pos
	var q3 []Pos
	var q4 []Pos
	for _, g := range gcdVectors {
		if 0 <= g.Radians && g.Radians <= math.Pi/2 {
			q1 = append(q1, g)
		} else if math.Pi/2 <= g.Radians && g.Radians <= math.Pi {
			q2 = append(q2, g)
		} else if math.Pi <= g.Radians && g.Radians <= (3*math.Pi/2) {
			q3 = append(q3, g)
		} else {
			q4 = append(q4, g)
		}
	}

	// order goes q1 -> q4 -> q3 -> q2
	var order []Pos
	slices.Reverse(q1)
	slices.Reverse(q2)
	slices.Reverse(q3)
	slices.Reverse(q4)
	order = append(order, q1...)
	order = append(order, q4...)
	order = append(order, q3...)
	order = append(order, q2...)
	for i, p := range order {
		p.Radians = 0
		order[i] = p
	}

	var count int
	i := 0

	// WIP: sort asteroids based on how close they are from the Station
	mp := station.Asteroids
	for vec, asteroids := range mp {
		slices.SortFunc(asteroids, func(a, b Pos) int {

			var aFactor int
			var bFactor int

			if vec.X == 0 {
				aFactor = (a.Y - station.Y) / vec.Y
				bFactor = (b.Y - station.Y) / vec.Y
			} else {
				aFactor = (a.X - station.X) / vec.X
				bFactor = (b.X - station.X) / vec.X
			}

			return cmp.Compare(aFactor, bFactor)
		})
		mp[vec] = asteroids
	}

	var destroyed Pos
	for count < 200 {
		vec := order[i]
		if len(mp[vec]) == 0 {
			i++
			i = i % len(order)
			continue
		}

		destroyed = mp[vec][0]
		mp[vec] = mp[vec][1:]

		count++
		i++
		i = i % len(order)
	}
	fmt.Println(destroyed)

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("inputExample2.txt")
	// data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	grid := CreateGrid(data)
	res := SolvePartOne(grid)
	fmt.Println(res.Pos, len(res.Asteroids))

	SolvePartTwo(res, grid)

}
