package main

import "fmt"

// Node is a binary tree node with an extra Next pointer to its right
// neighbor on the same level.
type Node struct {
	Val               int
	Left, Right, Next *Node
}

// Connect wires every node's Next to the node on its right in a PERFECT
// binary tree, in O(1) extra space: each completed level becomes the
// linked list used to walk and wire the level below.
func Connect(root *Node) *Node {
	leftmost := root
	for leftmost != nil && leftmost.Left != nil {
		for n := leftmost; n != nil; n = n.Next {
			n.Left.Next = n.Right // siblings
			if n.Next != nil {
				n.Right.Next = n.Next.Left // cousins, across the gap
			}
		}
		leftmost = leftmost.Left
	}
	return root
}

func main() {
	root := &Node{Val: 1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
		Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}}}
	Connect(root)
	for level := root; level != nil; level = level.Left {
		for n := level; n != nil; n = n.Next {
			fmt.Print(n.Val, " ")
		}
		fmt.Println()
	}
}
