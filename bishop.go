package chess

// BishopMaxMoves is the maximum number of possible moves for a bishop. A bishop can move diagonally for any number
// of squares. The maximum number of moves a bishop can make is 13.
const BishopMaxMoves = 13

// Bishop represents a bishop.
type Bishop struct {
	Piece
}

var bishopMoves = Moves{
	SimpleMove(ConditionTrue, MoveForwardLeft(8)),
	SimpleMove(ConditionTrue, MoveForwardRight(8)),
	SimpleMove(ConditionTrue, MoveBackwardLeft(8)),
	SimpleMove(ConditionTrue, MoveBackwardRight(8)),
}

// NewBishop creates a new Bishop with the given position.
func NewBishop(position string, color Color) (*Bishop, error) {
	p, err := NewPiece(position, color, &bishopMoves, BishopMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Bishop{Piece: p}, nil
}
