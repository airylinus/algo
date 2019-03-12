package basic

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		haystack []int64
		niddle   int64
	}
	tests := []struct {
		name    string
		args    args
		wantIdx int
	}{
		// TODO: Add test cases.
		{
			name: "targetAtLeftBorder",
			args: args{
				haystack: []int64{23, 35},
				niddle:   int64(23),
			},
			wantIdx: 0,
		},
		{
			name: "targetAtRightBorder",
			args: args{
				haystack: []int64{23, 35},
				niddle:   int64(23),
			},
			wantIdx: 0,
		},
		{
			name: "targetAtMiddle",
			args: args{
				haystack: []int64{23, 35, 97},
				niddle:   int64(35),
			},
			wantIdx: 1,
		},
		{
			name: "targetMissing",
			args: args{
				haystack: []int64{23, 35, 37, 57},
				niddle:   int64(99),
			},
			wantIdx: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIdx := BinarySearch(tt.args.haystack, tt.args.niddle); gotIdx != tt.wantIdx {
				t.Errorf(tt.name+" BinarySearch() = %v, want %v", gotIdx, tt.wantIdx)
			}
		})
	}
}
