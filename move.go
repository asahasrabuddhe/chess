package chess

import (
	"slices"
)

type MoveFunc func(Position, Color, Board) Position

func MoveForward(steps int) []MoveFunc {
	var out = make([]MoveFunc, steps)
	for i := 0; i < steps; i++ {
		out[i] = func(position Position, color Color, board Board) Position {
			newPosition := position
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
	}
	return out
}

func MoveBackward(steps int) []MoveFunc {
	var out = make([]MoveFunc, steps)
	for i := 0; i < steps; i++ {
		out[i] = func(position Position, color Color, board Board) Position {
			newPosition := position
			if color == White {
				if position.Rank+1 < 8 {
					newPosition = Position{Rank: position.Rank + 1, File: position.File}
				}
			} else {
				if position.Rank-1 >= 0 {
					newPosition = Position{Rank: position.Rank - 1, File: position.File}
				}
			}

			if !board.PositionIsEmpty(newPosition) {
				return position
			}

			return newPosition
		}
	}
	return out
}

func MoveLeft(steps int) []MoveFunc {
	var out = make([]MoveFunc, steps)
	for i := 0; i < steps; i++ {
		out[i] = func(position Position, color Color, board Board) Position {
			newPosition := position
			if color == White {
				if position.File-1 >= 0 {
					newPosition = Position{Rank: position.Rank, File: position.File - 1}
				}
			} else {
				if position.File+1 < 8 {
					newPosition = Position{Rank: position.Rank, File: position.File + 1}
				}
			}

			if !board.PositionIsEmpty(newPosition) {
				return position
			}

			return newPosition
		}
	}
	return out
}

func MoveRight(steps int) []MoveFunc {
	var out = make([]MoveFunc, steps)
	for i := 0; i < steps; i++ {
		out[i] = func(position Position, color Color, board Board) Position {
			newPosition := position
			if color == White {
				if position.File+1 < 8 {
					newPosition = Position{Rank: position.Rank, File: position.File + 1}
				}
			} else {
				if position.File-1 >= 0 {
					newPosition = Position{Rank: position.Rank, File: position.File - 1}
				}
			}

			if !board.PositionIsEmpty(newPosition) {
				return position
			}

			return newPosition
		}
	}
	return out
}

func MoveForwardLeft(steps int) []MoveFunc {
	var out = make([]MoveFunc, steps)
	for i := 0; i < steps; i++ {
		out[i] = func(position Position, color Color, board Board) Position {
			newPosition := position
			if color == White {
				if position.Rank-1 >= 0 && position.File-1 >= 0 {
					newPosition = Position{Rank: position.Rank - 1, File: position.File - 1}
				}
			} else {
				if position.Rank+1 < 8 && position.File+1 < 8 {
					newPosition = Position{Rank: position.Rank + 1, File: position.File + 1}
				}
			}

			if !board.PositionIsEmpty(newPosition) {
				return position
			}

			return newPosition
		}
	}
	return out
}

func MoveForwardRight(steps int) []MoveFunc {
	var out = make([]MoveFunc, steps)
	for i := 0; i < steps; i++ {
		out[i] = func(position Position, color Color, board Board) Position {
			newPosition := position
			if color == White {
				if position.Rank-1 >= 0 && position.File+1 < 8 {
					newPosition = Position{Rank: position.Rank - 1, File: position.File + 1}
				}
			} else {
				if position.Rank+1 < 8 && position.File-1 >= 0 {
					newPosition = Position{Rank: position.Rank + 1, File: position.File - 1}
				}
			}

			if !board.PositionIsEmpty(newPosition) {
				return position
			}

			return newPosition
		}
	}
	return out
}

func MoveBackwardLeft(steps int) []MoveFunc {
	var out = make([]MoveFunc, steps)
	for i := 0; i < steps; i++ {
		out[i] = func(position Position, color Color, board Board) Position {
			newPosition := position
			if color == White {
				if position.Rank+1 < 8 && position.File-1 >= 0 {
					newPosition = Position{Rank: position.Rank + 1, File: position.File - 1}
				}
			} else {
				if position.Rank-1 >= 0 && position.File+1 < 8 {
					newPosition = Position{Rank: position.Rank - 1, File: position.File + 1}
				}
			}

			if !board.PositionIsEmpty(newPosition) {
				return position
			}

			return newPosition
		}
	}
	return out
}

func MoveBackwardRight(steps int) []MoveFunc {
	var out = make([]MoveFunc, steps)
	for i := 0; i < steps; i++ {
		out[i] = func(position Position, color Color, board Board) Position {
			newPosition := position
			if color == White {
				if position.Rank+1 < 8 && position.File+1 < 8 {
					newPosition = Position{Rank: position.Rank + 1, File: position.File + 1}
				}
			} else {
				if position.Rank-1 >= 0 && position.File-1 >= 0 {
					newPosition = Position{Rank: position.Rank - 1, File: position.File - 1}
				}
			}

			if !board.PositionIsEmpty(newPosition) {
				return position
			}

			return newPosition
		}
	}
	return out
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

func ConditionTrue(_ Position, _ Color, _ Board) bool {
	return true
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
			for idx, f := range move.Move {
				position, color = piece.Position(), piece.Color()
				for i := 0; i < idx+1; i++ {
					position = f(position, color, board)
				}
				if board.PositionIsEmpty(position) && board.PositionIsValid(position) && position != piece.Position() && !slices.Contains(positions, position) {
					positions = append(positions, position)
				}
			}
		}
	}
	return positions
}
