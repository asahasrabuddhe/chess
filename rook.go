package chess

// RookMaxMoves is the maximum number of possible moves for a rook. A rook can move in any direction for any number
// of squares. The maximum number of moves a rook can make is 14.
const RookMaxMoves = 14

// Rook represents a rook.
type Rook struct {
	Piece
}

var rookMoves = Moves{
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
}

// NewRook creates a new Rook with the given position.
func NewRook(position string, color Color) (*Rook, error) {
	p, err := NewPiece(position, color, &rookMoves, RookMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Rook{Piece: p}, nil
}
