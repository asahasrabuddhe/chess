package chess

import (
	"fmt"
	"io"
	"testing"
)

func TestKing_Moves(t *testing.T) {
	type args struct {
		position string
		color    Color
		board    Board
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "King D5: White",
			args: args{
				position: "D5",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"D6", "D4", "C5", "E5", "C6", "E6", "C4", "E4"},
		},
		{
			name: "King D5: Black",
			args: args{
				position: "D5",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"D4", "D6", "E5", "C5", "E4", "C4", "E6", "C6"},
		},
		{
			name: "King E8: White",
			args: args{
				position: "E8",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"E7", "D8", "F8", "D7", "F7"},
		},
		{
			name: "King E8: Black",
			args: args{
				position: "E8",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"E7", "F8", "D8", "F7", "D7"},
		},
		{
			name: "King A1: White",
			args: args{
				position: "A1",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"A2", "B1", "B2"},
		},
		{
			name: "King A1: Black",
			args: args{
				position: "A1",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"A2", "B1", "B2"},
		},
		{
			name: "King H6: White",
			args: args{
				position: "H6",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"H7", "H5", "G6", "G7", "G5"},
		},
		{
			name: "King H6: Black",
			args: args{
				position: "H6",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"H5", "H7", "G6", "G5", "G7"},
		},
		{
			name: "King C3: White",
			args: args{
				position: "C3",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"C4", "C2", "B3", "D3", "B4", "D4", "B2", "D2"},
		},
		{
			name: "King C3: Black",
			args: args{
				position: "C3",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"C2", "C4", "D3", "B3", "D2", "B2", "D4", "B4"},
		},
		{
			name: "King F4: White",
			args: args{
				position: "F4",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"F5", "F3", "E4", "G4", "E5", "G5", "E3", "G3"},
		},
		{
			name: "King F4: Black",
			args: args{
				position: "F4",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"F3", "F5", "G4", "E4", "G3", "E3", "G5", "E5"},
		},
		{
			name: "King I9: White",
			args: args{
				position: "I9",
				color:    White,
				board:    EmptyBoard(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "King I9: Black",
			args: args{
				position: "I9",
				color:    Black,
				board:    EmptyBoard(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k, err := NewKing(tt.args.position, tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if k != nil {
				got := k.PossiblePositions(tt.args.board)
				if got != nil && tt.want != nil {
					if len(got) != len(tt.want) {
						t.Errorf("PossiblePositions() got=%v, want = %v", got, tt.want)
						return
					}
					for i, g := range got {
						if g.String() != tt.want[i] {
							t.Errorf("PossiblePositions() got=%v, want = %v", got, tt.want)
							return
						}
					}
				}
			}
		})
	}
}

func BenchmarkKing_Moves(b *testing.B) {
	var possiblePositions []Position
	var board = EmptyBoard()

	king, _ := NewKing("D5", White)
	for i := 0; i < b.N; i++ {
		possiblePositions = king.PossiblePositions(board)
	}

	_, _ = fmt.Fprintln(io.Discard, possiblePositions)
}
