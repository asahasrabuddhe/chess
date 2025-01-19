package chess

import (
	"fmt"
	"io"
	"testing"
)

func TestBishop_Moves(t *testing.T) {
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
			name: "Bishop D5: White",
			args: args{
				position: "D5",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"C6", "B7", "A8", "E6", "F7", "G8", "C4", "B3", "A2", "E4", "F3", "G2", "H1"},
		},
		{
			name: "Bishop D5: Black",
			args: args{
				position: "D5",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"E4", "F3", "G2", "H1", "C4", "B3", "A2", "E6", "F7", "G8", "C6", "B7", "A8"},
		},
		{
			name: "Bishop E8: White",
			args: args{
				position: "E8",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"D7", "C6", "B5", "A4", "F7", "G6", "H5"},
		},
		{
			name: "Bishop E8: Black",
			args: args{
				position: "E8",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"F7", "G6", "H5", "D7", "C6", "B5", "A4"},
		},
		{
			name: "Bishop A1: White",
			args: args{
				position: "A1",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"B2", "C3", "D4", "E5", "F6", "G7", "H8"},
		},
		{
			name: "Bishop A1: Black",
			args: args{
				position: "A1",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"B2", "C3", "D4", "E5", "F6", "G7", "H8"},
		},
		{
			name: "Bishop H6: White",
			args: args{
				position: "H6",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"G7", "F8", "G5", "F4", "E3", "D2", "C1"},
		},
		{
			name: "Bishop H6: Black",
			args: args{
				position: "H6",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"G5", "F4", "E3", "D2", "C1", "G7", "F8"},
		},
		{
			name: "Bishop C3: White",
			args: args{
				position: "C3",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"B4", "A5", "D4", "E5", "F6", "G7", "H8", "B2", "A1", "D2", "E1"},
		},
		{
			name: "Bishop C3: Black",
			args: args{
				position: "C3",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"D2", "E1", "B2", "A1", "D4", "E5", "F6", "G7", "H8", "B4", "A5"},
		},
		{
			name: "Bishop F4: White",
			args: args{
				position: "F4",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"E5", "D6", "C7", "B8", "G5", "H6", "E3", "D2", "C1", "G3", "H2"},
		},
		{
			name: "Bishop F4: Black",
			args: args{
				position: "F4",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"G3", "H2", "E3", "D2", "C1", "G5", "H6", "E5", "D6", "C7", "B8"},
		},
		{
			name: "Bishop I9: White",
			args: args{
				position: "I9",
				color:    White,
				board:    EmptyBoard(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Bishop I9: Black",
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
			b, err := NewBishop(tt.args.position, tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBishop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if b != nil {
				got := b.PossiblePositions(tt.args.board)
				if got != nil && tt.want != nil {
					if len(got) != len(tt.want) {
						t.Errorf("Bishop.PossiblePositions() = %v, want %v", got, tt.want)
						return
					}
					for i := range got {
						if got[i].String() != tt.want[i] {
							t.Errorf("Bishop.PossiblePositions() = %v, want %v", got, tt.want)
							return
						}
					}
				}
			}
		})
	}
}

func BenchmarkBishop_Moves(b *testing.B) {
	var possiblePositions []Position
	var board = EmptyBoard()

	bishop, _ := NewBishop("D5", White)
	for i := 0; i < b.N; i++ {
		possiblePositions = bishop.PossiblePositions(board)
	}

	_, _ = fmt.Fprintln(io.Discard, possiblePositions)
}
