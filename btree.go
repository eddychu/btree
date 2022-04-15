package btree

type Btree struct {
	root Node
}

func NewBtree() *Btree {
	return &Btree{
		root: nil,
	}
}

func (t *Btree) findLeaf(key uint8) *LeafNode {
	c := t.root
	if c.Type() == LEAF_NODE {
		return c.(*LeafNode)
	}
	for c.Type() == INTERNAL_NODE {
		i := 0
		for i < len(c.(*InternalNode).keys) && key > c.(*InternalNode).keys[i] {
			i++
		}
		if i == len(c.(*InternalNode).keys) {
			if c.(*InternalNode).children[i] == nil {
				c.(*InternalNode).children[i] = NewLeafNode()
			} else {
				c = c.(*InternalNode).children[i]
			}
		} else if c.(*InternalNode).keys[i] == key {
			c = c.(*InternalNode).children[i+1]
		} else {
			c = c.(*InternalNode).children[i]
		}
	}
	return c.(*LeafNode)
}

func (t *Btree) Insert(key uint8, value string) {
	if t.root == nil {
		t.root = NewLeafNode()
		t.root.(*LeafNode).Insert(key, value)
	} else {
		l := t.findLeaf(key)
		m, newNode := l.Insert(key, value)
		if newNode != nil {
			if t.root.Type() == LEAF_NODE {
				oldRoot := t.root.(*LeafNode)
				t.root = NewInternalNode()
				children := make([]Node, 2, NUMBER_OF_KEYS+1)
				children[0] = oldRoot
				children[1] = newNode
				t.root.(*InternalNode).children = children
				t.root.(*InternalNode).keys = append(t.root.(*InternalNode).keys, m)
				t.root.(*InternalNode).keys[0] = m
			} else {
				m, newNewNode := t.root.(*InternalNode).Insert(m, newNode)
				if newNewNode != nil {
					oldOldRoot := t.root.(*InternalNode)
					t.root = NewInternalNode()
					children := make([]Node, 2, NUMBER_OF_KEYS+1)
					children[0] = oldOldRoot
					children[1] = newNewNode
					t.root.(*InternalNode).children = children
					t.root.(*InternalNode).keys = append(t.root.(*InternalNode).keys, m)
					t.root.(*InternalNode).keys[0] = m
				}
			}
		}
	}
	// t.root.Insert(key, value)
}

func (t *Btree) insertNode(key uint8, node Node) {

}
