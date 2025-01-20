package chess

import (
	"sync"
)

// MoveFunc is a function that takes a Position, Color, and Board and returns the new Position after completing the
// Move. If the Move is not possible, the function will return the original Position.
type MoveFunc func(Position, Color) Position

// MoveForward returns a slice of MoveFunc that Move the piece forward by the given number of steps.
func MoveForward(steps int) []MoveFunc {
	// Create a slice of MoveFuncs with the given number of steps.
	var (
		// Create a slice of MoveFuncs with the given number of steps.
		out = make([]MoveFunc, steps)
		// Create a WaitGroup to wait for all the goroutines to finish.
		wg sync.WaitGroup
	)
	// Add the number of steps to the WaitGroup.
	wg.Add(steps)
	// Loop through the slice and create a MoveFunc that moves the piece forward by the given number of steps.
	for i := 0; i < steps; i++ {
		go func(index int) {
			out[index] = func(position Position, color Color) Position {
				// Calculate the new position based on the color of the piece.
				newPosition := Position{Rank: position.Rank + (1 * int(color)), File: position.File}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() {
					return position
				}
				// Return the new position.
				return newPosition
			}
			// Signal that the goroutine has finished.
			wg.Done()
		}(i)
	}
	// Wait for all the goroutines to finish.
	wg.Wait()
	return out
}

// MoveBackward returns a slice of MoveFunc that Move the piece backward by the given number of steps.
func MoveBackward(steps int) []MoveFunc {
	// Create a slice of MoveFuncs with the given number of steps.
	var (
		// Create a slice of MoveFuncs with the given number of steps.
		out = make([]MoveFunc, steps)
		// Create a WaitGroup to wait for all the goroutines to finish.
		wg sync.WaitGroup
	)
	// Add the number of steps to the WaitGroup.
	wg.Add(steps)
	// Loop through the slice and create a MoveFunc that moves the piece forward by the given number of steps.
	for i := 0; i < steps; i++ {
		go func(index int) {
			out[index] = func(position Position, color Color) Position {
				// Calculate the new position based on the color of the piece.
				newPosition := Position{Rank: position.Rank - (1 * int(color)), File: position.File}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() {
					return position
				}
				// Return the new position.
				return newPosition
			}
			// Signal that the goroutine has finished.
			wg.Done()
		}(i)
	}
	// Wait for all the goroutines to finish.
	wg.Wait()
	return out
}

