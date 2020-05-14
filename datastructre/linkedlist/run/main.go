package main

import (
	"fmt"
	"github.com/PhamDuyKhang/alg/datastructre/linkedlist"
)

func main()  {
	list := linkedlist.LinkedList{}

	list.Append("hello")
	list.Append("World")
	list.Append("I")
	list.Append("am")
	list.Append("Khang")

	fmt.Println(list.GetLength())

	list.RemoveByTime("am")

	fmt.Println(list.GetLength())

	fmt.Println()
	list.Traverse()
}