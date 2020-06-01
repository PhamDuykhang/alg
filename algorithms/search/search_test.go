package main

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		a []int
		i int
	}
	tests := []struct {
		name      string
		args      args
		wantIndex int
		wantErr   bool
	}{
		{
			name:"Normal case",
			args: args{
				a: []int{1,3,4,5,7,8},
				i: 4,
			},
			wantErr: false,
			wantIndex: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, err := BinarySearch(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinarySearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("BinarySearch() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}
