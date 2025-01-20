package chess

import (
	"sync"
)

// MoveFunc is a function that takes a Position, Color, and Board and returns the new Position after completing the
// move. If the move is not possible, the function will return the original Position.
type MoveFunc func(Position, Color) Position

// MoveForward returns a slice of MoveFuncs that move the piece forward by the given number of steps.
func MoveForward(steps int) MoveFuncs {
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
	return MoveFuncs{Compound: false, Moves: out}
}

// MoveBackward returns a slice of MoveFuncs that move the piece backward by the given number of steps.
func MoveBackward(steps int) MoveFuncs {
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
	return MoveFuncs{Compound: false, Moves: out}
}

// MoveLeft returns a slice of MoveFuncs that move the piece to the left by the given number of steps.
func MoveLeft(steps int) MoveFuncs {
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
	return MoveFuncs{Compound: false, Moves: out}
}

// MoveRight returns a slice of MoveFuncs that move the piece to the right by the given number of steps.
func MoveRight(steps int) MoveFuncs {
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
	return MoveFuncs{Compound: false, Moves: out}
}

// MoveForwardLeft returns a slice of MoveFuncs that move the piece forward and to the left by the given number of steps.
func MoveForwardLeft(steps int) MoveFuncs {
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
	return MoveFuncs{Compound: false, Moves: out}
}

// MoveForwardRight returns a slice of MoveFuncs that move the piece forward and to the right by the given number of steps.
func MoveForwardRight(steps int) MoveFuncs {
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
	return MoveFuncs{Compound: false, Moves: out}
}

// MoveBackwardLeft returns a slice of MoveFuncs that move the piece backward and to the left by the given number of steps.
func MoveBackwardLeft(steps int) MoveFuncs {
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
	return MoveFuncs{Compound: false, Moves: out}
}

// MoveBackwardRight returns a slice of MoveFuncs that move the piece backward and to the right by the given number of steps.
func MoveBackwardRight(steps int) MoveFuncs {
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
	return MoveFuncs{Compound: false, Moves: out}
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
	Move      MoveFuncs
}

type MoveFuncs struct {
	Compound bool
	Moves    []MoveFunc
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
			if !move.Move.Compound {
				// Loop through the MoveFuncs for the move.
				var lastPosition Position
				for idx, f := range move.Move.Moves {
					// re-initialize the position and color for each iteration.
					newPosition := position
					// Apply the MoveFunc to the position for the given number of steps.
					for i := 0; i < idx+1; i++ {
						newPosition = f(newPosition, color)
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
				// Loop through the MoveFuncs for the move.
				for _, f := range move.Move.Moves {
					// Apply the MoveFunc to the position.
					newPosition = f(newPosition, color)
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

func CompoundMove(moves ...MoveFuncs) MoveFuncs {
	var length int
	for _, m := range moves {
		length += len(m.Moves)
	}
	var out = make([]MoveFunc, 0, length)
	for _, m := range moves {
		out = append(out, m.Moves...)
	}
	return MoveFuncs{Compound: true, Moves: out}
}
