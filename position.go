package chess // import "go.ajitem.com/chess"

import "fmt"

// Position represents a position on the chess board.
type Position struct {
	Row int
	Col int
}

// ParsePosition parses a Position from a given string. In case of an invalid position, it returns an error.
func ParsePosition(position string) (Position, error) {
	// If the position string has a length not equal to 2, the position is invalid.
	if len(position) != 2 {
		return Position{}, fmt.Errorf("invalid position: %s", position)
	}
	// If the first character of the position string is not between 'A' and 'H', the position is invalid.
	if position[0] < 'A' || position[0] > 'H' {
		return Position{}, fmt.Errorf("invalid position: %s", position)
	}
	// If the second character of the position string is not between '1' and '8', the position is invalid.
	if position[1] < '1' || position[1] > '8' {
		return Position{}, fmt.Errorf("invalid position: %s", position)
	}
	// If the position is valid, return the parsed position.
	return Position{
		// Subtract the first character of the position string from the ASCII value of 'A' to get the column index.
		Col: int(position[0] - 'A'),
		// Subtract the ASCII value of '8' from the second character of the position string to get the row index.
		Row: int('8' - position[1]),
	}, nil
}

// String returns the string representation of a Position.
func (p Position) String() string {
	// Convert the row and column indices to string.
	return string(rune('A'+p.Col)) + string(rune(8-p.Row+'0'))
}
