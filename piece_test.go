package chess

import (
	"reflect"
	"testing"
)

func TestPawn_Moves(t *testing.T) {
	type fields struct {
		position Position
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Pawn D5",
			fields: fields{
				position: ParsePosition("D5"),
			},
			want: []string{"D6"},
		},
		{
			name: "Pawn E8",
			fields: fields{
				position: ParsePosition("E8"),
			},
			want: nil,
		},
		{
			name: "Pawn A1",
			fields: fields{
				position: ParsePosition("A1"),
			},
			want: []string{"A2"},
		},
		{
			name: "Pawn H6",
			fields: fields{
				position: ParsePosition("H6"),
			},
			want: []string{"H7"},
		},
		{
			name: "Pawn C3",
			fields: fields{
				position: ParsePosition("C3"),
			},
			want: []string{"C4"},
		},
		{
			name: "Pawn F4",
			fields: fields{
				position: ParsePosition("F4"),
			},
			want: []string{"F5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPawn(tt.fields.position.String())
			if got := p.Moves(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Moves() = %v, want %v", got, tt.want)
			}
		})
	}
}

