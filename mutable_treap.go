package treap

import "math/rand"

type MutableTreap struct {
	root *Node
}

// Returns a new, mutable treap wrapper.
//
// This is most useful when loading a treap with an initial set of data, before
// you want its safe, persistent properties. Once the load is complete, call
// Treap() to get the persistent data structure.
func NewMutable() *MutableTreap {
	return &MutableTreap{}
}

// Mutates an existing treap, setting key to value.
func (treap *MutableTreap) Set(key int, value interface{}) {
	treap.root = treap.root.setUnsafe(key, value, rand.Int())
}

// Return the underlying treap structure.
func (treap *MutableTreap) Treap() *Node {
	return treap.root
}

func (root *Node) setUnsafe(key int, value interface{}, priority int) *Node {
	switch {
	case root == nil:
		root = &Node{priority: priority, Key: key, Value: value}
	case key < root.Key:
		root.left = root.left.setUnsafe(key, value, priority)
		if root.left.priority < root.priority {
			root.leftRotateUnsafe()
		}
	case key > root.Key:
		root.right = root.right.setUnsafe(key, value, priority)
		if root.right.priority < root.priority {
			root.rightRotateUnsafe()
		}
	case key == root.Key:
		root.Value = value
	}

	return root
}

func (node *Node) leftRotateUnsafe() *Node {
	right := &Node{node.priority, node.Key, node.Value, node.left.right, node.right}

	node.priority = node.left.priority
	node.Key = node.left.Key
	node.Value = node.left.Value
	node.left = node.left.left
	node.right = right
	return node
}

func (node *Node) rightRotateUnsafe() *Node {
	left := &Node{node.priority, node.Key, node.Value, node.right.left, node.left}

	node.priority = node.right.priority
	node.Key = node.right.Key
	node.Value = node.right.Value
	node.right = node.right.right
	node.left = left
	return node
}
