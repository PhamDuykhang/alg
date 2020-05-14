package main

import (
	"fmt"
	"github.com/PhamDuyKhang/alg/datastructre/linkedlist"
)

func main()  {
	l := linkedlist.LList{}

	l.Add("Hello")
	l.Add("I")
	l.Add("Am")
	l.Add("Pham")
	l.Add("Duy")
	l.Add("Khang")

	l.Print(false)

	fmt.Println()

	l.Print(true)

}
