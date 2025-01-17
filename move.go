package chess

import (
	"sync"
)

// MoveFunc is a function that takes a Position, Color, and Board and returns the new Position after completing the
// move. If the move is not possible, the function will return the original Position.
type MoveFunc func(Position, Color, *Board) Position

// MoveForward returns a slice of MoveFuncs that move the piece forward by the given number of steps.
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
			out[index] = func(position Position, color Color, board *Board) Position {
				// Calculate the new position based on the color of the piece.
				newPosition := Position{Rank: position.Rank + (1 * int(color)), File: position.File}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() || !board.PositionIsEmpty(newPosition) {
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

// MoveBackward returns a slice of MoveFuncs that move the piece backward by the given number of steps.
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
			out[index] = func(position Position, color Color, board *Board) Position {
				// Calculate the new position based on the color of the piece.
				newPosition := Position{Rank: position.Rank - (1 * int(color)), File: position.File}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() || !board.PositionIsEmpty(newPosition) {
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

// MoveLeft returns a slice of MoveFuncs that move the piece to the left by the given number of steps.
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
			out[index] = func(position Position, color Color, board *Board) Position {
				// Calculate the new position based on the color of the piece.
				newPosition := Position{Rank: position.Rank, File: rune(int(position.File) + (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() || !board.PositionIsEmpty(newPosition) {
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

// MoveRight returns a slice of MoveFuncs that move the piece to the right by the given number of steps.
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
			out[index] = func(position Position, color Color, board *Board) Position {
				// Calculate the new position based on the color of the piece.
				newPosition := Position{Rank: position.Rank, File: rune(int(position.File) - (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() || !board.PositionIsEmpty(newPosition) {
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

// MoveForwardLeft returns a slice of MoveFuncs that move the piece forward and to the left by the given number of steps.
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
			out[index] = func(position Position, color Color, board *Board) Position {
				// Calculate the new position based on the color of the piece.
				newPosition := Position{Rank: position.Rank + (1 * int(color)), File: rune(int(position.File) + (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() || !board.PositionIsEmpty(newPosition) {
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

// MoveForwardRight returns a slice of MoveFuncs that move the piece forward and to the right by the given number of steps.
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
			out[index] = func(position Position, color Color, board *Board) Position {
				newPosition := Position{Rank: position.Rank + (1 * int(color)), File: rune(int(position.File) - (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() || !board.PositionIsEmpty(newPosition) {
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

// MoveBackwardLeft returns a slice of MoveFuncs that move the piece backward and to the left by the given number of steps.
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
			out[index] = func(position Position, color Color, board *Board) Position {
				newPosition := Position{Rank: position.Rank - (1 * int(color)), File: rune(int(position.File) + (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() || !board.PositionIsEmpty(newPosition) {
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

// MoveBackwardRight returns a slice of MoveFuncs that move the piece backward and to the right by the given number of steps.
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
			out[index] = func(position Position, color Color, board *Board) Position {
				newPosition := Position{Rank: position.Rank - (1 * int(color)), File: rune(int(position.File) - (1 * int(color)))}
				// If the new position is not empty, return the original position.
				if !newPosition.isValid() || !board.PositionIsEmpty(newPosition) {
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

// Move represents a move that a piece can make. It contains a Condition and a slice of MoveFuncs that represent the
// move. The Condition is used to determine if the move is possible or not.
type Move struct {
	Condition Condition
	Move      []MoveFunc
}

// Moves is a slice of Move.
type Moves []Move

// PossiblePositions returns a slice of Position that the piece can move to based on the given board.
func (m Moves) PossiblePositions(piece Piece, board *Board) []Position {
	// Create a slice of Position to store the possible positions.
	var positions = make([]Position, piece.MaxPossibleMoves())
	var positionIndex int
	// Loop through the moves for the piece.
	for _, move := range m {
		// Get the position and color of the piece.
		position, color := piece.Position(), piece.Color()
		// Check if the condition for the move is satisfied.
		if move.Condition(position, color, board) {
			// Loop through the MoveFuncs for the move.
			var lastPosition Position
			for idx, f := range move.Move {
				// re-initialize the position and color for each iteration.
				newPosition := position
				// Apply the MoveFunc to the position for the given number of steps.
				for i := 0; i < idx+1; i++ {
					newPosition = f(newPosition, color, board)
				}
				if newPosition == lastPosition {
					break
				}
				if newPosition != position {
					positions[positionIndex], lastPosition = newPosition, newPosition
					positionIndex++
				}
			}
		}
	}
	// Return the slice of possible positions.
	return positions[:positionIndex]
}
