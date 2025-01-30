package chess

// RookMaxMoves is the maximum number of possible moves for a rook. A rook can move in any direction for any number
// of squares. The maximum number of moves a rook can make is 14.
const RookMaxMoves = 14

// Rook represents a rook.
type Rook struct {
	Piece
}

var rookMoves = Moves{
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
}

// NewRook creates a new Rook with the given position.
func NewRook(position string, color Color) (*Rook, error) {
	p, err := NewPiece(position, color, &rookMoves, RookMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Rook{Piece: p}, nil
}
