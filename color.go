package main

type Color int

const (
	Black = iota
	White
)

func (c Color) String() string {
	return [...]string{"Black", "White"}[c]
}
