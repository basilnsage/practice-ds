package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewBST(t *testing.T) {
	bst := NewBST(nil)
	var got, want *Node
	if got, want = bst.Root, nil; got != want {
		t.Errorf("empty tree should have nil Root, got: %v", got)
	}
}

func TestSimpleBST(t *testing.T) {
	bst := NewBST(nil)
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(0)
	bst.Insert(3)
	vals := []int{
		bst.Root.Key,
		bst.Root.Left.Key,
		bst.Root.Right.Key,
		bst.Root.Left.Left.Key,
	}

	if diff := cmp.Diff(vals, []int{2, 1, 3, 0}); diff != "" {
		t.Errorf("manual insertion produced wrong order, diff: %v", diff)
	}

	bst = NewBST([]int{2, 3, 1, 0})
	if diff := cmp.Diff(vals, []int{2, 1, 3, 0}); diff != "" {
		t.Errorf("array insertion produced wrong order, diff: %v", diff)
	}
}

func TestPreOrderTraversal(t *testing.T) {
	vals := []int{2, 3, 1, -1, 9, 7, 4, 8}
	bst := NewBST(vals)
	if diff := cmp.Diff(bst.PreOrder(), []int{2, 1, -1, 3, 9, 7, 4, 8}); diff != "" {
		t.Errorf("wrong PreOrder traversal, diff: %v", diff)
	}

	bst = NewBST(nil)
	if diff := cmp.Diff(bst.PreOrder(), []int(nil)); diff != "" {
		t.Errorf("preOrder traversal of an empty tree should be empty, diff: %v", diff)
	}

	bst = NewBST([]int{0})
	if diff := cmp.Diff(bst.PreOrder(), []int{0}); diff != "" {
		t.Errorf("wrong preOrder traversal of a single-node tree, diff: %v", diff)
	}
}

func TestInOrderTraversal(t *testing.T) {
	vals := []int{2, 3, 1, -1, 9, 7, 4, 8}
	bst := NewBST(vals)
	if diff := cmp.Diff(bst.InOrder(), []int{-1, 1, 2, 3, 4, 7, 8, 9}); diff != "" {
		t.Errorf("wrong InOrder traversal, diff: %v", diff)
	}

	bst = NewBST(nil)
	var want []int
	if diff := cmp.Diff(bst.InOrder(), want); diff != "" {
		t.Errorf("InOrder traversal of an empty tree should be empty, diff: %v", diff)
	}

	bst = NewBST([]int{0})
	if diff := cmp.Diff(bst.InOrder(), []int{0}); diff != "" {
		t.Errorf("wrong InOrder traversal of a single-node tree, diff: %v", diff)
	}
}

func TestPostOrderTraversal(t *testing.T) {
	vals := []int{2, 3, 1, -1, 9, 7, 4, 8}
	bst := NewBST(vals)
	if diff := cmp.Diff(bst.PostOrder(), []int{-1, 1, 4, 8, 7, 9, 3, 2}); diff != "" {
		t.Errorf("wrong PostOrder traversal, diff: %v", diff)
	}

	bst = NewBST(nil)
	var want []int
	if diff := cmp.Diff(bst.PostOrder(), want); diff != "" {
		t.Errorf("PostOrder traversal of an empty tree should be empty, diff: %v", diff)
	}

	bst = NewBST([]int{0})
	if diff := cmp.Diff(bst.PostOrder(), []int{0}); diff != "" {
		t.Errorf("wrong PostOrder traversal of a single-node tree, diff: %v", diff)
	}
}

func TestSame(t *testing.T) {
	var n1 *Node = nil
	n2 := &Node{Key: 0}
	n3 := &Node{1, n2, nil}
	if !same(n1, n1) {
		t.Errorf("nil nodes should be equivalent")
	}
	if same(n1, n2) {
		t.Errorf("left-nil, right non-nil should not be the same")
	}
	if same(n2, n1) {
		t.Errorf("right non-nil, left nil should not be the same")
	}
	if !same(n2, n2) {
		t.Errorf("a node should be equivalent to itself")
	}
	if same(n2, n3) {
		diff := cmp.Diff(n2.inOrder(), n3.inOrder())
		t.Errorf("false node equivalence, diff: %v", diff)
	}

	bst1 := NewBST([]int{2, 3, 1, -1, 9, 7, 4, 8, 0})
	bst2 := NewBST([]int{2, 1, -1, 3, 0, 9, 7, 8, 4})
	if !same(bst1.Root, bst2.Root) {
		diff := cmp.Diff(bst1.InOrder(), bst2.InOrder())
		t.Errorf("trees are not the same, diff: %v", diff)
	}
}