// MoveLeft returns a slice of MoveFunc that Move the piece to the left by the given number of steps.
func MoveLeft(steps int) []MoveFunc {
	// Create a slice of MoveFuncs with the given number of steps.
	var (
		// Create a slice of MoveFuncs with the given number of steps.
		out = make([]MoveFunc, steps)
		// Create a WaitGroup to wait for all the goroutines to finish.
		wg sync.WaitGroup
	)
	// Add the number of steps to the WaitGroup.
	wg.Add(steps)
	// Loop through the slice and create a MoveFunc that moves the piece forward by the given number of steps.
	for i := 0; i < steps; i++ {
		go func(index int) {
			out[index] = func(position Position, color Color) Position {
				// Calculate the new position based on the color of the piece.
				newPosition := Position{Rank: position.Rank, File: rune(int(position.File) + (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() {
					return position
				}
				// Return the new position.
				return newPosition
			}
			// Signal that the goroutine has finished.
			wg.Done()
		}(i)
	}
	// Wait for all the goroutines to finish.
	wg.Wait()
	return out
}

// MoveRight returns a slice of MoveFunc that Move the piece to the right by the given number of steps.
func MoveRight(steps int) []MoveFunc {
	// Create a slice of MoveFuncs with the given number of steps.
	var (
		// Create a slice of MoveFuncs with the given number of steps.
		out = make([]MoveFunc, steps)
		// Create a WaitGroup to wait for all the goroutines to finish.
		wg sync.WaitGroup
	)
	// Add the number of steps to the WaitGroup.
	wg.Add(steps)
	// Loop through the slice and create a MoveFunc that moves the piece forward by the given number of steps.
	for i := 0; i < steps; i++ {
		go func(index int) {
			out[index] = func(position Position, color Color) Position {
				// Calculate the new position based on the color of the piece.
				newPosition := Position{Rank: position.Rank, File: rune(int(position.File) - (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() {
					return position
				}
				// Return the new position.
				return newPosition
			}
			// Signal that the goroutine has finished.
			wg.Done()
		}(i)
	}
	// Wait for all the goroutines to finish.
	wg.Wait()
	return out
}

// MoveForwardLeft returns a slice of MoveFunc that Move the piece forward and to the left by the given number of steps.
func MoveForwardLeft(steps int) []MoveFunc {
	// Create a slice of MoveFuncs with the given number of steps.
	var (
		// Create a slice of MoveFuncs with the given number of steps.
		out = make([]MoveFunc, steps)
		// Create a WaitGroup to wait for all the goroutines to finish.
		wg sync.WaitGroup
	)
	// Add the number of steps to the WaitGroup.
	wg.Add(steps)
	// Loop through the slice and create a MoveFunc that moves the piece forward by the given number of steps.
	for i := 0; i < steps; i++ {
		go func(index int) {
			out[index] = func(position Position, color Color) Position {
				// Calculate the new position based on the color of the piece.
				newPosition := Position{Rank: position.Rank + (1 * int(color)), File: rune(int(position.File) + (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() {
					return position
				}
				// Return the new position.
				return newPosition
			}
			// Signal that the goroutine has finished.
			wg.Done()
		}(i)
	}
	// Wait for all the goroutines to finish.
	wg.Wait()
	return out
}

// MoveForwardRight returns a slice of MoveFunc that Move the piece forward and to the right by the given number of steps.
func MoveForwardRight(steps int) []MoveFunc {
	// Create a slice of MoveFuncs with the given number of steps.
	var (
		// Create a slice of MoveFuncs with the given number of steps.
		out = make([]MoveFunc, steps)
		// Create a WaitGroup to wait for all the goroutines to finish.
		wg sync.WaitGroup
	)
	// Add the number of steps to the WaitGroup.
	wg.Add(steps)
	// Loop through the slice and create a MoveFunc that moves the piece forward by the given number of steps.
	for i := 0; i < steps; i++ {
		go func(index int) {
			out[index] = func(position Position, color Color) Position {
				newPosition := Position{Rank: position.Rank + (1 * int(color)), File: rune(int(position.File) - (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() {
					return position
				}
				// Return the new position.
				return newPosition
			}
			// Signal that the goroutine has finished.
			wg.Done()
		}(i)
	}
	// Wait for all the goroutines to finish.
	wg.Wait()
	return out
}

// MoveBackwardLeft returns a slice of MoveFunc that Move the piece backward and to the left by the given number of steps.
func MoveBackwardLeft(steps int) []MoveFunc {
	// Create a slice of MoveFuncs with the given number of steps.
	var (
		// Create a slice of MoveFuncs with the given number of steps.
		out = make([]MoveFunc, steps)
		// Create a WaitGroup to wait for all the goroutines to finish.
		wg sync.WaitGroup
	)
	// Add the number of steps to the WaitGroup.
	wg.Add(steps)
	// Loop through the slice and create a MoveFunc that moves the piece forward by the given number of steps.
	for i := 0; i < steps; i++ {
		go func(index int) {
			out[index] = func(position Position, color Color) Position {
				newPosition := Position{Rank: position.Rank - (1 * int(color)), File: rune(int(position.File) + (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() {
					return position
				}
				// Return the new position.
				return newPosition
			}
			// Signal that the goroutine has finished.
			wg.Done()
		}(i)
	}
	// Wait for all the goroutines to finish.
	wg.Wait()
	return out
}

// MoveBackwardRight returns a slice of MoveFunc that Move the piece backward and to the right by the given number of steps.
func MoveBackwardRight(steps int) []MoveFunc {
	// Create a slice of MoveFuncs with the given number of steps.
	var (
		// Create a slice of MoveFuncs with the given number of steps.
		out = make([]MoveFunc, steps)
		// Create a WaitGroup to wait for all the goroutines to finish.
		wg sync.WaitGroup
	)
	// Add the number of steps to the WaitGroup.
	wg.Add(steps)
	// Loop through the slice and create a MoveFunc that moves the piece forward by the given number of steps.
	for i := 0; i < steps; i++ {
		go func(index int) {
			out[index] = func(position Position, color Color) Position {
				newPosition := Position{Rank: position.Rank - (1 * int(color)), File: rune(int(position.File) - (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() {
					return position
				}
				// Return the new position.
				return newPosition
			}
			// Signal that the goroutine has finished.
			wg.Done()
		}(i)
	}
	// Wait for all the goroutines to finish.
	wg.Wait()
	return out
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

// Move represents a Move that a piece can make. It contains a Condition and a slice of MoveFuncs that represent the
// Move. The Condition is used to determine if the Move is possible or not.
type Move struct {
	Type      MoveType
	Condition Condition
	Move      []MoveFunc
}

func SimpleMove(c Condition, m []MoveFunc) Move {
	return Move{Type: Simple, Condition: c, Move: m}
}

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
		if move.Condition(position, color, board) {
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
			} else {
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
		}
	}
	// Return the slice of possible positions.
	return positions[:positionIndex]
}
