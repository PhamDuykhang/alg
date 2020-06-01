package main

import "fmt"

func BinarySearch(a []int, i int) (index int, err error) {
	index, err = binaryUtil(a, 0, len(a)-1, i)
	if err != nil {
		return
	}
	return
}

func binaryUtil(a []int, h, l, i int) (int, error) {
	if h > l {
		return -1, fmt.Errorf("not found")
	}
	mid := h + (l-h)/2
	if a[mid] == i {
		return mid, nil
	}
	if a[mid] > i {
		return binaryUtil(a, mid+1, l, i)
	}
	return binaryUtil(a, h, mid-1, i)

}
