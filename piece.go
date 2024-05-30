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

