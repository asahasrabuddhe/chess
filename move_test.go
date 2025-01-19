package chess

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestCompoundMove(t *testing.T) {
	type args struct {
		moves [][]MoveFunc
	}
	tests := []struct {
		name string
		args args
		want []MoveFunc
	}{
		{
			name: "CompoundMove 1",
			args: args{
				moves: [][]MoveFunc{
					MoveForward(2),
					MoveRight(1),
				},
			},
			want: append(MoveForward(2), MoveRight(1)...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompoundMove(tt.args.moves...)
			if len(got) != len(tt.want) {
				t.Errorf("CompoundMove() got = %v, want = %v", got, tt.want)
			}
			for i, g := range got {
				if reflect.TypeOf(g) != reflect.TypeOf(tt.want[i]) {
					t.Errorf("CompoundMove() got = %v, want = %v", got, tt.want)
				}
				if reflect.ValueOf(g).Pointer() != reflect.ValueOf(tt.want[i]).Pointer() {
					t.Errorf("CompoundMove() got = %v, want = %v", got, tt.want)
				}
			}
		})
	}
}

func BenchmarkCompoundMove(b *testing.B) {
	var moves []MoveFunc
	for i := 0; i < b.N; i++ {
		moves = CompoundMove(MoveForward(2), MoveRight(1))
	}
	_, _ = fmt.Fprintln(io.Discard, moves)
}
