package chess

type MoveFunc func(Position, Color, Board) Position

func MoveForward(position Position, color Color, board Board) Position {
	newPosition := Position{}
	if color == White {
		if position.Rank-1 >= 0 {
			newPosition = Position{Rank: position.Rank - 1, File: position.File}
		}
	} else {
		if position.Rank+1 < 8 {
			newPosition = Position{Rank: position.Rank + 1, File: position.File}
		}
	}

	if !board.PositionIsEmpty(newPosition) {
		return position
	}

	return newPosition
}

type Condition func(Position, Color, Board) bool

func (c Condition) And(other Condition) Condition {
	return func(position Position, color Color, board Board) bool {
		return c(position, color, board) && other(position, color, board)
	}
}

func (c Condition) Or(other Condition) Condition {
	return func(position Position, color Color, board Board) bool {
		return c(position, color, board) || other(position, color, board)
	}
}

type Move struct {
	Condition Condition
	Move      []MoveFunc
}

type Moves []Move

func (m Moves) PossiblePositions(piece Piece, board Board) []Position {
	var positions []Position
	for _, move := range m {
		position, color := piece.Position(), piece.Color()
		if move.Condition(position, color, board) {
			for _, f := range move.Move {
				position = f(position, color, board)
			}
			if board.PositionIsEmpty(position) {
				positions = append(positions, position)
			}
		}
	}
	return positions
}
