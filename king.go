package chess

// King represents a king.
type King struct {
	Piece
}

var kingMoves = Moves{
	{Condition: ConditionTrue, Move: []MoveFunc{MoveForward}},
	{Condition: ConditionTrue, Move: []MoveFunc{MoveBackward}},
	{Condition: ConditionTrue, Move: []MoveFunc{MoveLeft}},
	{Condition: ConditionTrue, Move: []MoveFunc{MoveRight}},
	{Condition: ConditionTrue, Move: []MoveFunc{MoveForwardLeft}},
	{Condition: ConditionTrue, Move: []MoveFunc{MoveForwardRight}},
	{Condition: ConditionTrue, Move: []MoveFunc{MoveBackwardLeft}},
	{Condition: ConditionTrue, Move: []MoveFunc{MoveBackwardRight}},
}

// NewKing creates a new King with the given position.
func NewKing(position string, color Color) (*King, error) {
	p, err := NewPiece(position, color, kingMoves)
	if err != nil {
		return nil, err
	}

	return &King{Piece: p}, nil
}
