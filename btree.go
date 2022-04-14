package btree

import "fmt"

type NodeType int

const (
	LEAF_NODE NodeType = iota
	INTERNAL_NODE
)

const NUMBER_OF_KEYS = 10

type Node interface {
	Type() NodeType
	Split() (string, Node, error)
	Find(key string) *Entry
	IsRoot() bool
	Parent() *InternalNode
}

type InternalNode struct {
	keys   []string
	nodes  []Node
	parent *InternalNode
}

func (n *InternalNode) IsRoot() bool {
	return n.parent == nil
}

func (n *InternalNode) Parent() *InternalNode {
	return n.parent
}

func (n *InternalNode) Type() NodeType {
	return INTERNAL_NODE
}

func NewInternalNode(p *InternalNode) *InternalNode {
	return &InternalNode{
		keys:   make([]string, 0, NUMBER_OF_KEYS),
		nodes:  make([]Node, 0, NUMBER_OF_KEYS+1),
		parent: p,
	}
}

func (n *InternalNode) Split() (string, Node, error) {
	if len(n.keys) < NUMBER_OF_KEYS {
		panic("should not happen")
	}
	sibling := NewInternalNode(n.parent)
	middle := NUMBER_OF_KEYS / 2
	middle_key := n.keys[middle]
	// copy the end half of the entries to the sibling
	copy(sibling.keys, n.keys[middle:])
	copy(sibling.nodes, n.nodes[middle+1:])
	return middle_key, sibling, nil
}

func (n *InternalNode) Find(key string) *Entry {
	var c Node = n
	for c.Type() == INTERNAL_NODE {
		i := 0
		for i < len(c.(*InternalNode).keys) && key > c.(*InternalNode).keys[i] {
			i++
		}
		if i == len(c.(*InternalNode).keys) {
			c = c.(*InternalNode).nodes[i]
		} else if key == c.(*InternalNode).keys[i] {
			c = c.(*InternalNode).nodes[i+1]
		} else {
			c = c.(*InternalNode).nodes[i]
		}
	}

	return c.(*LeafNode).Find(key)
}

func (n *InternalNode) findNode(key string) Node {
	if len(n.keys) == 0 {
		fmt.Println("bingo")
		return n
	}
	var c Node = n
	for c.Type() == INTERNAL_NODE {
		i := 0
		for i < len(c.(*InternalNode).keys) && key > c.(*InternalNode).keys[i] {
			i++
		}
		if i == len(c.(*InternalNode).keys) {
			c = c.(*InternalNode).nodes[i]
		} else if key == c.(*InternalNode).keys[i] {
			c = c.(*InternalNode).nodes[i+1]
		} else {
			c = c.(*InternalNode).nodes[i]
		}
	}
	return c
}

func (n *InternalNode) Insert(entry *Entry) error {

	var c Node = n.findNode(entry.Key)

	if c.Type() == LEAF_NODE {
		return c.(*LeafNode).Insert(entry)
	} else {
		// is an empty node.
		new_leaf := NewLeafNode(c.(*InternalNode))
		fmt.Println(new_leaf)
		new_leaf.Insert(entry)
		c.(*InternalNode).insertKey(entry.Key, new_leaf)
	}
	return nil
}

func (n *InternalNode) insertKey(key string, node Node) error {
	if len(n.keys) == NUMBER_OF_KEYS {
		// is full
		middle_key, sibling, err := n.Split()
		if err != nil {
			panic("should not happen")
		}
		// insert the middle key into the parent
		if n.parent == nil {
			new_root := NewInternalNode(nil)
			new_root.keys = append(new_root.keys, middle_key)
			new_root.nodes = append(new_root.nodes, n)
			new_root.nodes = append(new_root.nodes, sibling)
			n.parent = new_root
			sibling.(*InternalNode).parent = new_root
		} else {
			n.parent.insertKey(middle_key, sibling)
		}
	} else {
		if len(n.keys) == 0 {
			// this is a empty node
			n.keys = append(n.keys, key)
			n.nodes = append(n.nodes, nil)
			n.nodes = append(n.nodes, node)
		} else {
			i := 0
			for i < len(n.keys) && key > n.keys[i] {
				i++
			}
			// insert the key
			n.keys = append(n.keys, "")
			copy(n.keys[i+1:], n.keys[i:])
			n.keys[i] = key
			n.nodes = append(n.nodes, nil)
			copy(n.nodes[i+2:], n.nodes[i+1:])
			n.nodes[i+1] = node
		}

	}
	return nil
}

type LeafNode struct {
	entries []*Entry
	parent  *InternalNode
}

func NewLeafNode(p *InternalNode) *LeafNode {
	return &LeafNode{
		entries: make([]*Entry, 0, NUMBER_OF_KEYS),
		parent:  p,
	}
}

func (n *LeafNode) Type() NodeType {
	return LEAF_NODE
}

func (n *LeafNode) Find(key string) *Entry {
	for i := range n.entries {
		if n.entries[i].Key == key {
			return n.entries[i]
		}
	}
	return nil
}

func (n *LeafNode) Split() (string, Node, error) {
	if len(n.entries) < NUMBER_OF_KEYS {
		panic("should not happen")
	}
	sibling := NewLeafNode(n.parent)
	middle := NUMBER_OF_KEYS / 2
	middle_key := n.entries[middle].Key
	// copy the end half of the entries to the sibling
	copy(sibling.entries, n.entries[middle:])
	return middle_key, sibling, nil
}

func (n *LeafNode) Insert(entry *Entry) error {
	if len(n.entries) == NUMBER_OF_KEYS {
		// is full
		middle_key, sibling, err := n.Split()
		if err != nil {
			panic("should not happen")
		}
		// insert the middle key into the parent
		n.parent.insertKey(middle_key, sibling)
		if entry.Key >= middle_key {
			sibling.(*LeafNode).insertEntry(entry)
		} else {
			n.insertEntry(entry)
		}
	} else {
		n.insertEntry(entry)
	}
	return nil
}

func (n *LeafNode) insertEntry(entry *Entry) error {
	if len(n.entries) == NUMBER_OF_KEYS {
		panic("should not happen")
	} else {
		n.entries = append(n.entries, entry)
		return nil
	}
}

func (n *LeafNode) IsRoot() bool {
	return false
}

func (n *LeafNode) Parent() *InternalNode {
	return n.parent
}
