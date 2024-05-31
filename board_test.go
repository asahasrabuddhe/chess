package chess

import (
	"testing"
)

func TestEmptyBoard(t *testing.T) {
	tests := []struct {
		name string
		want *Board
	}{
		{
			name: "Empty Board",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EmptyBoard()
			for _, row := range got.pieces {
				for _, pc := range row {
					if pc != nil {
						t.Errorf("EmptyBoard() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}
