package main

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// n.Left and n.Right are always preserved EXCEPT FOR when one or the other is nil
// in this case we assign n.Left or n.Right a new node with val depending on relation to the current key
func insert(n *Node, val int) *Node {
	if n == nil {
		return &Node{Key: val}
	}
	if val < n.Key {
		n.Left = insert(n.Left, val)
	} else {
		n.Right = insert(n.Right, val)
	}
	return n
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	lHeight := height(n.Left)
	rHeight := height(n.Right)
	if lHeight < rHeight {
		return rHeight + 1
	} else {
		return lHeight + 1
	}
}

func (n *Node) delete(val int) *Node {
	if val < n.Key && n.Left != nil { // search left subtree for the value to delete
		n.Left = n.Left.delete(val)
	} else if val > n.Key && n.Right != nil { // search right subtree for the value to delete
		n.Right = n.Right.delete(val)
	} else if val == n.Key { // we found the value! delete it
		if n.Left == nil && n.Right == nil { // no subtrees so nothing to move into our current place
			return nil
		}
		if n.Left == nil && n.Right != nil { // move the only available subtree into our position
			return n.Right
		}
		if n.Left != nil && n.Right == nil { // move the only available subtree into our position
			return n.Left
		}
		// easy cases handled, now we have to consider what happens if there are two subtrees
		// delete from the longer subtree to minimize tree height
		lHeight, rHeight := height(n.Left), height(n.Right)
		if lHeight <= rHeight {
			n.Right, n.Key = pruneMin(n.Right)
		} else {
			n.Left, n.Key = pruneMax(n.Left)
		}
	}
	return n
}

// search the tree for the largest smallest value, delete it, and return the value from the just-deleted node
func pruneMax(n *Node) (*Node, int) {
	if n.Right == nil { // this is the largest value
		if n.Left != nil { // bump child into our current position
			return n.Left, n.Key
		}
		return nil, n.Key
	}
	r, max := pruneMax(n.Right)
	n.Right = r
	return n, max
}

func pruneMin(n *Node) (*Node, int) {
	if n.Left == nil {
		if n.Right != nil {
			return n.Right, n.Key
		}
		return nil, n.Key
	}
	l, min := pruneMin(n.Left)
	n.Left = l
	return n, min
}

func same(n1, n2 *Node) bool {
	if n1 == nil { // if n1 is nil, the nodes can only be equivalent iff n2 == nil
		return n2 == nil
	} else if n2 != nil { // neither n1 nor n2 are nil so we need to compare keys and child nodes
		if n1.Key != n2.Key {
			return false
		}
		return same(n1.Left, n2.Left) && same(n1.Right, n2.Right)
	}
	return false // n2 == nil and n1 != nil
}

func (n *Node) preOrder() []int {
	vals := []int{n.Key}
	if n.Left != nil {
		vals = append(vals, n.Left.preOrder()...)
	}
	if n.Right != nil {
		vals = append(vals, n.Right.preOrder()...)
	}
	return vals
}

func (n *Node) inOrder() []int {
	var vals []int
	if n.Left != nil {
		vals = append(vals, n.Left.inOrder()...)
	}
	vals = append(vals, n.Key)
	if n.Right != nil {
		vals = append(vals, n.Right.inOrder()...)
	}
	return vals
}

func (n *Node) postOrder() []int {
	var vals []int
	if n.Left != nil {
		vals = append(vals, n.Left.postOrder()...)
	}
	if n.Right != nil {
		vals = append(vals, n.Right.postOrder()...)
	}
	vals = append(vals, n.Key)
	return vals
}

type BST struct {
	Root *Node
}

func NewBST(vals []int) *BST {
	bst := &BST{nil}
	for _, val := range vals {
		bst.Insert(val)
	}
	return bst
}

func (b *BST) Insert(val int) {
	b.Root = insert(b.Root, val)
}

func (b *BST) Delete(val int) {
	b.Root = b.Root.delete(val)
}

func (b *BST) PreOrder() []int {
	var vals []int
	if b.Root != nil {
		vals = b.Root.preOrder()
	}
	return vals
}
func (b *BST) InOrder() []int {
	var vals []int
	if b.Root != nil {
		vals = b.Root.inOrder()
	}
	return vals
}
func (b *BST) PostOrder() []int {
	var vals []int
	if b.Root != nil {
		vals = b.Root.postOrder()
	}
	return vals
}
