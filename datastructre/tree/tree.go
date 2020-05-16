package tree

type (
	TreeNode struct {
		Left  *TreeNode
		Right *TreeNode
		Value int
	}

	Queue struct {
		Ar []*TreeNode
	}
)

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
