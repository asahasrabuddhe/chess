package chess

// Pawn represents a pawn.
type Pawn struct {
	Piece
}

var pawnMoves = Moves{
	{
		Condition: func(position Position, color Color, _ *Board) bool {
			return (color == White && position.Rank == 6) || (color == Black && position.Rank == 1)
		},
		Move: MoveForward(2),
	},
	{
		Condition: ConditionTrue,
		Move:      MoveForward(1),
	},
}

// NewPawn creates a new Pawn with the given position and color.
func NewPawn(position string, color Color) (*Pawn, error) {
	p, err := NewPiece(position, color, &pawnMoves)
	if err != nil {
		return nil, err
	}

	return &Pawn{Piece: p}, nil
}
