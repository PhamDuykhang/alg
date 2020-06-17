/*
 * Copyright (c) 2020. Author by KTECH Inc team. khang_duy.p
 */

package main

import "fmt"

var Price []int = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

func RodCutting(nL int) (maxR, s []int) {
	//init memories table
	var m = make([]int, len(Price))
	var n = make([]int, len(Price))
	m[0] = 0
	for i := 1; i <= nL; i++ {
		q := -9999
		for j := 1; j <= i; j++ {
			if q < Price[j]+m[i-j] {
				q = Price[j] + m[i-j]
				n[i] = j
			}
		}
		m[i] = q
	}
	return m, n
}

func main() {
	n := 7
	a, b := RodCutting(n)

	fmt.Println(a)
	fmt.Println(b)

	for n > 0 {
		fmt.Printf("%d ", b[n])
		n = n - b[n]
	}
	fmt.Println()
	fmt.Printf("max %d", a[n])
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
