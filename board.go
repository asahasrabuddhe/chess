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
func (b *Board) PositionIsEmpty(pos Position) bool {
	return b.pieces[pos.File][pos.Rank] == nil
}
