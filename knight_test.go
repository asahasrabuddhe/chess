package chess

import (
	"fmt"
	"io"
	"testing"
)

func TestKnight_Moves(t *testing.T) {
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
			name: "Knight D5: White",
			args: args{
				position: "D5",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"E7", "C7", "E3", "C3", "B6", "B4", "F6", "F4"},
		},
		{
			name: "Knight D5: Black",
			args: args{
				position: "D5",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"C3", "E3", "C7", "E7", "F4", "F6", "B4", "B6"},
		},
		{
			name: "Knight E8: White",
			args: args{
				position: "E8",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"F6", "D6", "C7", "G7"},
		},
		{
			name: "Knight E8: Black",
			args: args{
				position: "E8",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"D6", "F6", "G7", "C7"},
		},
		{
			name: "Knight A1: White",
			args: args{
				position: "A1",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"B3", "C2"},
		},
		{
			name: "Knight A1: Black",
			args: args{
				position: "A1",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"B3", "C2"},
		},
		{
			name: "Knight H6: White",
			args: args{
				position: "H6",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"G8", "G4", "F7", "F5"},
		},
		{
			name: "Knight H6: Black",
			args: args{
				position: "H6",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"G4", "G8", "F5", "F7"},
		},
		{
			name: "Knight C3: White",
			args: args{
				position: "C3",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"D5", "B5", "D1", "B1", "A4", "A2", "E4", "E2"},
		},
		{
			name: "Knight C3: Black",
			args: args{
				position: "C3",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"B1", "D1", "B5", "D5", "E2", "E4", "A2", "A4"},
		},
		{
			name: "Knight F4: White",
			args: args{
				position: "F4",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"G6", "E6", "G2", "E2", "D5", "D3", "H5", "H3"},
		},
		{
			name: "Knight F4: Black",
			args: args{
				position: "F4",
				color:    Black,
				board:    EmptyBoard(),
			},
			want: []string{"E2", "G2", "E6", "G6", "H3", "H5", "D3", "D5"},
		},
		{
			name: "Knight I9: White",
			args: args{
				position: "I9",
				color:    White,
				board:    EmptyBoard(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Knight I9: Black",
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
			k, err := NewKnight(tt.args.position, tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKnight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if k != nil {
				got := k.PossiblePositions(tt.args.board)
				if got != nil && tt.want != nil {
					if len(got) != len(tt.want) {
						t.Errorf("Knight.PossiblePositions() = %v, want %v", got, tt.want)
						return
					}
					for i := range got {
						if got[i].String() != tt.want[i] {
							t.Errorf("Knight.PossiblePositions() = %v, want %v", got, tt.want)
							return
						}
					}
				}
			}
		})
	}
}

func BenchmarkKnight_Moves(b *testing.B) {
	var possiblePositions []Position
	var board = EmptyBoard()

	knight, _ := NewKnight("D5", White)
	for i := 0; i < b.N; i++ {
		possiblePositions = knight.PossiblePositions(board)
	}

	_, _ = fmt.Fprintln(io.Discard, possiblePositions)
}
