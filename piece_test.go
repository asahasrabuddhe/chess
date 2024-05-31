package chess

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestPawn_Moves(t *testing.T) {
	type fields struct {
		position string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "Pawn D5",
			fields: fields{
				position: "D5",
			},
			want: []string{"D6"},
		},
		{
			name: "Pawn E8",
			fields: fields{
				position: "E8",
			},
			want: nil,
		},
		{
			name: "Pawn A1",
			fields: fields{
				position: "A1",
			},
			want: []string{"A2"},
		},
		{
			name: "Pawn H6",
			fields: fields{
				position: "H6",
			},
			want: []string{"H7"},
		},
		{
			name: "Pawn C3",
			fields: fields{
				position: "C3",
			},
			want: []string{"C4"},
		},
		{
			name: "Pawn F4",
			fields: fields{
				position: "F4",
			},
			want: []string{"F5"},
		},
		{
			name: "Pawn I9",
			fields: fields{
				position: "I9",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewPawn(tt.fields.position)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPawn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if p != nil {
				if got := p.Moves(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Moves() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func BenchmarkPawn_Moves(b *testing.B) {
	var moves []string

	pawn, _ := NewPawn("D5")
	for i := 0; i <= b.N; i++ {
		moves = pawn.Moves()
	}

	_, _ = fmt.Fprintln(io.Discard, moves)
}

func TestKing_Moves(t *testing.T) {
	type fields struct {
		position string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "King D5",
			fields: fields{
				position: "D5",
			},
			want: []string{"D6", "D4", "C5", "E5", "C4", "E4", "C6", "E6"},
		},
		{
			name: "King E8",
			fields: fields{
				position: "E8",
			},
			want: []string{"E7", "D8", "F8", "D7", "F7"},
		},
		{
			name: "King A1",
			fields: fields{
				position: "A1",
			},
			want: []string{"A2", "B1", "B2"},
		},
		{
			name: "King H6",
			fields: fields{
				position: "H6",
			},
			want: []string{"H7", "H5", "G6", "G5", "G7"},
		},
		{
			name: "King C3",
			fields: fields{
				position: "C3",
			},
			want: []string{"C4", "C2", "B3", "D3", "B2", "D2", "B4", "D4"},
		},
		{
			name: "King F4",
			fields: fields{
				position: "F4",
			},
			want: []string{"F5", "F3", "E4", "G4", "E3", "G3", "E5", "G5"},
		},
		{
			name: "King I9",
			fields: fields{
				position: "I9",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k, err := NewKing(tt.fields.position)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if k != nil {
				if got := k.Moves(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Moves() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func BenchmarkKing_Moves(b *testing.B) {
	var moves []string

	king, _ := NewKing("D5")
	for i := 0; i < b.N; i++ {
		moves = king.Moves()
	}

	_, _ = fmt.Fprintln(io.Discard, moves)
}

func TestQueen_Moves(t *testing.T) {
	type fields struct {
		position string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "Queen D5",
			fields: fields{
				position: "D5",
			},
			want: []string{"D6", "D7", "D8", "D4", "D3", "D2", "D1", "C5", "B5", "A5", "E5", "F5", "G5", "H5", "C6", "B7", "A8", "E6", "F7", "G8", "C4", "B3", "A2", "E4", "F3", "G2", "H1"},
		},
		{
			name: "Queen E8",
			fields: fields{
				position: "E8",
			},
			want: []string{"E7", "E6", "E5", "E4", "E3", "E2", "E1", "D8", "C8", "B8", "A8", "F8", "G8", "H8", "D7", "C6", "B5", "A4", "F7", "G6", "H5"},
		},
		{
			name: "Queen A1",
			fields: fields{
				position: "A1",
			},
			want: []string{"A2", "A3", "A4", "A5", "A6", "A7", "A8", "B1", "C1", "D1", "E1", "F1", "G1", "H1", "B2", "C3", "D4", "E5", "F6", "G7", "H8"},
		},
		{
			name: "Queen H6",
			fields: fields{
				position: "H6",
			},
			want: []string{"H7", "H8", "H5", "H4", "H3", "H2", "H1", "G6", "F6", "E6", "D6", "C6", "B6", "A6", "G7", "F8", "G5", "F4", "E3", "D2", "C1"},
		},
		{
			name: "Queen C3",
			fields: fields{
				position: "C3",
			},
			want: []string{"C4", "C5", "C6", "C7", "C8", "C2", "C1", "B3", "A3", "D3", "E3", "F3", "G3", "H3", "B4", "A5", "D4", "E5", "F6", "G7", "H8", "B2", "A1", "D2", "E1"},
		},
		{
			name: "Queen F4",
			fields: fields{
				position: "F4",
			},
			want: []string{"F5", "F6", "F7", "F8", "F3", "F2", "F1", "E4", "D4", "C4", "B4", "A4", "G4", "H4", "E5", "D6", "C7", "B8", "G5", "H6", "E3", "D2", "C1", "G3", "H2"},
		},
		{
			name: "Queen I9",
			fields: fields{
				position: "I9",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q, err := NewQueen(tt.fields.position)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQueen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if q != nil {
				if got := q.Moves(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Moves() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
