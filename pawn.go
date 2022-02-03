package main

import "math"

type Pawn struct {
	pos   Pos
	colr  Color
	moves []Pos
}

func NewPawn(pos Pos, colr Color) {
	return Pawn{pos: pos, colr: colr}
}

func (p Pawn) Position() Pos {
	return p.pos
}

func (p Pawn) Color() Color {
	return p.colr
}

func (p *Pawn) AssignColor(colr Color) {
	p.colr = colr
}

func (p *Pawn) Move(bd *Board, dest Pos) {
	mv1 := Pos{p.pos[0], p.pos[1] + 1*p.direction()}
	mv2 := Pos{p.pos[0], p.pos[1] + 2*p.direction()}
	strL := Pos{p.pos[0] - 1, p.pos[1] + 1*p.direction()}
	strR := Pos{p.pos[0] + 1, p.pos[1] + 1*p.direction()}
	if bd.fields[cand] == nil {
		append(p.moves, cand)
	}
	p.appendMoves(bd, mv1, mv2)
	p.appendHits(bd, strL, strR)
	p.pos = p

	// en passent: if the last move of an opponent pawn from the home row
	// by 2 beside this pawn the pawn can move diagonal behind it and take
	// the opponents pawn

	// if this pawn moves to the base row of the opponent player can
	// replace this pawn with a piece of choice

	// cannot move if king it is chess

	p.pos = dest
}

func (p Pawn) direction() int {
	if p.colr == White {
		return 1
	} else if p.colr == Black {
		return -1
	}
}

func (p Pawn) isPawnHome() bool {
	return (p.colr == White && p.pos[0] == 1) || (p.colr == Black && p.pos[0] == 6)
}

func (p Pawn) appendMoves(bd *Board, mvs ...Pos) {
	for mv := range mvs {
		if !p.isBlocked() {
			append(p.moves, mv)
		}
	}
}

func (p Pawn) isBlocked(bd *Board, mv Pos) bool {
	dy := int(math.Abs(float64(p.pos[1]) - float64(mv[1])))
	if dy == 2 && p.isPawnHome() {
		btw := Pos{mv[0], mv[1] + 1*p.direction()}
		bd.fieldIsEmpty(btw)
	}
	bd.fieldIsEmpty(mv)
	return true
}

func (p Pawn) appendHits(bd *Board, hits ...Pos) {
	for hit := range hits {
		if p.canHit(bd, hit) {
			append(p.moves, hit)
		}
	}
}

func (p Pawn) canHit(bd *Board, dest Pos) bool {
	fig := bd.fields[dest]
	if fig == nil {
		return false
	}
	return fig.Color() != p.colr
}

//#####
//# ♚ #
//#####
//     #####
//     #####
//     #####
