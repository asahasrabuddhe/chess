package chess

type Color int

const (
	Black Color = -1
	White Color = 1
)

// Piece represents a chess piece.
type Piece interface {
	// PossiblePositions returns a list of possible moves for the Piece.
	PossiblePositions(*Board) []Position
	// Position returns the position of the Piece.
	Position() Position
	// Color returns the color of the Piece.
	Color() Color
	// PossibleMoves returns the possible moves for the Piece.
	PossibleMoves() *Moves
	// MaxPossibleMoves returns the maximum number of possible moves for the Piece.
	MaxPossibleMoves() int
}

type piece struct {
	position         Position
	color            Color
	possibleMoves    *Moves
	maxPossibleMoves int
}

// NewPiece creates a new Piece with the given position and color.
func NewPiece(position string, color Color, possibleMoves *Moves, maxPossibleMoves int) (Piece, error) {
	pos, err := ParsePosition(position)
	if err != nil {
		return nil, err
	}
	return &piece{position: pos, color: color, possibleMoves: possibleMoves, maxPossibleMoves: maxPossibleMoves}, nil
}

// PossiblePositions returns a list of possible moves for the piece.
func (p *piece) PossiblePositions(board *Board) []Position {
	return p.possibleMoves.PossiblePositions(p, board)
}

// Position returns the position of the piece.
func (p *piece) Position() Position {
	return p.position
}

// Color returns the color of the piece.
func (p *piece) Color() Color {
	return p.color
}

// PossibleMoves returns the possible moves for the piece.
func (p *piece) PossibleMoves() *Moves {
	return p.possibleMoves
}

// MaxPossibleMoves returns the maximum number of possible moves for the piece.
func (p *piece) MaxPossibleMoves() int {
	return p.maxPossibleMoves
}
