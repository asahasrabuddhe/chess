package chess

type Piece interface {
	Moves() []string
}

type Pawn struct {
	position Position
}

func NewPawn(position string) (*Pawn, error) {
	pos, err := ParsePosition(position)
	if err != nil {
		return nil, err
	}
	return &Pawn{position: pos}, nil
}

func (p *Pawn) Moves() []string {
	var moves []string
	// move one step forward
	if p.position.Row-1 >= 0 {
		newPos := Position{Row: p.position.Row - 1, Col: p.position.Col}
		moves = append(moves, newPos.String())
	}
	return moves
}

type King struct {
	position Position
}

func NewKing(position string) (*King, error) {
	pos, err := ParsePosition(position)
	if err != nil {
		return nil, err
	}
	return &King{position: pos}, nil
}

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

type Queen struct {
	position Position
}

func NewQueen(position string) (*Queen, error) {
	pos, err := ParsePosition(position)
	if err != nil {
		return nil, err
	}
	return &Queen{position: pos}, nil
}

func (q *Queen) Moves() []string {
	var moves []string
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
	// move diagonally forward to the left
	for i := 1; i <= q.position.Row && i <= q.position.Col; i++ {
		//if q.position.Row+1 <= 7 && q.position.Col-1 >= 0 {
		newPos := Position{Row: q.position.Row - i, Col: q.position.Col - i}
		moves = append(moves, newPos.String())
		//}
	}
	// move diagonally forward to the right
	for i := 1; i <= q.position.Row && i <= 7-q.position.Col; i++ {
		//if q.position.Row+1 <= 7 && q.position.Col+1 <= 7 {
		newPos := Position{Row: q.position.Row - i, Col: q.position.Col + i}
		moves = append(moves, newPos.String())
		//}
	}
	// move diagonally backward to the left
	for i := 1; i <= 7-q.position.Row && i <= q.position.Col; i++ {
		//if q.position.Row-1 >= 0 && q.position.Col-1 >= 0 {
		newPos := Position{Row: q.position.Row + i, Col: q.position.Col - i}
		moves = append(moves, newPos.String())
		//}
	}
	// move diagonally backward to the right
	for i := 1; i <= 7-q.position.Row && i <= 7-q.position.Col; i++ {
		//if q.position.Row-1 >= 0 && q.position.Col+1 <= 7 {
		newPos := Position{Row: q.position.Row + i, Col: q.position.Col + i}
		moves = append(moves, newPos.String())
		//}
	}
	return moves
}
