package chess

// BishopMaxMoves is the maximum number of possible moves for a bishop. A bishop can move diagonally for any number
// of squares. The maximum number of moves a bishop can make is 13.
const BishopMaxMoves = 13

// Bishop represents a bishop.
type Bishop struct {
	Piece
}

var bishopMoves = Moves{
	SimpleMove(ConditionTrue, MoveForwardLeft(1)), SimpleMove(ConditionTrue, MoveForwardLeft(2)),
	SimpleMove(ConditionTrue, MoveForwardLeft(3)), SimpleMove(ConditionTrue, MoveForwardLeft(4)),
	SimpleMove(ConditionTrue, MoveForwardLeft(5)), SimpleMove(ConditionTrue, MoveForwardLeft(6)),
	SimpleMove(ConditionTrue, MoveForwardLeft(7)), SimpleMove(ConditionTrue, MoveForwardLeft(8)),

	SimpleMove(ConditionTrue, MoveForwardRight(1)), SimpleMove(ConditionTrue, MoveForwardRight(2)),
	SimpleMove(ConditionTrue, MoveForwardRight(3)), SimpleMove(ConditionTrue, MoveForwardRight(4)),
	SimpleMove(ConditionTrue, MoveForwardRight(5)), SimpleMove(ConditionTrue, MoveForwardRight(6)),
	SimpleMove(ConditionTrue, MoveForwardRight(7)), SimpleMove(ConditionTrue, MoveForwardRight(8)),

	SimpleMove(ConditionTrue, MoveBackwardLeft(1)), SimpleMove(ConditionTrue, MoveBackwardLeft(2)),
	SimpleMove(ConditionTrue, MoveBackwardLeft(3)), SimpleMove(ConditionTrue, MoveBackwardLeft(4)),
	SimpleMove(ConditionTrue, MoveBackwardLeft(5)), SimpleMove(ConditionTrue, MoveBackwardLeft(6)),
	SimpleMove(ConditionTrue, MoveBackwardLeft(7)), SimpleMove(ConditionTrue, MoveBackwardLeft(8)),

	SimpleMove(ConditionTrue, MoveBackwardRight(1)), SimpleMove(ConditionTrue, MoveBackwardRight(2)),
	SimpleMove(ConditionTrue, MoveBackwardRight(3)), SimpleMove(ConditionTrue, MoveBackwardRight(4)),
	SimpleMove(ConditionTrue, MoveBackwardRight(5)), SimpleMove(ConditionTrue, MoveBackwardRight(6)),
	SimpleMove(ConditionTrue, MoveBackwardRight(7)), SimpleMove(ConditionTrue, MoveBackwardRight(8)),
}

// NewBishop creates a new Bishop with the given position.
func NewBishop(position string, color Color) (*Bishop, error) {
	p, err := NewPiece(position, color, &bishopMoves, BishopMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Bishop{Piece: p}, nil
}
