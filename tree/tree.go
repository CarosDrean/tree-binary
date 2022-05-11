package tree

import (
	"fmt"
)

type Ordered interface {
	IsLeft(s, t interface{}) bool
	IsEqual(s, t interface{}) bool
	IsValidType(data interface{}) bool
}

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
	size int

	ordered Ordered
}

func New(ordered Ordered) Tree {
	return Tree{ordered: ordered}
}

func (t *Tree) Add(data interface{}) error {
	if !t.ordered.IsValidType(data) {
		return fmt.Errorf("the type is not valid")
	}

	if t.size == 0 {
		t.root = &Node{data: data}
		t.size++
		return nil
	}

	t.add(t.root, data)

	return nil
}

func (t Tree) InOrder() string {
	result := ""
	t.inOrder(t.root, &result)
	return result
}

func (t Tree) PreOrder() string {
	result := ""
	t.preOrder(t.root, &result)
	return result
}

func (t Tree) PostOrder() string {
	result := ""
	t.postOrder(t.root, &result)
	return result
}

func (t *Tree) Search(data interface{}) (interface{}, bool) {
	if !t.ordered.IsValidType(data) {
		return nil, false
	}

	return t.search(t.root, data)
}

func (t *Tree) Delete(data interface{}) {
	if !t.ordered.IsValidType(data) {
		return
	}

	t.delete(t.root, data)
}

func (t *Tree) add(current *Node, data interface{}) {
	if t.ordered.IsLeft(current.data, data) {
		if current.left == nil {
			current.left = &Node{data: data}
			t.size++
			return
		}

		t.add(current.left, data)
		return
	}

	if current.right == nil {
		current.right = &Node{data: data}
		t.size++
		return
	}

	t.add(current.right, data)
}

func (t Tree) inOrder(n *Node, result *string) {
	if n != nil {
		t.inOrder(n.left, result)

		if *result == "" {
			*result = fmt.Sprintf("%v", n.data)
		} else {
			*result = fmt.Sprintf("%s, %v ", *result, n.data)
		}

		t.inOrder(n.right, result)
	}
}

func (t Tree) preOrder(n *Node, result *string) {
	if n != nil {
		if *result == "" {
			*result = fmt.Sprintf("%v", n.data)
		} else {
			*result = fmt.Sprintf("%s, %v ", *result, n.data)
		}

		t.preOrder(n.left, result)
		t.preOrder(n.right, result)
	}
}

func (t Tree) postOrder(n *Node, result *string) {
	if n != nil {
		t.postOrder(n.left, result)
		t.postOrder(n.right, result)

		if *result == "" {
			*result = fmt.Sprintf("%v", n.data)
		} else {
			*result = fmt.Sprintf("%s, %v ", *result, n.data)
		}
	}
}

func (t Tree) search(n *Node, data interface{}) (interface{}, bool) {
	if n == nil {
		return nil, false
	}

	if t.ordered.IsEqual(n.data, data) {
		return n, true
	}

	if t.ordered.IsLeft(n.data, data) {
		return t.search(n.left, data)
	}

	return t.search(n.right, data)
}

func (t Tree) findMin(n Node) Node {
	if n.left == nil {
		return n
	}

	return t.findMin(*n.left)
}

func (t Tree) delete(n *Node, data interface{}) {
	if n == nil {
		return
	}

	if t.ordered.IsEqual(n.data, data) {
		if n.left == nil && n.right == nil {
			n = nil
			return
		}

		if n.left == nil {
			n = n.right
			return
		}

		if n.right == nil {
			n = n.left
			return
		}

		minRight := t.findMin(*n.right)
		n.data = minRight.data
		t.delete(n.right, minRight.data)

		return
	}

	if t.ordered.IsLeft(n.data, data) {
		t.delete(n.left, data)
		return
	}

	t.delete(n.right, data)
}
