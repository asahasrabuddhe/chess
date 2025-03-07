package chess

// PawnMaxMoves represents the maximum number of moves a pawn can make. A pawn can move forward 2 squares if it is in
// its starting position, otherwise it can only move forward 1 square. It can also capture pieces diagonally.
const PawnMaxMoves = 4

// Pawn represents a pawn.
type Pawn struct {
	Piece
}

var pawnMoves = Moves{
	SimpleMove(MoveForward(1)),
	SimpleConditionalMove(func(position Position, color Color, board *Board) bool {
		return (color == White && position.Rank == 6) || (color == Black && position.Rank == 1)
	}, MoveForward(2)),
}

// NewPawn creates a new Pawn with the given position and color.
func NewPawn(position string, color Color) (*Pawn, error) {
	p, err := NewPiece(position, color, &pawnMoves, PawnMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Pawn{Piece: p}, nil
}
