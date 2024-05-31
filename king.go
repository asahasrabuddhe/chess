package chess

// King represents a king.
type King struct {
	Piece
}

var kingMoves = Moves{
	{Condition: ConditionTrue, Move: MoveForward(1)},
	{Condition: ConditionTrue, Move: MoveBackward(1)},
	{Condition: ConditionTrue, Move: MoveLeft(1)},
	{Condition: ConditionTrue, Move: MoveRight(1)},
	{Condition: ConditionTrue, Move: MoveForwardLeft(1)},
	{Condition: ConditionTrue, Move: MoveForwardRight(1)},
	{Condition: ConditionTrue, Move: MoveBackwardLeft(1)},
	{Condition: ConditionTrue, Move: MoveBackwardRight(1)},
}

// NewKing creates a new King with the given position.
func NewKing(position string, color Color) (*King, error) {
	p, err := NewPiece(position, color, kingMoves)
	if err != nil {
		return nil, err
	}

	return &King{Piece: p}, nil
}
