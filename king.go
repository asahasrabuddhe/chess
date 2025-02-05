package chess

// KingMaxMoves is the maximum number of moves a king can make. Except for when the king is near the edges, it can
// make 8 possible moves. One in each direction.
const KingMaxMoves = 8

// King represents a king.
type King struct {
	Piece
}

var kingMoves = Moves{
	SimpleMove(MoveForward(1)),
	SimpleMove(MoveBackward(1)),
	SimpleMove(MoveLeft(1)),
	SimpleMove(MoveRight(1)),
	SimpleMove(MoveForwardLeft(1)),
	SimpleMove(MoveForwardRight(1)),
	SimpleMove(MoveBackwardLeft(1)),
	SimpleMove(MoveBackwardRight(1)),
}

// NewKing creates a new King with the given position.
func NewKing(position string, color Color) (*King, error) {
	p, err := NewPiece(position, color, &kingMoves, KingMaxMoves)
	if err != nil {
		return nil, err
	}

	return &King{Piece: p}, nil
}
