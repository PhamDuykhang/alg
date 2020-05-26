package main

import "fmt"

func main() {
	a := "82-*45*9*-"

	root := BuildExpressionTree(a)

	inOrder(root)
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Printf("%s ", root.Value)
	inOrder(root.Right)
}

func BuildExpressionTree(ex string) *TreeNode {
	st := Stack{}
	for _, o := range ex {
		if !isOperator(uint(o)) {
			st.Push(BuildNode(uint(o)))
		} else {
			s1Child := st.Pop()
			s2Child := st.Pop()
			n := BuildNode(uint(o))
			n.Left = s1Child
			n.Right = s2Child
			st.Push(n)
		}
	}
	return st.Top()
}

func isOperator(un uint) bool {
	if un != '+' && un != '-' && un != '*' && un != ':' && un != '^' {
		return false
	}
	return true
}

func BuildNode(un uint) *TreeNode {
	return &TreeNode{
		Value: string(un),
	}
}
