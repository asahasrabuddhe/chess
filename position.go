package chess

type Position struct {
	Row int
	Col int
}

func ParsePosition(position string) Position {
	return Position{}
}

func (p Position) String() string {
	return ""
}
