package main

import (
	"fmt"
	"sync"
)

type (
	Stack struct {
		A []*AdjListNode
	}
	Queue struct {
		Ar []*AdjListNode
	}
	AdjListNode struct {
		V    int
		next *AdjListNode
	}

	AdjList struct {
		head *AdjListNode
		tail *AdjListNode
	}

	Graph struct {
		Visited []bool
		TotalV  int
		list    []*AdjList
	}

	GraphV2 struct {
		nodes []*Node
		edge  map[Node][]*Node
		mux   *sync.RWMutex
	}
	Node struct {
		v interface{}
	}
)

func (g GraphV2) AddEdge(n1, n2 *Node) {
	g.mux.Lock()
	if g.edge == nil {
		g.edge = make(map[Node][]*Node)
	}
	g.edge[*n1] = append(g.edge[*n1], n2)
	g.edge[*n2] = append(g.edge[*n2], n1)
	g.mux.Unlock()
}

func (l *AdjList) AddToList(n *AdjListNode) {
	//List empty
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	l.tail = n
}
func NewQ() Queue {
	return Queue{}
}

func (q *Queue) Add(v *AdjListNode) {
	q.Ar = append(q.Ar, v)
}

func (q *Queue) IsEmpty() bool {
	if len(q.Ar) == 0 {
		return true
	}
	return false
}
func (q *Queue) Get() *AdjListNode {
	v := q.Ar[0]
	q.Ar = q.Ar[1:]
	return v
}

func (q *Queue) Length() int {
	return len(q.Ar)
}

func (s *Stack) Pop() *AdjListNode {
	if len(s.A) == 0 {
		return nil
	}
	el := s.A[0]
	s.A = s.A[1:]
	return el
}
func (s *Stack) Top() *AdjListNode {
	if len(s.A) == 0 {
		return nil
	}
	return s.A[0]
}

func (s *Stack) Push(v *AdjListNode) {
	s.A = append(s.A, v)
}
func (s *Stack) IsEmpty() bool {
	if len(s.A) <= 0 {
		return true
	}
	return false
}

func NewGraph(numNode int) Graph {
	l := make([]*AdjList, numNode)

	for i, _ := range l {
		l[i] = &AdjList{}
	}

	return Graph{
		Visited: make([]bool, numNode),
		TotalV:  numNode,
		list:    l,
	}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.list[v1].AddToList(&AdjListNode{
		V: v2,
	})
	g.list[v2].AddToList(&AdjListNode{
		V: v1,
	})
	return
}

func (g *Graph) PrintGraph() {
	for i, n := range g.list {
		if n == nil {
			continue
		}
		printLine(i, n)
		fmt.Println()
	}
}

func (g *Graph) DeepFirstSearch() {
	g.dfs(2)
}

func (g *Graph) dfs(n int) {
	now := g.list[n]
	g.Visited[n] = true
	fmt.Printf("visted v=%d ", n)
	nn := now.head
	for nn != nil {
		if !g.Visited[nn.V] {
			g.dfs(nn.V)
		}
		nn = nn.next
	}

}

func (g *Graph) BreathFirstSearch(start int) {
	queue := NewQ()
	visited := make([]bool, g.TotalV)
	node := g.list[start]
	queue.g

}

func printLine(v int, n *AdjList) {
	current := n.head
	fmt.Printf("%d ---> ", v)
	for current != nil {
		fmt.Printf("%d", current.V)
		current = current.next
		if current != nil {
			fmt.Print(" ---> ")
		}
	}
}

func main() {
	g := NewGraph(5)

	g.AddEdge(0, 1)
	g.AddEdge(0, 4)
	g.AddEdge(1, 2)
	g.AddEdge(1, 4)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	g.PrintGraph()

	//fmt.Println()
	g.DeepFirstSearch()

	fmt.Println()

}
