package linkedlist

import (
	"fmt"
	"time"
)

type (
	Element struct {
		Data      interface{}
		Next      *Element
		CreatedAt int64
	}
	LinkedList struct {
		length int
		Start  *Element
		End    *Element
	}
)

func (ll *LinkedList) GetLength() int {
	return ll.length
}

func (ll *LinkedList)Traverse() {
	if ll.length == 0{
		panic("the list is empty")
	}
	current := ll.Start

	for current.Next != nil{
		fmt.Printf("%s at time: %d\n",current.Data,current.CreatedAt)
		fmt.Printf("\t\t\t|\n")
		fmt.Printf("\t\t\t|\n")
		fmt.Printf("\t\t\t|\n")
		fmt.Printf("\t\t\t\\|/\n")
		current = current.Next
	}
	fmt.Printf("%s at time: %d\n",ll.End.Data,ll.End.CreatedAt)
}

func (ll *LinkedList) Append(e interface{}) {
	el := &Element{
		Data:      e,
		CreatedAt: time.Now().Unix(),
	}
	if ll.length == 0 {
		ll.Start = el
		ll.End = el
	} else {
		last := ll.End
		last.Next = el
		ll.End = el

	}
	ll.length++
}
func (ll *LinkedList) RemoveByTime(t string){
	if ll.length == 0{
		panic("list is empty")
	}
	current := ll.Start
	var pre *Element
	for current.Data != t{
		if current.Next == nil{
			panic("not found")
		}
		pre = current
		current = current.Next
	}
	pre.Next = current.Next
	ll.length--
}