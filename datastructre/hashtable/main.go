package main

import (
	"fmt"
	"math"
)

const (
	arrayLength         = 100
	ThresholdLoadFactor = 0.75
)

type (
	LinkedList struct {
		Head *Node
		Tail *Node
	}
	Node struct {
		Next *Node
		Data interface{}
	}

	HashTable struct {
		NumberBucket int
		Size         int
		data         [arrayLength]*LinkedList
	}

	listData struct {
		key   string
		value interface{}
	}
)

func (h *HashTable) CalculateLF() float64 {
	return float64(h.Size) / float64(h.NumberBucket)
}

func hash(st string) int {
	h := 0

	for pos, char := range st {
		h += int(char) * int(math.Pow(31, float64(len(st)-pos+1)))
	}
	return h
}

func (h *HashTable) Get(k string) (result interface{}, bool2 bool) {
	hash := hash(k)
	index := index(hash)
	if h.data[index] == nil {
		return "", false
	}
	currentNode := h.data[index].Head
	for currentNode != nil {
		d := currentNode.Data.(listData)
		if d.key == k {
			return currentNode.Data, true
		}
	}
	return "", false

}

func (h *HashTable) Set(k string, v interface{}) {
	hash := hash(k)
	index := index(hash)

	if h.data[index] == nil {
		list := NewLikedList()
		h.data[index] = list
		h.data[index].Append(listData{
			key:   k,
			value: v,
		})
	} else {
		node := h.data[index].Head
		for node != nil {
			d := node.Data.(listData)
			if d.key == k {
				d.value = v
				break
			}
			node = node.Next
		}
		h.data[index].Append(listData{
			key:   k,
			value: v,
		})
	}
	h.Size++
	nowLF := h.CalculateLF()
	if nowLF > ThresholdLoadFactor {

	}

	return
}

func index(h int) int {
	return int(math.Abs(float64(h % arrayLength)))
}

func NewHashTable() *HashTable {
	return &HashTable{
		NumberBucket: arrayLength,
		Size:         0,
		data:         [arrayLength]*LinkedList{},
	}
}
func NewLikedList() *LinkedList {
	emptyNode := &Node{
		Next: nil,
		Data: nil,
	}
	return &LinkedList{
		Head: emptyNode,
		Tail: emptyNode,
	}
}
func (ll *LinkedList) Append(d interface{}) *LinkedList {
	nextNode := &Node{
		Next: nil,
		Data: d,
	}
	if ll.Head.Data == nil {
		ll.Head = nextNode
	} else {
		ll.Tail.Next = nextNode
	}
	ll.Tail = nextNode
	return ll
}

// DeleteWithValue : Deletes node which has a specific value
func (ll *LinkedList) DeleteWithValue(v interface{}) *LinkedList {
	var node = ll.Head
	if node.Data == v {
		ll.Head = ll.Head.Next
		return ll
	}
	for {
		if v == node.Next.Data {
			if node.Next.Next != nil {
				node.Next = node.Next.Next
				return ll
			}
			node.Next = nil
			return ll
		}
		node = node.Next
	}
	return ll
}

func (ll *LinkedList) PrintAll() {
	var node = ll.Head
	for {
		fmt.Println(node.Data)
		if node.Next == nil {
			return
		}
		node = node.Next
	}
}

func main() {
	hashMap := NewHashTable()

	hashMap.Set("header", "application/json")
	hashMap.Set("authentication", "jwd")
	hashMap.Set("accept", "application")

	d, ok := hashMap.Get("header")
	if ok {
		fmt.Println(d)
	}
	d1, ok := hashMap.Get("as")
	if ok {
		fmt.Println(d1)
	}
}
