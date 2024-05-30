package chess

type Position struct {
	Row int
	Col int
}

func ParsePosition(position string) Position {
	return Position{
		Row: int('8' - position[1]),
		Col: int(position[0] - 'A'),
	}
}

func (p Position) String() string {
	return ""
}
