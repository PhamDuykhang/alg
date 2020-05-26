package main

type (
	TreeNode struct {
		Left  *TreeNode
		Right *TreeNode
		Value string
	}

	Queue struct {
		Ar []*TreeNode
	}
	Stack struct {
		A []*TreeNode
	}
)

func (s *Stack) Pop() *TreeNode {
	if len(s.A) == 0 {
		return nil
	}
	el := s.A[0]
	s.A = s.A[1:]
	return el
}
func (s *Stack) Top() *TreeNode {
	if len(s.A) == 0 {
		return nil
	}
	return s.A[0]
}

func (s *Stack) Length() int {
	return len(s.A)
}

func (s *Stack) Push(v *TreeNode) {
	s.A = append(s.A, v)
}

func NewQ() Queue {
	return Queue{}
}

func (q *Queue) Add(v *TreeNode) {
	q.Ar = append(q.Ar, v)
}

func (q *Queue) IsEmpty() bool {
	if len(q.Ar) == 0 {
		return true
	}
	return false
}
func (q *Queue) Get() *TreeNode {
	v := q.Ar[0]
	q.Ar = q.Ar[1:]
	return v
}

func (q *Queue) Lenth() int {
	return len(q.Ar)
}
