package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pos struct {
	X int
	Y int
}

func (p Pos) GetManhattanDistance() int {
	return common.IntAbs(p.X) + common.IntAbs(p.Y)
}

type Wire struct {
	Commands []Command
	Points   map[Pos]int
}

func (w *Wire) GetPoints() {
	curr := Pos{X: 0, Y: 0}
	time := 0

	for _, cmd := range w.Commands {
		dx := 0
		dy := 0
		switch cmd.Dir {
		case 'U':
			dy = cmd.Steps
		case 'R':
			dx = cmd.Steps
		case 'L':
			dx = -cmd.Steps
		case 'D':
			dy = -cmd.Steps
		}

		if dx != 0 {
			startX := curr.X
			if dx < 0 {
				for x := startX; x > startX+dx-1; x-- {
					curr.X = x
					if _, exists := w.Points[curr]; !exists {
						w.Points[curr] = time
					}
					time++
				}
			} else {
				for x := startX; x < startX+dx+1; x++ {
					curr.X = x
					if _, exists := w.Points[curr]; !exists {
						w.Points[curr] = time
					}
					time++
				}
			}

		}

		if dy != 0 {
			startY := curr.Y
			if dy < 0 {
				for y := startY; y > startY+dy-1; y-- {
					curr.Y = y
					if _, exists := w.Points[curr]; !exists {
						w.Points[curr] = time
					}
					time++
				}
			} else {
				for y := startY; y < startY+dy+1; y++ {
					curr.Y = y
					if _, exists := w.Points[curr]; !exists {
						w.Points[curr] = time
					}
					time++
				}
			}

		}

		// For loops progress time one too far
		time--

	}

	delete(w.Points, Pos{
		X: 0,
		Y: 0,
	})

}

type Command struct {
	Dir   Direction
	Steps int
}

type Direction byte

const (
	U Direction = 'U'
	R Direction = 'R'
	L Direction = 'L'
	D Direction = 'D'
)

func GetDirection(d byte) Direction {
	var res Direction
	switch d {
	case 'U':
		res = 'U'
	case 'R':
		res = 'R'
	case 'L':
		res = 'L'
	case 'D':
		res = 'D'
	}

	return res
}

func GetWires(data []byte) ([]Wire, error) {
	var res []Wire
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryArr := strings.Split(string(entry), ",")
		var cmds []Command
		for _, cmd := range entryArr {
			dir := cmd[0]
			steps, err := strconv.Atoi(cmd[1:])
			if err != nil {
				return res, err
			}
			cmds = append(cmds, Command{
				Dir:   GetDirection(dir),
				Steps: steps,
			})
		}

		wire := Wire{
			Commands: cmds,
			Points:   map[Pos]int{},
		}
		wire.GetPoints()

		res = append(res, wire)
	}

	return res, nil

}

func SolvePartOne(wires []Wire) int {

	collissions := make(map[Pos]bool)
	for pos := range wires[0].Points {
		if _, exists := wires[1].Points[pos]; exists {
			collissions[pos] = true
		}
	}

	// Set random value min
	var res int
	for c := range collissions {
		res = c.GetManhattanDistance()
		break
	}

	for c := range collissions {
		res = min(res, c.GetManhattanDistance())
	}

	return res

}

func SolvePartTwo(wires []Wire) int {

	collissions := make(map[Pos]bool)
	for pos := range wires[0].Points {
		if _, exists := wires[1].Points[pos]; exists {
			collissions[pos] = true
		}
	}

	// Set random value min
	var res int
	for c := range collissions {
		w0 := wires[0].Points[c]
		w1 := wires[1].Points[c]
		res = w0 + w1
		break
	}

	for c := range collissions {
		w0 := wires[0].Points[c]
		w1 := wires[1].Points[c]
		res = min(res, w0+w1)
	}

	return res

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	// data, err := common.ReadInput("inputExample2.txt")
	// data, err := common.ReadInput("inputExample3.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	wires, err := GetWires(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(wires)
	fmt.Println(res)

	res2 := SolvePartTwo(wires)
	fmt.Println(res2)

}
