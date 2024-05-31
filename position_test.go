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
		Row int
		Col int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "A8",
			fields: fields{Row: 0, Col: 0},
			want:   "A8",
		},
		{
			name:   "H1",
			fields: fields{Row: 7, Col: 7},
			want:   "H1",
		},
		{
			name:   "E4",
			fields: fields{Row: 4, Col: 4},
			want:   "E4",
		},
		{
			name:   "D7",
			fields: fields{Row: 1, Col: 3},
			want:   "D7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				Row: tt.fields.Row,
				Col: tt.fields.Col,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPosition_String(b *testing.B) {
	var position string
	pos := Position{Row: 5, Col: 2}

	for i := 0; i < b.N; i++ {
		position = pos.String()
	}

	_, _ = fmt.Fprintln(io.Discard, position)
}
