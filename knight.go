package chess

// KnightMaxMoves is the maximum number of possible moves for a knight. A knight can move in an L-shape, two squares
// in one direction and one square in a perpendicular direction. The maximum number of moves a knight can make is 8.
const KnightMaxMoves = 8

// Knight represents a knight.
type Knight struct {
	Piece
}

var knightMoves = Moves{
	CompoundMove(MoveForward(2), MoveRight(1)),
	CompoundMove(MoveForward(2), MoveLeft(1)),
	CompoundMove(MoveBackward(2), MoveRight(1)),
	CompoundMove(MoveBackward(2), MoveLeft(1)),
	CompoundMove(MoveLeft(2), MoveForward(1)),
	CompoundMove(MoveLeft(2), MoveBackward(1)),
	CompoundMove(MoveRight(2), MoveForward(1)),
	CompoundMove(MoveRight(2), MoveBackward(1)),
}

// NewKnight creates a new Knight with the given position.
func NewKnight(position string, color Color) (*Knight, error) {
	p, err := NewPiece(position, color, &knightMoves, KnightMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Knight{Piece: p}, nil
}
