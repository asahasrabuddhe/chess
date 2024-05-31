package chess

//func TestQueen_Moves(t *testing.T) {
//	type args struct {
//		position string
//	}
//	tests := []struct {
//		name    string
//		args  args
//		want    []string
//		wantErr bool
//	}{
//		{
//			name: "Queen D5",
//			args: args{
//				position: "D5",
//			},
//			want: []string{"D6", "D7", "D8", "D4", "D3", "D2", "D1", "C5", "B5", "A5", "E5", "F5", "G5", "H5", "C6", "B7", "A8", "E6", "F7", "G8", "C4", "B3", "A2", "E4", "F3", "G2", "H1"},
//		},
//		{
//			name: "Queen E8",
//			args: args{
//				position: "E8",
//			},
//			want: []string{"E7", "E6", "E5", "E4", "E3", "E2", "E1", "D8", "C8", "B8", "A8", "F8", "G8", "H8", "D7", "C6", "B5", "A4", "F7", "G6", "H5"},
//		},
//		{
//			name: "Queen A1",
//			args: args{
//				position: "A1",
//			},
//			want: []string{"A2", "A3", "A4", "A5", "A6", "A7", "A8", "B1", "C1", "D1", "E1", "F1", "G1", "H1", "B2", "C3", "D4", "E5", "F6", "G7", "H8"},
//		},
//		{
//			name: "Queen H6",
//			args: args{
//				position: "H6",
//			},
//			want: []string{"H7", "H8", "H5", "H4", "H3", "H2", "H1", "G6", "F6", "E6", "D6", "C6", "B6", "A6", "G7", "F8", "G5", "F4", "E3", "D2", "C1"},
//		},
//		{
//			name: "Queen C3",
//			args: args{
//				position: "C3",
//			},
//			want: []string{"C4", "C5", "C6", "C7", "C8", "C2", "C1", "B3", "A3", "D3", "E3", "F3", "G3", "H3", "B4", "A5", "D4", "E5", "F6", "G7", "H8", "B2", "A1", "D2", "E1"},
//		},
//		{
//			name: "Queen F4",
//			args: args{
//				position: "F4",
//			},
//			want: []string{"F5", "F6", "F7", "F8", "F3", "F2", "F1", "E4", "D4", "C4", "B4", "A4", "G4", "H4", "E5", "D6", "C7", "B8", "G5", "H6", "E3", "D2", "C1", "G3", "H2"},
//		},
//		{
//			name: "Queen I9",
//			args: args{
//				position: "I9",
//			},
//			want:    nil,
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			q, err := NewQueen(tt.args.position)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("NewQueen() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if q != nil {
//				if got := q.PossiblePositions(); !reflect.DeepEqual(got, tt.want) {
//					t.Errorf("PossiblePositions() = %v, want %v", got, tt.want)
//				}
//			}
//		})
//	}
//}
//
//func BenchmarkQueen_Moves(b *testing.B) {
//	var moves []string
//
//	queen, _ := NewQueen("D5")
//	for i := 0; i < b.N; i++ {
//		moves = queen.PossiblePositions()
//	}
//
//	_, _ = fmt.Fprintln(io.Discard, moves)
//}
