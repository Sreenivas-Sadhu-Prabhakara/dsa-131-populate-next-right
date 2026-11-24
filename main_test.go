package main

import "testing"

func TestConnect(t *testing.T) {
	n4, n5, n6, n7 := &Node{Val: 4}, &Node{Val: 5}, &Node{Val: 6}, &Node{Val: 7}
	n2 := &Node{Val: 2, Left: n4, Right: n5}
	n3 := &Node{Val: 3, Left: n6, Right: n7}
	root := &Node{Val: 1, Left: n2, Right: n3}
	Connect(root)

	if root.Next != nil {
		t.Error("root has no right neighbor")
	}
	if n2.Next != n3 || n3.Next != nil {
		t.Error("level 1 wiring wrong")
	}
	if n4.Next != n5 || n5.Next != n6 || n6.Next != n7 || n7.Next != nil {
		t.Error("level 2 wiring wrong (cousin link 5->6 is the hard one)")
	}

	if Connect(nil) != nil {
		t.Error("nil root")
	}
	single := &Node{Val: 1}
	if Connect(single) != single || single.Next != nil {
		t.Error("single node")
	}
}
