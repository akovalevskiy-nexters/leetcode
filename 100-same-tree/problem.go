package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type DataSource struct {
	P *TreeNode
	q *TreeNode
}

func main() {
	ds := DataSource{
		&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
	}
	// ds := DataSource{
	// 	&TreeNode{1, &TreeNode{2, nil, nil}, nil},
	// 	&TreeNode{1, nil, &TreeNode{2, nil, nil}},
	// }
	fmt.Println(isSameTree(ds.P, ds.q))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return q == p
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
