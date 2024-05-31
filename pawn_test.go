package chess

import (
	"fmt"
	"io"
	"testing"
)

func TestPawn_Moves(t *testing.T) {
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
			name: "Pawn D5: White",
			args: args{
				position: "D5",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"D6"},
		},
		{
			name: "Pawn D5: Black",
			args: args{
				position: "D5",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"D4"},
		},
		{
			name: "Pawn E8: White",
			args: args{
				position: "E8",
				color:    White,
				board:    EmptyBoard(),
			},
			want: nil,
		},
		{
			name: "Pawn E8: Black",
			args: args{
				position: "E8",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: nil,
		},
		{
			name: "Pawn A1: White",
			args: args{
				position: "A1",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"A2"},
		},
		{
			name: "Pawn A1: Black",
			args: args{
				position: "A1",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: nil,
		},
		{
			name: "Pawn H6: White",
			args: args{
				position: "H6",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"H7"},
		},
		{
			name: "Pawn H6: Black",
			args: args{
				position: "H6",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"H5"},
		},
		{
			name: "Pawn C3: White",
			args: args{
				position: "C3",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"C4"},
		},
		{
			name: "Pawn C3: Black",
			args: args{
				position: "C3",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"C2"},
		},
		{
			name: "Pawn F4: White",
			args: args{
				position: "F4",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"F5"},
		},
		{
			name: "Pawn F4: Black",
			args: args{
				position: "F4",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"F3"},
		},
		{
			name: "Pawn I9: White",
			args: args{
				position: "I9",
				color:    White,
				board:    EmptyBoard(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Pawn I9: Black",
			args: args{
				position: "I9",
				color:    Black,
				board:    EmptyBoard(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Pawn D2: White",
			args: args{
				position: "D2",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"D3", "D4"},
		},
		{
			name: "Pawn D2: Black",
			args: args{
				position: "D2",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"D1"},
		},
		{
			name: "Pawn D7: White",
			args: args{
				position: "D7",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"D8"},
		},
		{
			name: "Pawn D7: Black",
			args: args{
				position: "D7",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"D6", "D5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewPawn(tt.args.position, tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPawn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if p != nil {
				got := p.PossiblePositions(tt.args.board)
				if got != nil && tt.want != nil {
					if len(got) != len(tt.want) {
						t.Errorf("PossiblePositions() got=%v, want = %v", got, tt.want)
					}
					for i, g := range got {
						if g.String() != tt.want[i] {
							t.Errorf("PossiblePositions() got=%v, want = %v", g.String(), tt.want[i])
						}
					}
				}
			}
		})
	}
}

func BenchmarkPawn_Moves(b *testing.B) {
	var possiblePositions []Position
	var board = EmptyBoard()

	pawn, _ := NewPawn("D5", White)
	for i := 0; i <= b.N; i++ {
		possiblePositions = pawn.PossiblePositions(board)
	}

	_, _ = fmt.Fprintln(io.Discard, possiblePositions)
}
