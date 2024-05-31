package chess

// Queen represents a queen.
type Queen struct {
	Piece
}

var queenMoves = Moves{
	{Condition: ConditionTrue, Move: MoveForward(8)},
	{Condition: ConditionTrue, Move: MoveBackward(8)},
	{Condition: ConditionTrue, Move: MoveLeft(8)},
	{Condition: ConditionTrue, Move: MoveRight(8)},
	{Condition: ConditionTrue, Move: MoveForwardLeft(8)},
	{Condition: ConditionTrue, Move: MoveForwardRight(8)},
	{Condition: ConditionTrue, Move: MoveBackwardLeft(8)},
	{Condition: ConditionTrue, Move: MoveBackwardRight(8)},
}

// NewQueen creates a new Queen with the given position.
func NewQueen(position string, color Color) (*Queen, error) {
	p, err := NewPiece(position, color, queenMoves)
	if err != nil {
		return nil, err
	}

	return &Queen{Piece: p}, nil
}
