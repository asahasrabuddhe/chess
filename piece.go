package chess

// Piece represents a chess piece.
type Piece interface {
	// Moves returns a list of possible moves for the piece.
	Moves() []string
}

// Pawn represents a pawn.
type Pawn struct {
	position Position
}

// NewPawn creates a new Pawn with the given position.
func NewPawn(position string) (*Pawn, error) {
	pos, err := ParsePosition(position)
	if err != nil {
		return nil, err
	}
	return &Pawn{position: pos}, nil
}

// Moves returns a list of possible moves for the pawn.
func (p *Pawn) Moves() []string {
	// move one step forward
	if p.position.Row-1 >= 0 {
		newPos := Position{Row: p.position.Row - 1, Col: p.position.Col}
		return []string{newPos.String()}
	}
	return nil
}

// King represents a king.
type King struct {
	position Position
}

// NewKing creates a new King with the given position.
func NewKing(position string) (*King, error) {
	pos, err := ParsePosition(position)
	if err != nil {
		return nil, err
	}
	return &King{position: pos}, nil
}

// Moves returns a list of possible moves for the king.
func (k *King) Moves() []string {
	// a king can possibly move 1 step in 8 directions.
	var moves = make([]string, 0, 8)
	// move one step forward
	if k.position.Row-1 >= 0 {
		newPos := Position{Row: k.position.Row - 1, Col: k.position.Col}
		moves = append(moves, newPos.String())
	}
	// move one step backward
	if k.position.Row+1 <= 7 {
		newPos := Position{Row: k.position.Row + 1, Col: k.position.Col}
		moves = append(moves, newPos.String())
	}
	// move one step to the left
	if k.position.Col-1 >= 0 {
		newPos := Position{Row: k.position.Row, Col: k.position.Col - 1}
		moves = append(moves, newPos.String())
	}
	// move one step to the right
	if k.position.Col+1 <= 7 {
		newPos := Position{Row: k.position.Row, Col: k.position.Col + 1}
		moves = append(moves, newPos.String())
	}
	// move one step diagonally forward to the left
	if k.position.Row+1 <= 7 && k.position.Col-1 >= 0 {
		newPos := Position{Row: k.position.Row + 1, Col: k.position.Col - 1}
		moves = append(moves, newPos.String())
	}
	// move one step diagonally forward to the right
	if k.position.Row+1 <= 7 && k.position.Col+1 <= 7 {
		newPos := Position{Row: k.position.Row + 1, Col: k.position.Col + 1}
		moves = append(moves, newPos.String())
	}
	// move one step diagonally backward to the left
	if k.position.Row-1 >= 0 && k.position.Col-1 >= 0 {
		newPos := Position{Row: k.position.Row - 1, Col: k.position.Col - 1}
		moves = append(moves, newPos.String())
	}
	// move one step diagonally backward to the right
	if k.position.Row-1 >= 0 && k.position.Col+1 <= 7 {
		newPos := Position{Row: k.position.Row - 1, Col: k.position.Col + 1}
		moves = append(moves, newPos.String())
	}
	return moves
}

// Queen represents a queen.
type Queen struct {
	position Position
}

// NewQueen creates a new Queen with the given position.
func NewQueen(position string) (*Queen, error) {
	pos, err := ParsePosition(position)
	if err != nil {
		return nil, err
	}
	return &Queen{position: pos}, nil
}

// Moves returns a list of possible moves for the queen.
func (q *Queen) Moves() []string {
	up, down, left, right := q.position.Row, 7-q.position.Row, q.position.Col, 7-q.position.Col
	var diagFwdLeft, diagFwdRight, diagBackLeft, diagBackRight int
	for i := 1; i <= q.position.Row && i <= q.position.Col; i++ {
		diagFwdLeft++
	}
	for i := 1; i <= q.position.Row && i <= 7-q.position.Col; i++ {
		diagFwdRight++
	}
	for i := 1; i <= 7-q.position.Row && i <= q.position.Col; i++ {
		diagBackLeft++
	}
	for i := 1; i <= 7-q.position.Row && i <= 7-q.position.Col; i++ {
		diagBackRight++
	}
	// pre allocate the moves slice with the maximum possible moves
	var moves = make([]string, 0, up+down+left+right+diagFwdLeft+diagFwdRight+diagBackLeft+diagBackRight)
	// move forward across the board
	for i := 1; i <= q.position.Row; i++ {
		newPos := Position{Row: q.position.Row - i, Col: q.position.Col}
		moves = append(moves, newPos.String())
	}
	// move backward across the board
	for i := 1; i <= 7-q.position.Row; i++ {
		newPos := Position{Row: q.position.Row + i, Col: q.position.Col}
		moves = append(moves, newPos.String())
	}
	// move to the left across the board
	for i := 1; i <= q.position.Col; i++ {
		newPos := Position{Row: q.position.Row, Col: q.position.Col - i}
		moves = append(moves, newPos.String())
	}
	// move to the right across the board
	for i := 1; i <= 7-q.position.Col; i++ {
		newPos := Position{Row: q.position.Row, Col: q.position.Col + i}
		moves = append(moves, newPos.String())
	}
	// move diagonally forward to the left across the board
	for i := 1; i <= q.position.Row && i <= q.position.Col; i++ {
		newPos := Position{Row: q.position.Row - i, Col: q.position.Col - i}
		moves = append(moves, newPos.String())
	}
	// move diagonally forward to the right across the board
	for i := 1; i <= q.position.Row && i <= 7-q.position.Col; i++ {
		newPos := Position{Row: q.position.Row - i, Col: q.position.Col + i}
		moves = append(moves, newPos.String())
	}
	// move diagonally backward to the left across the board
	for i := 1; i <= 7-q.position.Row && i <= q.position.Col; i++ {
		newPos := Position{Row: q.position.Row + i, Col: q.position.Col - i}
		moves = append(moves, newPos.String())
	}
	// move diagonally backward to the right across the board
	for i := 1; i <= 7-q.position.Row && i <= 7-q.position.Col; i++ {
		newPos := Position{Row: q.position.Row + i, Col: q.position.Col + i}
		moves = append(moves, newPos.String())
	}
	return moves
}
