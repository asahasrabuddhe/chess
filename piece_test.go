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

func TestKing_Moves(t *testing.T) {
	type fields struct {
		position Position
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "King D5",
			fields: fields{
				position: ParsePosition("D5"),
			},
			want: []string{"D6", "D4", "C5", "E5", "C4", "E4", "C6", "E6"},
		},
		{
			name: "King E8",
			fields: fields{
				position: ParsePosition("E8"),
			},
			want: []string{"E7", "D8", "F8", "D7", "F7"},
		},
		{
			name: "King A1",
			fields: fields{
				position: ParsePosition("A1"),
			},
			want: []string{"A2", "B1", "B2"},
		},
		{
			name: "King H6",
			fields: fields{
				position: ParsePosition("H6"),
			},
			want: []string{"H7", "H5", "G6", "G5", "G7"},
		},
		{
			name: "King C3",
			fields: fields{
				position: ParsePosition("C3"),
			},
			want: []string{"C4", "C2", "B3", "D3", "B2", "D2", "B4", "D4"},
		},
		{
			name: "King F4",
			fields: fields{
				position: ParsePosition("F4"),
			},
			want: []string{"F5", "F3", "E4", "G4", "E3", "G3", "E5", "G5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := NewKing(tt.fields.position.String())
			if got := k.Moves(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Moves() = %v, want %v", got, tt.want)
			}
		})
	}
}
