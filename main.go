package main

import (
	"fmt"
	"sort"
)

const MAX_KEYS = 2

type NodeType int

const (
	LEAF NodeType = iota
	INTERNAL
)

type Node interface {
	Type() NodeType
	Parent() *InternalNode
	SetParent(node *InternalNode)
}

// ------------------------------------------------------------
// Leaf

type LeafNode struct {
	keys   []int
	parent *InternalNode
}

func (n *LeafNode) Type() NodeType {
	return LEAF
}

func NewLeafNode() *LeafNode {
	return &LeafNode{
		keys:   make([]int, 0, MAX_KEYS),
		parent: nil,
	}
}

func (n *LeafNode) Parent() *InternalNode {
	return n.parent
}

func (n *LeafNode) SetParent(parent *InternalNode) {
	n.parent = parent
}

type InternalNode struct {
	keys     []int
	children []Node
	parent   *InternalNode
}

func (n *InternalNode) Type() NodeType {
	return INTERNAL
}

func NewInternalNode() *InternalNode {
	return &InternalNode{
		keys:     make([]int, 0, MAX_KEYS),
		children: make([]Node, 0, MAX_KEYS+1),
	}
}

func (n *InternalNode) Parent() *InternalNode {
	return n.parent
}

func (n *InternalNode) SetParent(parent *InternalNode) {
	n.parent = parent
}

// ------------------------------------------------------------

type BTree struct {
	root Node
}

func NewBTree() *BTree {
	return &BTree{
		root: NewLeafNode(),
	}
}

func (tree *BTree) Insert(key int) {
	root := tree.root
	if root.Type() == LEAF {
		midKey, newNode := tree.insertIntoLeaf(root.(*LeafNode), key)
		if newNode != nil {
			newRoot := NewInternalNode()
			newRoot.keys = append(newRoot.keys, midKey)
			newRoot.children = append(newRoot.children, root)
			newRoot.children = append(newRoot.children, newNode)
			root.SetParent(newRoot)
			newNode.SetParent(newRoot)
			tree.root = newRoot
		}
	} else {
		tree.insertIntoInternal(root.(*InternalNode), key)
	}
}

func (tree *BTree) Print() {
	tree.printNode(tree.root)
}

func (tree *BTree) printNode(node Node) {

	if node.Type() == LEAF {
		tree.printLeaf(node.(*LeafNode))
	} else {
		tree.printInternal(node.(*InternalNode))
	}
}

func (tree *BTree) printInternal(node *InternalNode) {
	fmt.Printf("keys: %v\n", node.keys)
	for _, v := range node.children {
		if v != nil {
			tree.printNode(v)
		} else {
			fmt.Println("[nil]")
		}

	}
}

func (tree *BTree) printLeaf(node *LeafNode) {
	fmt.Printf("%v\n", node.keys)
}

func (tree *BTree) insertIntoLeaf(node *LeafNode, key int) (int, Node) {
	if len(node.keys) == MAX_KEYS {
		// TODO: split
		mid := len(node.keys) / 2
		midKey := node.keys[mid]
		newNode := NewLeafNode()
		newNode.keys = append(newNode.keys, node.keys[mid:]...)
		node.keys = node.keys[:mid]
		if key < midKey {
			tree.insertIntoLeaf(node, key)
		} else {
			tree.insertIntoLeaf(newNode, key)
		}
		return midKey, newNode
	} else {
		node.keys = append(node.keys, key)
		sort.Ints(node.keys)
		return 0, nil
	}
}

func (tree *BTree) insertIntoInternal(node *InternalNode, key int) {
	idx := findInChildren(node, key)
	for node.children[idx].Type() != LEAF {
		node = node.children[idx].(*InternalNode)
		idx = findInChildren(node, key)
	}

	midKey, newLeaf := tree.insertIntoLeaf(node.children[idx].(*LeafNode), key)

	if newLeaf != nil {

		tree.insertIntoParent(node, midKey, newLeaf)
	}

	// if newLeaf != nil {
	// 	newNode := NewInternalNode()
	// 	newNode.keys = append(newNode.keys, midKey)
	// 	newNode.children = append(newNode.children, node.children[idx])
	// 	newNode.children = append(newNode.children, newLeaf)
	// 	fmt.Printf("inserting %d newnode %v\n", midKey, newNode)
	// 	tree.insertIntoParent(node, midKey, newNode)
	// }
}

func (tree *BTree) insertIntoParent(node *InternalNode, key int, child Node) {
	if len(node.keys) == MAX_KEYS {
		// split
		mid := len(node.keys) / 2
		midKey := node.keys[mid]
		newNode := NewInternalNode()
		newNode.keys = append(newNode.keys, node.keys[mid:]...)
		newNode.children = append(newNode.children, nil)
		newNode.children = append(newNode.children, node.children[mid+1:]...)
		node.keys = node.keys[:mid]
		node.children = node.children[:mid+1]
		if key < midKey {
			tree.insertIntoParent(node, key, child)
		} else {
			tree.insertIntoParent(newNode, key, child)
		}

		if node == tree.root {
			newRoot := NewInternalNode()
			newRoot.keys = append(newRoot.keys, midKey)
			newRoot.children = append(newRoot.children, node)
			newRoot.children = append(newRoot.children, newNode)
			node.parent = newRoot
			tree.root = newRoot
		} else {
			tree.insertIntoParent(node.parent, midKey, newNode)
		}
	} else {
		i := 0
		for len(node.keys) > i && key > node.keys[i] {
			i++
		}
		node.keys = append(node.keys, 0)
		copy(node.keys[i+1:], node.keys[i:])
		node.keys[i] = key
		node.children = append(node.children, nil)
		copy(node.children[i+2:], node.children[i+1:])
		node.children[i+1] = child
		child.SetParent(node)
	}
}

// func insertIntoLeaf(node *LeafNode, key int) (int, Node) {
// 	if len(node.keys) == MAX_KEYS {
// 		// TODO: split
// 		mid := len(node.keys) / 2
// 		midKey := node.keys[mid]
// 		newNode := NewLeafNode()
// 		newNode.keys = append(newNode.keys, node.keys[mid:]...)
// 		node.keys = node.keys[:mid]
// 		if key < midKey {
// 			insertIntoLeaf(node, key)
// 		} else {
// 			insertIntoLeaf(newNode, key)
// 		}
// 		return midKey, newNode
// 	} else {
// 		node.keys = append(node.keys, key)
// 		sort.Ints(node.keys)
// 		return 0, nil
// 	}
// }

func findInChildren(subroot *InternalNode, key int) int {
	i := 0
	for len(subroot.keys) > i && key > subroot.keys[i] {
		i++
	}
	return i
}

func main() {
	tree := NewBTree()
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(8)
	tree.Print()

}
