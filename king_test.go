package chess

import (
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
			name: "King D5",
			args: args{
				position: "D5",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"D6", "D4", "C5", "E5", "C6", "E6", "C4", "E4"},
		},
		{
			name: "King E8",
			args: args{
				position: "E8",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"E7", "D8", "F8", "D7", "F7"},
		},
		{
			name: "King A1",
			args: args{
				position: "A1",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"B1", "A2", "B2"},
		},
		{
			name: "King H6",
			args: args{
				position: "H6",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"H7", "H5", "G6", "G7", "G5"},
		},
		{
			name: "King C3",
			args: args{
				position: "C3",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"C4", "C2", "B3", "D3", "B4", "D4", "B2", "D2"},
		},
		{
			name: "King F4",
			args: args{
				position: "F4",
				color:    White,
				board:    EmptyBoard(),
			},
			want: []string{"F5", "F3", "E4", "G4", "E5", "G5", "E3", "G3"},
		},
		{
			name: "King I9",
			args: args{
				position: "I9",
				color:    White,
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

//func BenchmarkKing_Moves(b *testing.B) {
//	var moves []string
//
//	king, _ := NewKing("D5")
//	for i := 0; i < b.N; i++ {
//		moves = king.PossiblePositions()
//	}
//
//	_, _ = fmt.Fprintln(io.Discard, moves)
//}
