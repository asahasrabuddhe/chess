package chess

// Pawn represents a pawn.
type Pawn struct {
	*piece
}

var pawnMoves = Moves{
	{
		Condition: func(piece Piece, board Board) bool {
			color := piece.Color()
			rank := piece.Position().Rank

			return (color == White && rank == 6) || (color == Black && rank == 1)
		},
		Move: []MoveFunc{MoveForward, MoveForward},
	},
	{
		Condition: func(piece Piece, board Board) bool {
			return true
		},
		Move: []MoveFunc{MoveForward},
	},
}

// NewPawn creates a new Pawn with the given position and color.
func NewPawn(position string, color Color) (*Pawn, error) {
	p, err := NewPiece(position, color, pawnMoves)
	if err != nil {
		return nil, err
	}

	return &Pawn{piece: p.(*piece)}, nil
}
