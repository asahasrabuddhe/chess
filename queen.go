package chess

// QueenMaxMoves is the maximum number of possible moves for a queen. A queen can move in any direction for any number
// of squares. The maximum number of moves a queen can make is 27.
const QueenMaxMoves = 27

// Queen represents a queen.
type Queen struct {
	Piece
}

var queenMoves = Moves{
	SimpleMove(ConditionTrue, MoveForward(1)), SimpleMove(ConditionTrue, MoveForward(2)),
	SimpleMove(ConditionTrue, MoveForward(3)), SimpleMove(ConditionTrue, MoveForward(4)),
	SimpleMove(ConditionTrue, MoveForward(5)), SimpleMove(ConditionTrue, MoveForward(6)),
	SimpleMove(ConditionTrue, MoveForward(7)), SimpleMove(ConditionTrue, MoveForward(8)),

	SimpleMove(ConditionTrue, MoveBackward(1)), SimpleMove(ConditionTrue, MoveBackward(2)),
	SimpleMove(ConditionTrue, MoveBackward(3)), SimpleMove(ConditionTrue, MoveBackward(4)),
	SimpleMove(ConditionTrue, MoveBackward(5)), SimpleMove(ConditionTrue, MoveBackward(6)),
	SimpleMove(ConditionTrue, MoveBackward(7)), SimpleMove(ConditionTrue, MoveBackward(8)),

	SimpleMove(ConditionTrue, MoveLeft(1)), SimpleMove(ConditionTrue, MoveLeft(2)),
	SimpleMove(ConditionTrue, MoveLeft(3)), SimpleMove(ConditionTrue, MoveLeft(4)),
	SimpleMove(ConditionTrue, MoveLeft(5)), SimpleMove(ConditionTrue, MoveLeft(6)),
	SimpleMove(ConditionTrue, MoveLeft(7)), SimpleMove(ConditionTrue, MoveLeft(8)),

	SimpleMove(ConditionTrue, MoveRight(1)), SimpleMove(ConditionTrue, MoveRight(2)),
	SimpleMove(ConditionTrue, MoveRight(3)), SimpleMove(ConditionTrue, MoveRight(4)),
	SimpleMove(ConditionTrue, MoveRight(5)), SimpleMove(ConditionTrue, MoveRight(6)),
	SimpleMove(ConditionTrue, MoveRight(7)), SimpleMove(ConditionTrue, MoveRight(8)),

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

// NewQueen creates a new Queen with the given position.
func NewQueen(position string, color Color) (*Queen, error) {
	p, err := NewPiece(position, color, &queenMoves, QueenMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Queen{Piece: p}, nil
}
