package chess

import (
	"fmt"
	"io"
	"testing"
)

func TestRook_Moves(t *testing.T) {
	type args struct {
		position string
		color    Color
		board    *Board
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Rook D5: White",
			args: args{
				position: "D5",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"D6", "D7", "D8", "D4", "D3", "D2", "D1", "C5", "B5", "A5", "E5", "F5", "G5", "H5"},
		},
		{
			name: "Rook D5: Black",
			args: args{
				position: "D5",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"D4", "D3", "D2", "D1", "D6", "D7", "D8", "E5", "F5", "G5", "H5", "C5", "B5", "A5"},
		},
		{
			name: "Rook E8: White",
			args: args{
				position: "E8",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"E7", "E6", "E5", "E4", "E3", "E2", "E1", "D8", "C8", "B8", "A8", "F8", "G8", "H8"},
		},
		{
			name: "Rook E8: Black",
			args: args{
				position: "E8",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"E7", "E6", "E5", "E4", "E3", "E2", "E1", "D8", "C8", "B8", "A8", "F8", "G8", "H8"},
		},
		{
			name: "Rook A1: White",
			args: args{
				position: "A1",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"A2", "A3", "A4", "A5", "A6", "A7", "A8", "B1", "C1", "D1", "E1", "F1", "G1", "H1"},
		},
		{
			name: "Rook A1: Black",
			args: args{
				position: "A1",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"A2", "A3", "A4", "A5", "A6", "A7", "A8", "B1", "C1", "D1", "E1", "F1", "G1", "H1"},
		},
		{
			name: "Rook H6: White",
			args: args{
				position: "H6",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"H7", "H8", "H5", "H4", "H3", "H2", "H1", "G6", "F6", "E6", "D6", "C6", "B6", "A6"},
		},
		{
			name: "Rook H6: Black",
			args: args{
				position: "H6",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"H5", "H4", "H3", "H2", "H1", "H7", "H8", "G6", "F6", "E6", "D6", "C6", "B6", "A6"},
		},
		{
			name: "Rook C3: White",
			args: args{
				position: "C3",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"C4", "C5", "C6", "C7", "C8", "C2", "C1", "B3", "A3", "D3", "E3", "F3", "G3", "H3"},
		},
		{
			name: "Rook C3: Black",
			args: args{
				position: "C3",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"C2", "C1", "C4", "C5", "C6", "C7", "C8", "D3", "E3", "F3", "G3", "H3", "B3", "A3"},
		},
		{
			name: "Rook F4: White",
			args: args{
				position: "F4",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"F5", "F6", "F7", "F8", "F3", "F2", "F1", "E4", "D4", "C4", "B4", "A4", "G4", "H4"},
		},
		{
			name: "Rook F4: Black",
			args: args{
				position: "F4",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"F3", "F2", "F1", "F5", "F6", "F7", "F8", "G4", "H4", "E4", "D4", "C4", "B4", "A4"},
		},
		{
			name: "Rook I9: White",
			args: args{
				position: "I9",
				color:    White,
				board:    EmptyBoard(),
			},
			wantErr: true,
		},
		{
			name: "Rook I9: Black",
			args: args{
				position: "I9",
				color:    Black,
				board:    EmptyBoard(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewRook(tt.args.position, tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if r != nil {
				got := r.PossiblePositions(tt.args.board)
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

func BenchmarkRook_Moves(b *testing.B) {
	var possiblePositions []Position
	var board = EmptyBoard()

	for i := 0; i < b.N; i++ {
		r, _ := NewRook("D5", White)
		r.PossiblePositions(board)
	}

	_, _ = fmt.Fprintln(io.Discard, possiblePositions)
}
