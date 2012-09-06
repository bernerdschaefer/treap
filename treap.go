// Package treap implements a persistent (immutable) treap structure
// for associating integer keys with arbitrary values.
//
// The implementation is based on (Treaps: The Magical Awesome
// BBT)[http://pavpanchekha.com/blog/treap.html]
package treap

import "math/rand"

type Node struct {
	priority int
	Key      int
	Value    interface{}
	left     *Node
	right    *Node
}

// Returns an empty treap node.
func New() (root *Node) {
	return root
}

// Get a value from the treap by key.
//
// Returns the value stored for key. The boolean ok is false if the key was not
// found.
func (root *Node) Get(key int) (value interface{}, ok bool) {
	switch {
	case root == nil:
	case key < root.Key:
		value, ok = root.left.Get(key)
	case key > root.Key:
		value, ok = root.right.Get(key)
	case key == root.Key:
		value, ok = root.Value, true
	}
	return
}

// Determine if a key is stored in the treap.
func (root *Node) Contains(key int) (found bool) {
	switch {
	case root == nil:
	case key < root.Key:
		found = root.left.Contains(key)
	case key > root.Key:
		found = root.right.Contains(key)
	case key == root.Key:
		found = true
	}

	return
}

// Set a value for the provided key.
//
// Returns a new treap with the value set.
func (root *Node) Set(key int, value interface{}) *Node {
	return root.set(key, value, rand.Int())
}

func (root *Node) set(key int, value interface{}, priority int) (node *Node) {
	switch {
	case root == nil:
		node = &Node{priority: priority, Key: key, Value: value}
	case key < root.Key:
		// recurse leftwards and do a left-rotation if necessary
		newNode := &Node{
			priority: root.priority,
			Key:      root.Key,
			Value:    root.Value,
			left:     root.left.set(key, value, priority),
			right:    root.right,
		}

		if newNode.left.priority < newNode.priority {
			node = newNode.leftRotate()
		} else {
			node = newNode
		}
	case key > root.Key:
		// recurse rightwards and do a right-rotation if necessary
		newNode := &Node{
			priority: root.priority,
			Key:      root.Key,
			Value:    root.Value,
			left:     root.left,
			right:    root.right.set(key, value, priority),
		}

		if newNode.right.priority < newNode.priority {
			node = newNode.rightRotate()
		} else {
			node = newNode
		}
	case key == root.Key:
		node = &Node{
			priority: root.priority,
			Key:      root.Key,
			Value:    value,
			left:     root.left,
			right:    root.right,
		}
	}

	return
}

// Split the treap based on the provided key. The key is assumed not to be
// present in the treap.
//
// Returns the new left and right treaps.
func (root *Node) Split(key int) (*Node, *Node) {
	// negative priority ensures it becomes the new root
	new := root.set(key, nil, -1)
	return new.left, new.right
}

// Merge valid left and right treaps and return the new treap.
func Merge(left *Node, right *Node) (node *Node) {
	switch {
	case left == nil:
		node = right
	case right == nil:
		node = left
	case left.priority < right.priority:
		node = &Node{left.priority, left.Key, left.Value, left.left, Merge(left.right, right)}
	case left.priority >= right.priority:
		node = &Node{right.priority, right.Key, right.Value, Merge(right.left, left), right.right}
	}
	return
}

// Removes the node with the provided key from the treap.
//
// Returns the new treap without the keyed node.
func (root *Node) Delete(key int) (node *Node) {
	switch {
	case root == nil:
	case key < root.Key:
		node = &Node{
			root.priority,
			root.Key,
			root.Value,
			root.left.Delete(key),
			root.right,
		}
	case key > root.Key:
		node = &Node{
			root.priority,
			root.Key,
			root.Value,
			root.left,
			root.right.Delete(key),
		}
	case key == root.Key:
		node = Merge(root.left, root.right)
	}

	return
}

func (root *Node) Keys(ch chan int) {
	if root == nil {
		return
	}

	root.left.Keys(ch)
	ch <- root.Key
	root.left.Keys(ch)
}

func (node *Node) leftRotate() *Node {
	return &Node{
		priority: node.left.priority,
		Key:      node.left.Key,
		Value:    node.left.Value,
		left:     node.left.left,
		right:    &Node{node.priority, node.Key, node.Value, node.left.right, node.right},
	}
}

func (node *Node) rightRotate() *Node {
	return &Node{
		priority: node.right.priority,
		Key:      node.right.Key,
		Value:    node.right.Value,
		left:     &Node{node.priority, node.Key, node.Value, node.left, node.right.left},
		right:    node.right.right,
	}
}
