package chess

import (
	"sync"
)

// MoveFunc is a function that takes a Position and Color and returns the new Position after completing the Move. If
// the Move is not possible, the function will return the original Position.
type MoveFunc func(position Position, color Color) Position

// generateMoves generates a slice of MoveFunc that can be used to move a piece in a given direction represented by the
// rankDelta and fileDelta. The number of steps denotes the number of steps the piece can take in that direction.
func generateMoves(steps int, rankDelta, fileDelta int) []MoveFunc {
	var (
		out = make([]MoveFunc, steps)
		wg  sync.WaitGroup
	)
	wg.Add(steps)

	for i := 0; i < steps; i++ {
		go func(index int) {
			out[index] = func(position Position, color Color) Position {
				newPosition := Position{
					Rank: position.Rank + rankDelta*int(color),
					File: rune(int(position.File) + fileDelta*int(color)),
				}
				if !newPosition.isValid() {
					return position
				}
				return newPosition
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	return out
}

// MoveForward generates MoveFuncs to move a piece forward.
func MoveForward(steps int) []MoveFunc {
	return generateMoves(steps, 1, 0)
}

// MoveBackward generates MoveFuncs to move a piece backward.
func MoveBackward(steps int) []MoveFunc {
	return generateMoves(steps, -1, 0)
}

// MoveLeft generates MoveFuncs to move a piece left.
func MoveLeft(steps int) []MoveFunc {
	return generateMoves(steps, 0, 1)
}

// MoveRight generates MoveFuncs to move a piece right.
func MoveRight(steps int) []MoveFunc {
	return generateMoves(steps, 0, -1)
}

// MoveForwardLeft generates MoveFuncs to move a piece forward-left.
func MoveForwardLeft(steps int) []MoveFunc {
	return generateMoves(steps, 1, 1)
}

// MoveForwardRight generates MoveFuncs to move a piece forward-right.
func MoveForwardRight(steps int) []MoveFunc {
	return generateMoves(steps, 1, -1)
}

// MoveBackwardLeft generates MoveFuncs to move a piece backward-left.
func MoveBackwardLeft(steps int) []MoveFunc {
	return generateMoves(steps, -1, 1)
}

// MoveBackwardRight generates MoveFuncs to move a piece backward-right.
func MoveBackwardRight(steps int) []MoveFunc {
	return generateMoves(steps, -1, -1)
}

// Condition is a function that takes a Position, Color, and Board and returns a boolean value indicating if the
// condition is satisfied or not.
type Condition func(Position, Color, *Board) bool

// ConditionTrue is a Condition that always returns true.
func ConditionTrue(_ Position, _ Color, _ *Board) bool {
	return true
}

type MoveType int

const (
	Simple MoveType = iota
	Compound
)

// Move represents a move that a piece can make. It contains a Condition and a slice of MoveFuncs that represent the
// Move. The Condition is used to determine if the Move is possible or not.
type Move struct {
	Type      MoveType
	Condition Condition
	Move      []MoveFunc
}

// SimpleMove returns a Move with the given Condition and MoveFuncs.
func SimpleMove(c Condition, m []MoveFunc) Move {
	return Move{Type: Simple, Condition: c, Move: m}
}

// CompoundMove returns a Move with the given Condition and MoveFuncs.
func CompoundMove(c Condition, mf ...[]MoveFunc) Move {
	var length int
	for _, m := range mf {
		length += len(m)
	}
	var out = make([]MoveFunc, 0, length)
	for _, m := range mf {
		out = append(out, m...)
	}
	return Move{Type: Compound, Condition: c, Move: out}
}

// Moves is a slice of Move.
type Moves []Move

// PossiblePositions returns a slice of Position that the piece can Move to based on the given board.
func (m Moves) PossiblePositions(piece Piece, board *Board) []Position {
	// Create a slice of Position to store the possible positions.
	var positions = make([]Position, piece.MaxPossibleMoves())
	var positionIndex int
	// Loop through the moves for the piece.
	for _, move := range m {
		// Get the position and color of the piece.
		position, color := piece.Position(), piece.Color()
		// Check if the condition for the Move is satisfied.
		if !move.Condition(position, color, board) {
			continue
		}
		if move.Type == Simple {
			// Loop through the MoveFuncs for the Move.
			var lastPosition Position
			for idx, fn := range move.Move {
				// re-initialize the position and color for each iteration.
				newPosition := position
				// Apply the MoveFunc to the position for the given number of steps.
				for i := 0; i < idx+1; i++ {
					newPosition = fn(newPosition, color)
				}
				if newPosition == lastPosition {
					break
				}
				if newPosition != position || !board.PositionIsEmpty(newPosition) {
					positions[positionIndex], lastPosition = newPosition, newPosition
					positionIndex++
				}
			}
			continue
		}
		// re-initialize the position and color for each iteration.
		lastPosition, newPosition := position, position
		// Loop through the MoveFuncs for the Move.
		for _, fn := range move.Move {
			// Apply the MoveFunc to the position.
			newPosition = fn(newPosition, color)
			if newPosition == lastPosition {
				newPosition = position
				break
			}
			lastPosition = newPosition
		}
		// If the new position is not the original position, add it to the slice of possible positions.
		if newPosition != position || !board.PositionIsEmpty(newPosition) {
			positions[positionIndex] = newPosition
			positionIndex++
		}

	}
	// Return the slice of possible positions.
	return positions[:positionIndex]
}
