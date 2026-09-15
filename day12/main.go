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
	Z int
}

type Moon struct {
	Pos
	D Pos // Velocity
}

func (m *Moon) Move() {
	m.X += m.D.X
	m.Y += m.D.Y
	m.Z += m.D.Z
}

func (m Moon) GetPotentialEnergy() int {
	var res int
	res += common.IntAbs(m.X)
	res += common.IntAbs(m.Y)
	res += common.IntAbs(m.Z)
	return res
}

func (m Moon) GetKineticEnergy() int {
	var res int
	res += common.IntAbs(m.D.X)
	res += common.IntAbs(m.D.Y)
	res += common.IntAbs(m.D.Z)
	return res
}

func (m Moon) GetTotalEnergy() int {
	return m.GetPotentialEnergy() * m.GetKineticEnergy()
}

func SolvePartOne(moons []Moon, steps int) int {
	var totalEnergy int

	for range steps {
		// Apply Gravity
		for i := range moons {
			for j := i + 1; j < len(moons); j++ {

				m1 := &moons[i]
				m2 := &moons[j]

				if m1.X > m2.X {
					m1.D.X -= 1
					m2.D.X += 1
				} else if m1.X < m2.X {
					m1.D.X += 1
					m2.D.X -= 1
				}
				if m1.Y > m2.Y {
					m1.D.Y -= 1
					m2.D.Y += 1
				} else if m1.Y < m2.Y {
					m1.D.Y += 1
					m2.D.Y -= 1
				}
				if m1.Z > m2.Z {
					m1.D.Z -= 1
					m2.D.Z += 1
				} else if m1.Z < m2.Z {
					m1.D.Z += 1
					m2.D.Z -= 1
				}
			}
		}

		// Move Moons
		for i, m := range moons {
			m.Move()
			moons[i] = m
		}

	}
	// Totalize Energy after steps
	for _, m := range moons {
		totalEnergy += m.GetTotalEnergy()
	}

	return totalEnergy

}

func GetMoons(data []byte) ([]Moon, error) {
	var res []Moon
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryStrArr := strings.Split(string(entry), ",")

		xStr := entryStrArr[0]
		xStr = strings.TrimSpace(xStr)
		xStr = strings.TrimPrefix(xStr, "<")
		xStr = strings.TrimPrefix(xStr, "x=")
		xInt, err := strconv.Atoi(xStr)
		if err != nil {
			return res, err
		}

		yStr := entryStrArr[1]
		yStr = strings.TrimSpace(yStr)
		yStr = strings.TrimPrefix(yStr, "y=")
		yInt, err := strconv.Atoi(yStr)
		if err != nil {
			return res, err
		}

		zStr := entryStrArr[2]
		zStr = strings.TrimSpace(zStr)
		zStr = strings.TrimSuffix(zStr, ">")
		zStr = strings.TrimPrefix(zStr, "z=")
		zInt, err := strconv.Atoi(zStr)
		if err != nil {
			return res, err
		}

		res = append(res, Moon{
			X: xInt,
			Y: yInt,
			Z: zInt,
		})

	}

	return res, nil
}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	// data, err := common.ReadInput("inputExample2.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	moons, err := GetMoons(data)
	if err != nil {
		log.Fatal(err)
	}
	res := SolvePartOne(moons, 1000)
	fmt.Println(res)

}
