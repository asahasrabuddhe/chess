package chess

type MoveFunc func(Piece, Board) Position

func MoveForward(piece Piece, board Board) Position {
	position, newPosition := piece.Position(), Position{}
	if piece.Color() == White {
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

type Condition func(Piece, Board) bool

func (c Condition) And(other Condition) Condition {
	return func(piece Piece, board Board) bool {
		return c(piece, board) && other(piece, board)
	}
}

func (c Condition) Or(other Condition) Condition {
	return func(piece Piece, board Board) bool {
		return c(piece, board) || other(piece, board)
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
		if move.Condition(piece, board) {
			for _, f := range move.Move {
				newPosition := f(piece, board)
				if !board.PositionIsEmpty(newPosition) {
					break
				}
				positions = append(positions, newPosition)
			}
		}
	}
	return positions
}
