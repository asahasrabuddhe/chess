package chess

type Board struct {
	pieces [8][8]Piece
}

func EmptyBoard() Board {
	return Board{}
}

func (b Board) PositionIsEmpty(pos Position) bool {
	return b.pieces[pos.File][pos.Rank] == nil
}

func (b Board) PositionIsValid(pos Position) bool {
	return pos.File >= 0 && pos.File < 8 && pos.Rank >= 0 && pos.Rank < 8
}
