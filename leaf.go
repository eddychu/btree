package btree

import (
	"fmt"
	"sort"
)

type LeafNode struct {
	entries []*Entry
}

func (n *LeafNode) Type() NodeType {
	return LEAF_NODE
}

func NewLeafNode() *LeafNode {
	return &LeafNode{
		entries: make([]*Entry, 0, NUMBER_OF_KEYS),
	}
}

func (n *LeafNode) isEmpty() bool {
	return len(n.entries) == 0
}

func (n *LeafNode) isFull() bool {
	return len(n.entries) == NUMBER_OF_KEYS
}

func (n *LeafNode) Insert(key uint8, value string) (uint8, *LeafNode) {
	if !n.isFull() {
		n.entries = append(n.entries, NewEntry(key, value))
		sort.SliceStable(n.entries, func(i, j int) bool {
			return n.entries[i].key < n.entries[j].key
		})
		return NUMBER_OF_KEYS + 1, nil
	}
	m, newNode := n.Split()
	if key < m {
		n.Insert(key, value)
		return m, newNode
	} else {
		newNode.Insert(key, value)
		return m, newNode
	}
}

func (n *LeafNode) Split() (uint8, *LeafNode) {
	if !n.isFull() {
		panic("Split: should not split an empty or non-full node")
	}
	index := uint8(len(n.entries) / 2)
	key := n.entries[index].key
	fmt.Println("Split: m =", index)
	newNode := NewLeafNode()
	newNode.entries = append(newNode.entries, n.entries[index:]...)
	n.entries = n.entries[:index]
	return key, newNode
}
