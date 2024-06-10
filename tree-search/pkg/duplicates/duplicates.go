package duplicates

import (
	"container/list"
	"errors"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NaryTreeNode struct {
	Val int
}

type NaryNodeAndLevel struct {
	Node  NaryTreeNode
	Level int
}

type NodeAndLevel struct {
	Node  *TreeNode
	Level int
}

type Results struct {
	Value *int
	Level int
}

func CheckDuplicateIDs(root *TreeNode) (*Results, error) {
	if root == nil {
		return &Results{Value: nil, Level: 0}, errors.New("no tree provided")
	}

	queue := list.New()
	queue.PushBack(&NodeAndLevel{Node: root, Level: 0})
	seen := make(map[int]bool)

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		curr := front.Value.(*NodeAndLevel)

		if seen[curr.Node.Val] {
			return &Results{Value: &curr.Node.Val, Level: curr.Level}, nil
		} else {
			seen[curr.Node.Val] = true
		}

		if curr.Node.Left != nil {
			queue.PushBack(&NodeAndLevel{Node: curr.Node.Left, Level: curr.Level + 1})
		}

		if curr.Node.Right != nil {
			queue.PushBack(&NodeAndLevel{Node: curr.Node.Right, Level: curr.Level + 1})
		}
	}

	return &Results{Value: nil, Level: 0}, nil
}
