package chess

// QueenMaxMoves is the maximum number of possible moves for a queen. A queen can move in any direction for any number
// of squares. The maximum number of moves a queen can make is 27.
const QueenMaxMoves = 27

// Queen represents a queen.
type Queen struct {
	Piece
}

var queenMoves = Moves{
	SimpleMove(ConditionTrue, MoveForward(8)),
	SimpleMove(ConditionTrue, MoveBackward(8)),
	SimpleMove(ConditionTrue, MoveLeft(8)),
	SimpleMove(ConditionTrue, MoveRight(8)),
	SimpleMove(ConditionTrue, MoveForwardLeft(8)),
	SimpleMove(ConditionTrue, MoveForwardRight(8)),
	SimpleMove(ConditionTrue, MoveBackwardLeft(8)),
	SimpleMove(ConditionTrue, MoveBackwardRight(8)),
}

// NewQueen creates a new Queen with the given position.
func NewQueen(position string, color Color) (*Queen, error) {
	p, err := NewPiece(position, color, &queenMoves, QueenMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Queen{Piece: p}, nil
}
