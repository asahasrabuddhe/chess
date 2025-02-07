package chess

// BishopMaxMoves is the maximum number of possible moves for a bishop. A bishop can move diagonally for any number
// of squares. The maximum number of moves a bishop can make is 13.
const BishopMaxMoves = 13

// Bishop represents a bishop.
type Bishop struct {
	Piece
}

var bishopMoves = Moves{
	SimpleMove(MoveForwardLeft(1)), SimpleMove(MoveForwardLeft(2)),
	SimpleMove(MoveForwardLeft(3)), SimpleMove(MoveForwardLeft(4)),
	SimpleMove(MoveForwardLeft(5)), SimpleMove(MoveForwardLeft(6)),
	SimpleMove(MoveForwardLeft(7)), SimpleMove(MoveForwardLeft(8)),

	SimpleMove(MoveForwardRight(1)), SimpleMove(MoveForwardRight(2)),
	SimpleMove(MoveForwardRight(3)), SimpleMove(MoveForwardRight(4)),
	SimpleMove(MoveForwardRight(5)), SimpleMove(MoveForwardRight(6)),
	SimpleMove(MoveForwardRight(7)), SimpleMove(MoveForwardRight(8)),

	SimpleMove(MoveBackwardLeft(1)), SimpleMove(MoveBackwardLeft(2)),
	SimpleMove(MoveBackwardLeft(3)), SimpleMove(MoveBackwardLeft(4)),
	SimpleMove(MoveBackwardLeft(5)), SimpleMove(MoveBackwardLeft(6)),
	SimpleMove(MoveBackwardLeft(7)), SimpleMove(MoveBackwardLeft(8)),

	SimpleMove(MoveBackwardRight(1)), SimpleMove(MoveBackwardRight(2)),
	SimpleMove(MoveBackwardRight(3)), SimpleMove(MoveBackwardRight(4)),
	SimpleMove(MoveBackwardRight(5)), SimpleMove(MoveBackwardRight(6)),
	SimpleMove(MoveBackwardRight(7)), SimpleMove(MoveBackwardRight(8)),
}

// NewBishop creates a new Bishop with the given position.
func NewBishop(position string, color Color) (*Bishop, error) {
	p, err := NewPiece(position, color, &bishopMoves, BishopMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Bishop{Piece: p}, nil
}
