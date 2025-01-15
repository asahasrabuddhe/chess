package chess

// QueenMaxMoves is the maximum number of possible moves for a queen. A queen can move in any direction for any number
// of squares. The maximum number of moves a queen can make is 27.
const QueenMaxMoves = 27

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
	p, err := NewPiece(position, color, &queenMoves, QueenMaxMoves)
	if err != nil {
		return nil, err
	}

	return &Queen{Piece: p}, nil
}
