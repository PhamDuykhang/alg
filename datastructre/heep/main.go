/*
 * Copyright (c) 2020. Author by KTECH Inc team. khang_duy.p
 */

package main

import "fmt"

func main() {
	h := NewMaxHeep()
	//c := make([]int, 10)
	a := []int{2, 4, 1, 37, 10, 7, 8, 16, 14}

	h.BuildMaxHeep(a)

	h.PrintArr()

	fmt.Printf("max = %d\n", h.ExtractMax())

	h.PrintArr()
}
