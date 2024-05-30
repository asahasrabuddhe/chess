package chess

import (
	"reflect"
	"testing"
)

func TestParsePosition(t *testing.T) {
	type args struct {
		position string
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{
			name: "A8",
			args: args{position: "A8"},
			want: Position{Row: 0, Col: 0},
		},
		{
			name: "H1",
			args: args{position: "H1"},
			want: Position{Row: 7, Col: 7},
		},
		{
			name: "E4",
			args: args{position: "E4"},
			want: Position{Row: 4, Col: 4},
		},
		{
			name: "D7",
			args: args{position: "D7"},
			want: Position{Row: 1, Col: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParsePosition(tt.args.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
