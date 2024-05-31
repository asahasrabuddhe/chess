package chess

// Pawn represents a pawn.
type Pawn struct {
	*piece
}

var pawnMoves = Moves{
	{
		Condition: func(position Position, color Color, board Board) bool {
			return (color == White && position.Rank == 6) || (color == Black && position.Rank == 1)
		},
		Move: []MoveFunc{MoveForward, MoveForward},
	},
	{
		Condition: func(position Position, color Color, board Board) bool {
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
