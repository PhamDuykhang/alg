/*
 * Copyright (c) 2020. Author by KTECH Inc team. khang_duy.p
 */

package main

import "fmt"

const NegInfinity = -99999

func main() {
	a := []int{-6, 3, 4, 5, 7}

	MaximumSubArr(a)
}

func MaximumSubArr(a []int) {
	l, r, s := maxMaxSub(a, 0, len(a)-1)
	fmt.Print("[")
	for i := l; i <= r; i++ {
		fmt.Printf("%d", a[i])
		if i != r {
			fmt.Print(",")
		}
	}
	fmt.Printf("] total sum = %d \n", s)
}

func maxCrossPoint(a []int, low, mid, high int) (maxL, maxR, sum int) {
	leftSum := NegInfinity
	s := 0
	for i := mid; i >= low; i-- {
		s = s + a[i]
		if s > leftSum {
			leftSum = s
			maxL = i
		}
	}
	rightSum := NegInfinity
	s = 0
	for i := mid + 1; i <= high; i++ {
		s = s + a[i]
		if s > rightSum {
			rightSum = s
			maxR = i
		}
	}
	sum = leftSum + rightSum
	return
}

func maxMaxSub(a []int, low, high int) (idxL, idxR, sum int) {
	if low == high {
		return low, high, a[low]
	}
	mid := (low + high) / 2
	lefLow, leftHigh, leftSum := maxMaxSub(a, low, mid)
	rightLow, rightHigh, rightSum := maxMaxSub(a, mid+1, high)
	crossLow, crossHigh, crossSum := maxCrossPoint(a, low, mid, high)

	if leftSum >= rightSum && leftSum >= crossSum {
		return lefLow, leftHigh, leftSum
	}
	if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	}
	return crossLow, crossHigh, crossSum
}
