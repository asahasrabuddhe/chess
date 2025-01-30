package chess

type MoveDelta struct {
	rankDelta int
	fileDelta rune
}

func (m MoveDelta) ApplyTo(position Position, color Color) Position {
	newPosition := Position{
		Rank: position.Rank + m.rankDelta*int(color),
		File: rune(int(position.File) + int(m.fileDelta)*int(color)),
	}
	if !newPosition.isValid() {
		return position
	}
	return newPosition
}

func generateMove(steps, rankDelta int, fileDelta rune) MoveDelta {
	return MoveDelta{rankDelta: rankDelta * steps, fileDelta: rune(int(fileDelta) * steps)}
}

// MoveForward generates MoveFuncs to move a piece forward.
func MoveForward(steps int) MoveDelta {
	return generateMove(steps, 1, 0)
}

// MoveBackward generates MoveFuncs to move a piece backward.
func MoveBackward(steps int) MoveDelta {
	return generateMove(steps, -1, 0)
}

// MoveLeft generates MoveFuncs to move a piece left.
func MoveLeft(steps int) MoveDelta {
	return generateMove(steps, 0, 1)
}

// MoveRight generates MoveFuncs to move a piece right.
func MoveRight(steps int) MoveDelta {
	return generateMove(steps, 0, -1)
}

// MoveForwardLeft generates MoveFuncs to move a piece forward-left.
func MoveForwardLeft(steps int) MoveDelta {
	return generateMove(steps, 1, 1)
}

// MoveForwardRight generates MoveFuncs to move a piece forward-right.
func MoveForwardRight(steps int) MoveDelta {
	return generateMove(steps, 1, -1)
}

// MoveBackwardLeft generates MoveFuncs to move a piece backward-left.
func MoveBackwardLeft(steps int) MoveDelta {
	return generateMove(steps, -1, 1)
}

// MoveBackwardRight generates MoveFuncs to move a piece backward-right.
func MoveBackwardRight(steps int) MoveDelta {
	return generateMove(steps, -1, -1)
}

// Condition is a function that takes a Position, Color, and Board and returns a boolean value indicating if the
// condition is satisfied or not.
type Condition func(Position, Color, *Board) bool

// ConditionTrue is a Condition that always returns true.
func ConditionTrue(_ Position, _ Color, _ *Board) bool {
	return true
}

type Move struct {
	condition Condition
	delta     MoveDelta
}

func SimpleMove(condition Condition, delta MoveDelta) Move {
	return Move{condition: condition, delta: delta}
}

func CompoundMove(condition Condition, delta ...MoveDelta) Move {
	var mergedDelta MoveDelta
	for _, d := range delta {
		mergedDelta.rankDelta += d.rankDelta
		mergedDelta.fileDelta += d.fileDelta
	}
	return Move{condition: condition, delta: mergedDelta}
}

// Moves is a slice of Move.
type Moves []Move

// PossiblePositions returns a slice of Position that the piece can MoveDelta to based on the given board.
func (m Moves) PossiblePositions(piece Piece, board *Board) []Position {
	// Create a slice of Position to store the possible positions.
	var positions = make([]Position, piece.MaxPossibleMoves())
	var positionIndex int
	// Loop through the moves for the piece.
	for _, move := range m {
		// Get the position and color of the piece.
		position, color := piece.Position(), piece.Color()
		// Check if the condition for the MoveDelta is satisfied.
		if !move.condition(position, color, board) {
			continue
		}
		newPosition := move.delta.ApplyTo(position, color)
		if newPosition != position && newPosition.isValid() && board.PositionIsEmpty(newPosition) {
			positions[positionIndex] = newPosition
			positionIndex++
		}

	}
	// Return the slice of possible positions.
	return positions[:positionIndex]
}
