package sort

import (
	"reflect"
	"testing"
)

func Test_findBiggest(t *testing.T) {
	type args struct {
		haystack []int64
	}
	tests := []struct {
		name            string
		args            args
		wantVal         int64
		wantNewHaystack []int64
	}{
		// TODO: Add test cases.
		{
			name: "找最大数并移除",
			args: args{
				haystack: []int64{1, 14, 48, 7, 32, 99},
			},
			wantVal:         int64(99),
			wantNewHaystack: []int64{1, 14, 48, 7, 32},
		},
		{
			name: "找最大数并移除",
			args: args{
				haystack: []int64{1, 14, 48, 7, 32},
			},
			wantVal:         int64(48),
			wantNewHaystack: []int64{1, 14, 7, 32},
		},
		{
			name: "找最大数并移除",
			args: args{
				haystack: []int64{10002, 14, 48, 7, 32, 99},
			},
			wantVal:         int64(10002),
			wantNewHaystack: []int64{14, 48, 7, 32, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotNewHaystack := findBiggest(tt.args.haystack)
			if gotVal != tt.wantVal {
				t.Errorf("findBiggest() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if !reflect.DeepEqual(gotNewHaystack, tt.wantNewHaystack) {
				t.Errorf("findBiggest() gotNewHaystack = %v, want %v", gotNewHaystack, tt.wantNewHaystack)
			}
		})
	}
}

func TestSelectSort(t *testing.T) {
	type args struct {
		haystack []int64
	}
	tests := []struct {
		name       string
		args       args
		wantSorted []int64
	}{
		// TODO: Add test cases.
		{
			name: "排序1",
			args: args{
				haystack: []int64{2, 99, 32, 100, 43},
			},
			wantSorted: []int64{100, 99, 43, 32, 2},
		},
		{
			name: "排序2",
			args: args{
				haystack: []int64{1, 2, 3, 6, 43},
			},
			wantSorted: []int64{43, 6, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSorted := SelectSort(tt.args.haystack); !reflect.DeepEqual(gotSorted, tt.wantSorted) {
				t.Errorf(tt.name+" SelectSort() = %v, want %v", gotSorted, tt.wantSorted)
			}
		})
	}
}
