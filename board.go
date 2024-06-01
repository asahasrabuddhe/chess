package chess

// Board represents a chess board.
type Board struct {
	pieces [8][8]Piece
}

// EmptyBoard returns an empty chess board.
func EmptyBoard() *Board {
	return &Board{}
}

// PositionIsEmpty returns true if the position is empty.
func (b Board) PositionIsEmpty(pos Position) bool {
	return b.pieces[pos.File][pos.Rank] == nil
}

// PositionIsValid returns true if the position is valid.
func (b Board) PositionIsValid(pos Position) bool {
	return pos.File >= 0 && pos.File < 8 && pos.Rank >= 0 && pos.Rank < 8
}
