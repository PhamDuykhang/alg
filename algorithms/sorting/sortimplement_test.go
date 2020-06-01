package sorting

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	a := []int{4, 5, 6, 7}
	swap(&a[1], &a[2])
	fmt.Println(a)
}

func TestSelectionSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Normal",
			args: args{a: []int{3, 6, 4, 1, 6, 8}},
		},
		{
			name: "Nil",
			args: args{a: []int{}},
		},
		{
			name: "Same",
			args: args{a: []int{1,1,1,1,1,1,1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.a)
			fmt.Println("result",tt.args.a)
		})
	}
}

func TestBubbleSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Normal",
			args: args{a: []int{3, 6, 4, 1, 6, 8}},
		},
		{
			name: "Nil",
			args: args{a: []int{}},
		},
		{
			name: "Same",
			args: args{a: []int{1,1,1,1,1,1,1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.a)
			fmt.Println("result",tt.args.a)
		})
	}
}

func TestInsertSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Normal",
			args: args{a: []int{3, 6, 4, 1, 6, 8}},
		},
		{
			name: "Nil",
			args: args{a: []int{}},
		},
		{
			name: "Same",
			args: args{a: []int{1,1,1,1,1,1,1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("age   ",tt.args.a)
			InsertSort(tt.args.a)
			fmt.Println("result",tt.args.a)
		})
	}
}

func Test_merge(t *testing.T) {
	a := []int{1,4,3,8,2,10,5,7,9}

	merge(a,2,3,5)

}
func TestMergeSort(t *testing.T) {
	a:= []int{32, 45, 67, 2, 7}
	fmt.Println(a)
	MergeSort(a)
	fmt.Println(a)
}