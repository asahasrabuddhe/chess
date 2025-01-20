package chess

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestCompoundMove(t *testing.T) {
	type args struct {
		moves []MoveFuncs
	}
	tests := []struct {
		name string
		args args
		want MoveFuncs
	}{
		{
			name: "CompoundMove 1",
			args: args{
				moves: []MoveFuncs{
					MoveForward(2),
					MoveRight(1),
				},
			},
			want: MoveFuncs{
				Compound: true,
				Moves:    append(MoveForward(2).Moves, MoveRight(1).Moves...),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompoundMove(tt.args.moves...)
			if got.Compound != tt.want.Compound {
				t.Errorf("CompoundMove() got = %v, want = %v", got, tt.want)
			}
			if len(got.Moves) != len(tt.want.Moves) {
				t.Errorf("CompoundMove() got = %v, want = %v", got, tt.want)
			}
			for i, g := range got.Moves {
				if reflect.TypeOf(g) != reflect.TypeOf(tt.want.Moves[i]) {
					t.Errorf("CompoundMove() got = %v, want = %v", got, tt.want)
				}
				if reflect.ValueOf(g).Pointer() != reflect.ValueOf(tt.want.Moves[i]).Pointer() {
					t.Errorf("CompoundMove() got = %v, want = %v", got, tt.want)
				}
			}
		})
	}
}

func BenchmarkCompoundMove(b *testing.B) {
	var moves MoveFuncs
	for i := 0; i < b.N; i++ {
		moves = CompoundMove(MoveForward(2), MoveRight(1))
	}
	_, _ = fmt.Fprintln(io.Discard, moves)
}