func TestPruneMax(t *testing.T) {
	{
		n := &Node{0, nil, nil}
		gotNode, gotKey := pruneMax(n)
		var wantNode *Node = nil
		wantKey := 0
		if gotNode != wantNode {
			t.Errorf("pruning a single-item tree did not return nil value")
		}
		if gotKey != wantKey {
			t.Errorf("wrong max key: %v, want: %v", gotKey, wantKey)
		}
	}

	{
		n := NewBST([]int{2, 5, 1, 7, 6, 9, 8}).Root
		gotNode, gotKey := pruneMax(n)
		wantNode := NewBST([]int{2, 5, 1, 7, 6, 8}).Root
		wantKey := 9
		if !same(gotNode, wantNode) {
			diff := cmp.Diff(gotNode.inOrder(), wantNode.inOrder())
			t.Errorf("wrong pruneMax'd node, diff: %v", diff)
		}
		if gotKey != wantKey {
			t.Errorf("wrong max key: %v, want %v", gotKey, wantKey)
		}
	}
}

func TestPruneMin(t *testing.T) {
	{
		n := &Node{0, nil, nil}
		gotNode, gotKey := pruneMin(n)
		var wantNode *Node = nil
		wantKey := 0
		if gotNode != wantNode {
			t.Errorf("pruning a single-item tree did not return nil value")
		}
		if gotKey != wantKey {
			t.Errorf("wrong min key: %v, want: %v", gotKey, wantKey)
		}
	}

	{
		n := NewBST([]int{7, 4, 1, 2, 6, 5}).Root
		gotNode, gotKey := pruneMin(n)
		wantNode := NewBST([]int{7, 4, 2, 6, 5}).Root
		wantKey := 1
		if !same(gotNode, wantNode) {
			diff := cmp.Diff(gotNode.inOrder(), wantNode.inOrder())
			t.Errorf("wrong pruneMin'd node, diff: %v", diff)
		}
		if gotKey != wantKey {
			t.Errorf("wrong min key: %v, want %v", gotKey, wantKey)
		}
	}
}

func TestHeight(t *testing.T) {
	var n *Node = nil
	if got, want := height(n), 0; got != want {
		t.Errorf("wrong height of nil Node: %v, want: %v", got, want)
	}

	n1 := &Node{0, nil, nil}
	if got, want := height(n1), 1; got != want {
		t.Errorf("wrong height of single Node: %v, want: %v", got, want)
	}

	n2 := NewBST([]int{2, 1, 5, 7}).Root
	if got, want := height(n2), 3; got != want {
		t.Errorf("wrong height for imbalanced tree: %v, want: %v", got, want)
	}

	n3 := NewBST([]int{2, 1, 3}).Root
	if got, want := height(n3), 2; got != want {
		t.Errorf("wrong height for balanced tree: %v, want: %v", got, want)
	}
}

func TestDelete(t *testing.T) {
	{
		n := NewBST([]int{0})
		n.Delete(0)
		var want *Node = nil
		if got := n.Root; got != want {
			t.Errorf("deleting single-item node did not produce nil tree, got: %v", got.inOrder())
		}
	}

	{
		n := NewBST([]int{2, 1})
		n.Delete(2)
		if got, want := n.Root, NewBST([]int{1}).Root; !same(got, want) {
			diff := cmp.Diff(got.inOrder(), want.inOrder())
			t.Errorf("wrong result when deleting with nil left subtree, diff: %v", diff)
		}
	}

	{
		n := NewBST([]int{1, 2})
		n.Delete(1)
		if got, want := n.Root, NewBST([]int{2}).Root; !same(got, want) {
			diff := cmp.Diff(got.inOrder(), want.inOrder())
			t.Errorf("wrong result when deleting with nil left subtree, diff: %v", diff)
		}
	}

	{
		n := NewBST([]int{2, 1, 3})
		n.Delete(2)
		if got, want := n.Root, NewBST([]int{3, 1}).Root; !same(got, want) {
			diff := cmp.Diff(got.inOrder(), want.inOrder())
			t.Errorf("wrong result when deleting from small balanced tree, diff: %v", diff)
		}
	}

	{
		n := NewBST([]int{4, 1, 7, 6, 5})
		n.Delete(4)
		if got, want := n.Root, NewBST([]int{5, 1, 7, 6}).Root; !same(got, want) {
			diff := cmp.Diff(got.inOrder(), want.inOrder())
			t.Errorf("wrong result when deleting from small imbalanced tree, diff: %v", diff)
		}
	}

	{
		n := NewBST([]int{3, 1, 8, 7, 5, 4, 6})
		n.Delete(5)
		if got, want := n.Root, NewBST([]int{3, 1, 8, 7, 6, 4}).Root; !same(got, want) {
			diff := cmp.Diff(got.inOrder(), want.inOrder())
			t.Errorf("wrong result when deleting from small imbalanced tree, diff: %v", diff)
		}
	}
}
