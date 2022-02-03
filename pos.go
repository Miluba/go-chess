package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Pos [2]int

func (p Pos) String() string {
	x := p[0]
	y := p[1]
	switch x {
	case 0:
		return fmt.Sprintf("A%d", y)
	case 1:
		return fmt.Sprintf("B%d", y)
	case 2:
		return fmt.Sprintf("C%d", y)
	case 3:
		return fmt.Sprintf("D%d", y)
	case 4:
		return fmt.Sprintf("E%d", y)
	case 5:
		return fmt.Sprintf("F%d", y)
	case 6:
		return fmt.Sprintf("G%d", y)
	case 7:
		return fmt.Sprintf("H%d", y)
	default:
		return "Not on field"
	}
}

func Convert(pos string) (Pos, error) {
	if len(pos) != 2 {
		return [2]int{-1, -1}, errors.New("Invalid position string")
	}
	c := pos[:1]
	y, err := strconv.Atoi(pos[1:])
	if err != nil {
		return [2]int{-1, -1}, errors.New("Invalid position string")
	}
	switch c {
	case "A":
		return [2]int{0, y}, nil
	case "B":
		return [2]int{1, y}, nil
	case "C":
		return [2]int{2, y}, nil
	case "D":
		return [2]int{3, y}, nil
	case "E":
		return [2]int{4, y}, nil
	case "F":
		return [2]int{5, y}, nil
	case "G":
		return [2]int{6, y}, nil
	case "H":
		return [2]int{7, y}, nil
	default:
		return [2]int{-1, -1}, errors.New("Invalid position string")
	}
}
