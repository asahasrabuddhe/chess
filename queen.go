package chess

// QueenMaxMoves is the maximum number of possible moves for a queen. A queen can move in any direction for any number
// of squares. The maximum number of moves a queen can make is 27.
const QueenMaxMoves = 27

// Queen represents a queen.
type Queen struct {
	Piece
}

var queenMoves = Moves{
	SimpleMove(MoveForward(1)), SimpleMove(MoveForward(2)),
	SimpleMove(MoveForward(3)), SimpleMove(MoveForward(4)),
	SimpleMove(MoveForward(5)), SimpleMove(MoveForward(6)),
	SimpleMove(MoveForward(7)), SimpleMove(MoveForward(8)),

	SimpleMove(MoveBackward(1)), SimpleMove(MoveBackward(2)),
	SimpleMove(MoveBackward(3)), SimpleMove(MoveBackward(4)),
	SimpleMove(MoveBackward(5)), SimpleMove(MoveBackward(6)),
	SimpleMove(MoveBackward(7)), SimpleMove(MoveBackward(8)),

	SimpleMove(MoveLeft(1)), SimpleMove(MoveLeft(2)),
	SimpleMove(MoveLeft(3)), SimpleMove(MoveLeft(4)),
	SimpleMove(MoveLeft(5)), SimpleMove(MoveLeft(6)),
	SimpleMove(MoveLeft(7)), SimpleMove(MoveLeft(8)),

	SimpleMove(MoveRight(1)), SimpleMove(MoveRight(2)),
	SimpleMove(MoveRight(3)), SimpleMove(MoveRight(4)),
	SimpleMove(MoveRight(5)), SimpleMove(MoveRight(6)),
	SimpleMove(MoveRight(7)), SimpleMove(MoveRight(8)),

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

// NewQueen creates a new Queen with the given position.
func NewQueen(position string, color Color) (*Queen, error) {
	p, err := NewPiece(position, color, &queenMoves, QueenMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Queen{Piece: p}, nil
}
