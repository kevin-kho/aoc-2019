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

type Wire struct {
	Commands []Command
	Points   map[Pos]bool
}

func (w *Wire) GetPoints() {
	curr := Pos{X: 0, Y: 0}

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
				for x := startX; x > startX+dx; x-- {
					curr.X = x
					w.Points[curr] = true
				}
			} else {
				for x := startX; x < startX+dx; x++ {
					curr.X = x
					w.Points[curr] = true
				}
			}

		}

		if dy != 0 {
			startY := curr.Y
			if dy < 0 {
				for y := startY; y > startY+dy; y-- {
					curr.Y = y
					w.Points[curr] = true
				}
			} else {
				for y := startY; y < startY+dy; y++ {
					curr.Y = y
					w.Points[curr] = true
				}
			}

		}

	}
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
		res = append(res, Wire{
			Commands: cmds,
			Points:   map[Pos]bool{},
		})
	}

	return res, nil

}

func SolvePartOne(wires []Wire) {

	for _, wire := range wires {
		wire.GetPoints()
	}

	collissions := make(map[Pos]bool)
	for pos := range wires[0].Points {
		if _, exists := wires[1].Points[pos]; exists {
			collissions[pos] = true
		}

	}
	fmt.Println(collissions)

}

func main() {
	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	wires, err := GetWires(data)
	if err != nil {
		log.Fatal(err)
	}

	SolvePartOne(wires)

}
