package chess

type Piece interface {
	Moves() []string
}

type Pawn struct {
	position Position
}

func NewPawn(position string) *Pawn {
	return &Pawn{position: ParsePosition(position)}
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

func NewKing(position string) *King {
	return &King{position: ParsePosition(position)}
}

func (k *King) Moves() []string {
	var moves []string
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

