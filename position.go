package chess

import "fmt"

type Position struct {
	Row int
	Col int
}

func ParsePosition(position string) (Position, error) {
	if len(position) != 2 {
		return Position{}, fmt.Errorf("invalid position: %s", position)
	}
	if position[0] < 'A' || position[0] > 'H' {
		return Position{}, fmt.Errorf("invalid position: %s", position)
	}
	if position[1] < '1' || position[1] > '8' {
		return Position{}, fmt.Errorf("invalid position: %s", position)
	}
	return Position{
		Row: int('8' - position[1]),
		Col: int(position[0] - 'A'),
	}, nil
}

func (p Position) String() string {
	return string(rune('A'+p.Col)) + fmt.Sprint(8-p.Row)
}
