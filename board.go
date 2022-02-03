package main

type Board struct {
	fields map[Pos]Piece
}

func (b Board) fieldIsEmpty(f Pos) bool {
	if b.fields[f] != nil {
		return false
	}
}
