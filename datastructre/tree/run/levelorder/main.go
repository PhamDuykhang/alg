package main

import (
	"fmt"

	"github.com/PhamDuyKhang/alg/datastructre/tree"
)

func main() {
	root := &tree.TreeNode{}

	result := levelOrder(root)
	fmt.Println(result)

}
func levelOrder(root *tree.TreeNode) [][]int {
	queue := tree.NewQ()
	queue.Add(root)
	var level [][]int
	for !queue.IsEmpty() {
		var curentLevel []int
		l := queue.Lenth()
		for i := 1; i <= l; i++ {
			cNode := queue.Get()
			curentLevel = append(curentLevel, cNode.Value)
			if cNode.Left != nil {
				queue.Add(cNode.Left)
			}
			if cNode.Right != nil {
				queue.Add(cNode.Right)
			}
		}
		level = append(level, curentLevel)
	}
	return level
}
