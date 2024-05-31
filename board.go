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
