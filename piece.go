package main

type Piece interface {
	Position() Pos
	Move(bd *Board, dest Pos) error
	Color() Color
	AssignColor(colr Color)
}
