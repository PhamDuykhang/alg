package linkedlist

import (
	"errors"
	"fmt"
	"time"
)

type (
	Node struct {
		Next      *Node
		Data      interface{}
		CreatedAt time.Time
	}
	LList struct {
		Head   *Node
		Tail   *Node
		Length int
	}

	DNode struct {
		Data      interface{}
		CreatedAt time.Time
		Next      *DNode
		Prev      *DNode
	}

	DoublyList struct {
		Font *DNode
		End  *DNode
		Len  int
	}
)

func (l *LList) Add(value interface{}) {
	n := &Node{
		Data:      value,
		CreatedAt: time.Now(),
	}
	if l.Length == 0 {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n

	}
	l.Length++

}
func (l *LList) Prepend(value interface{}) {
	n := &Node{
		Data:      value,
		CreatedAt: time.Now(),
	}
	n.Next = l.Head
	l.Head = n
	if l.Tail == nil {
		l.Tail = n
	}
	l.Length++
}
func (l *LList) Search(v interface{}) (ok bool, n Node) {
	if l.Length == 0 {
		panic("list is empty")
	}
	curr := l.Head

	for curr.Next != nil && curr.Data != v {
		curr = curr.Next
	}
	if curr.Next == nil {
		return false, Node{}
	}
	return true, Node{
		Data:      curr.Data,
		CreatedAt: curr.CreatedAt,
	}

}

func (l LList) Print(rev bool) {

	if l.Length == 0 {
		panic("list is empty")
	}
	if rev {
		if l.Tail == nil {
			panic("list don't have tail element")
		}
		fmt.Printf("Liked list have %d element\n", l.Length)
		cur := l.Tail
		for cur != l.Head {
			per := l.Head
			for per.Next != cur {
				per = per.Next
			}
			fmt.Printf("%v -----> ", cur.Data)
			cur = per
		}
		fmt.Printf("%v\n", l.Head.Data)

	} else {
		curr := l.Head
		fmt.Printf("Liked list have %d element\n", l.Length)
		for curr.Next != nil {
			fmt.Printf("%v -----> ", curr.Data)
			curr = curr.Next
		}
		fmt.Printf("%v\n", l.Tail.Data)
	}

}

func (l *LList) Remove(v interface{}) error {
	if l.Length == 0 {
		panic("list is empty")
	}
	if l.Head.Data == v {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}
	curr := l.Head
	var per *Node = l.Head
	for curr.Next != nil && curr.Data != v {
		per = curr
		curr = curr.Next
	}
	if curr.Next == nil {
		l.Tail = per
		l.Length--
		return nil
	}
	if curr.Next == l.Tail {
		per.Next = l.Tail
		l.Length--
		return nil
	}
	per.Next = curr.Next
	l.Length--
	return nil
}

/* Double linked list*/

func (dl *DoublyList) AddEnd(v interface{}) {
	n := &DNode{
		Data:      v,
		CreatedAt: time.Now(),
		Next:      nil,
		Prev:      dl.End,
	}
	if dl.End == nil {
		dl.Font = n
	}

	n.Prev = dl.End
	n.Next = nil
	dl.End = n
	dl.Len++
}

func (dl *DoublyList) AddFont(v interface{}) {
	n := &DNode{
		Data:      v,
		CreatedAt: time.Now(),
		Next:      dl.Font,
		Prev:      nil,
	}
	if dl.Font == nil {
		dl.End = n
	}
	dl.Font.Prev = n
	dl.Font = n
	dl.Len++
}

func (dl *DoublyList) AddBefore(at interface{}, v interface{}) error {
	if dl.Font == nil || dl.End == nil {
		panic("nil list")
	}

	if dl.Font.Data == at {
		dl.AddFont(v)
		return nil
	}
	if dl.End.Data == at {
		dl.AddFont(v)
		return nil
	}
	// find node to add

	cur := dl.Font

	for cur != nil && cur.Data != at {
		cur = cur.Next
	}
	if cur == nil {
		return errors.New("can't find")
	}

	n := &DNode{
		Data:      v,
		CreatedAt: time.Now(),
		Next:      cur,
		Prev:      cur.Prev,
	}
	cur.Prev = n
	return nil
}
func (dl *DoublyList) AddAfter(at interface{}, v interface{}) error {
	if dl.Font == nil || dl.End == nil {
		panic("nil list")
	}

	if dl.Font.Data == at {
		dl.AddEnd(v)
		return nil
	}
	if dl.End.Data == at {
		dl.AddEnd(v)
		return nil
	}
	// find node to add

	cur := dl.Font

	for cur != nil && cur.Data != at {
		cur = cur.Next
	}
	if cur == nil {
		return errors.New("can't find")
	}

	n := &DNode{
		Data:      v,
		CreatedAt: time.Now(),
		Next:      cur,
		Prev:      cur.Prev,
	}
	cur.Prev = n
	return nil
}

func (dl *DoublyList)Remove(){
	
}