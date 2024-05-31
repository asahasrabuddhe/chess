package chess

import "sort"

type Color int

const (
	Black Color = -1
	White Color = 1
)

// Piece represents a chess piece.
type Piece interface {
	// PossiblePositions returns a list of possible moves for the Piece.
	PossiblePositions(Board) []Position
	// Position returns the position of the Piece.
	Position() Position
	// Color returns the color of the Piece.
	Color() Color
	// PossibleMoves returns the possible moves for the Piece.
	PossibleMoves() Moves
}

type piece struct {
	position      Position
	color         Color
	possibleMoves Moves
}

// NewPiece creates a new Piece with the given position and color.
func NewPiece(position string, color Color, possibleMoves Moves) (Piece, error) {
	pos, err := ParsePosition(position)
	if err != nil {
		return nil, err
	}
	return &piece{position: pos, color: color, possibleMoves: possibleMoves}, nil
}

// PossiblePositions returns a list of possible moves for the piece.
func (p *piece) PossiblePositions(board Board) []Position {
	positions := p.possibleMoves.PossiblePositions(p, board)
	sort.Slice(positions, func(i, j int) bool {
		if p.color == White {
			return (positions[i].Rank > positions[j].Rank) && (positions[i].File > positions[j].File)
		} else {
			return (positions[i].Rank < positions[j].Rank) && (positions[i].File < positions[j].File)
		}
	})
	return positions
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
func (p *piece) PossibleMoves() Moves {
	return p.possibleMoves
}

//// Queen represents a queen.
//type Queen struct {
//	position Position
//}
//
//// NewQueen creates a new Queen with the given position.
//func NewQueen(position string) (*Queen, error) {
//	pos, err := ParsePosition(position)
//	if err != nil {
//		return nil, err
//	}
//	return &Queen{position: pos}, nil
//}
//
//// PossiblePositions returns a list of possible moves for the queen.
//func (q *Queen) PossiblePositions(board Board) []string {
//	up, down, left, right := q.position.File, 7-q.position.File, q.position.Rank, 7-q.position.Rank
//	var diagFwdLeft, diagFwdRight, diagBackLeft, diagBackRight int
//	for i := 1; i <= q.position.File && i <= q.position.Rank; i++ {
//		diagFwdLeft++
//	}
//	for i := 1; i <= q.position.File && i <= 7-q.position.Rank; i++ {
//		diagFwdRight++
//	}
//	for i := 1; i <= 7-q.position.File && i <= q.position.Rank; i++ {
//		diagBackLeft++
//	}
//	for i := 1; i <= 7-q.position.File && i <= 7-q.position.Rank; i++ {
//		diagBackRight++
//	}
//	// pre allocate the moves slice with the maximum possible moves
//	var moves = make([]string, 0, up+down+left+right+diagFwdLeft+diagFwdRight+diagBackLeft+diagBackRight)
//	// move forward across the board
//	for i := 1; i <= q.position.File; i++ {
//		newPos := Position{File: q.position.File - i, Rank: q.position.Rank}
//		moves = append(moves, newPos.String())
//	}
//	// move backward across the board
//	for i := 1; i <= 7-q.position.File; i++ {
//		newPos := Position{File: q.position.File + i, Rank: q.position.Rank}
//		moves = append(moves, newPos.String())
//	}
//	// move to the left across the board
//	for i := 1; i <= q.position.Rank; i++ {
//		newPos := Position{File: q.position.File, Rank: q.position.Rank - i}
//		moves = append(moves, newPos.String())
//	}
//	// move to the right across the board
//	for i := 1; i <= 7-q.position.Rank; i++ {
//		newPos := Position{File: q.position.File, Rank: q.position.Rank + i}
//		moves = append(moves, newPos.String())
//	}
//	// move diagonally forward to the left across the board
//	for i := 1; i <= q.position.File && i <= q.position.Rank; i++ {
//		newPos := Position{File: q.position.File - i, Rank: q.position.Rank - i}
//		moves = append(moves, newPos.String())
//	}
//	// move diagonally forward to the right across the board
//	for i := 1; i <= q.position.File && i <= 7-q.position.Rank; i++ {
//		newPos := Position{File: q.position.File - i, Rank: q.position.Rank + i}
//		moves = append(moves, newPos.String())
//	}
//	// move diagonally backward to the left across the board
//	for i := 1; i <= 7-q.position.File && i <= q.position.Rank; i++ {
//		newPos := Position{File: q.position.File + i, Rank: q.position.Rank - i}
//		moves = append(moves, newPos.String())
//	}
//	// move diagonally backward to the right across the board
//	for i := 1; i <= 7-q.position.File && i <= 7-q.position.Rank; i++ {
//		newPos := Position{File: q.position.File + i, Rank: q.position.Rank + i}
//		moves = append(moves, newPos.String())
//	}
//	return moves
//}
