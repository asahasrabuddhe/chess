package chess

// RookMaxMoves is the maximum number of possible moves for a rook. A rook can move in any direction for any number
// of squares. The maximum number of moves a rook can make is 14.
const RookMaxMoves = 14

// Rook represents a rook.
type Rook struct {
	Piece
}

var rookMoves = Moves{
	SimpleMove(ConditionTrue, MoveForward(8)),
	SimpleMove(ConditionTrue, MoveBackward(8)),
	SimpleMove(ConditionTrue, MoveLeft(8)),
	SimpleMove(ConditionTrue, MoveRight(8)),
}

// NewRook creates a new Rook with the given position.
func NewRook(position string, color Color) (*Rook, error) {
	p, err := NewPiece(position, color, &rookMoves, RookMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Rook{Piece: p}, nil
}
