package chess

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestParsePosition(t *testing.T) {
	type args struct {
		position string
	}
	tests := []struct {
		name    string
		args    args
		want    Position
		wantErr bool
	}{
		{
			name: "A8",
			args: args{position: "A8"},
			want: Position{File: 0, Rank: 0},
		},
		{
			name: "H1",
			args: args{position: "H1"},
			want: Position{Rank: 7, File: 7},
		},
		{
			name: "E4",
			args: args{position: "E4"},
			want: Position{Rank: 4, File: 4},
		},
		{
			name: "D7",
			args: args{position: "D7"},
			want: Position{Rank: 1, File: 3},
		},
		{
			name: "H1",
			args: args{position: "H1"},
			want: Position{Rank: 7, File: 7},
		},
		{
			name:    "A9",
			args:    args{position: "A9"},
			want:    Position{},
			wantErr: true,
		},
		{
			name:    "I2",
			args:    args{position: "I2"},
			want:    Position{},
			wantErr: true,
		},
		{
			name:    "AC21",
			args:    args{position: "AC21"},
			want:    Position{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePosition(tt.args.position)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_String(t *testing.T) {
	type fields struct {
		Rank int
		File rune
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "A8",
			fields: fields{Rank: 0, File: 0},
			want:   "A8",
		},
		{
			name:   "H1",
			fields: fields{Rank: 7, File: 7},
			want:   "H1",
		},
		{
			name:   "E4",
			fields: fields{Rank: 4, File: 4},
			want:   "E4",
		},
		{
			name:   "D7",
			fields: fields{Rank: 1, File: 3},
			want:   "D7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				File: tt.fields.File,
				Rank: tt.fields.Rank,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPosition_String(b *testing.B) {
	var position string
	pos := Position{File: 5, Rank: 2}

	for i := 0; i < b.N; i++ {
		position = pos.String()
	}

	_, _ = fmt.Fprintln(io.Discard, position)
}
