package btree

type InternalNode struct {
	keys     []uint8
	children []Node
}

func (n *InternalNode) Type() NodeType {
	return INTERNAL_NODE
}

func NewInternalNode() *InternalNode {
	return &InternalNode{
		keys:     make([]uint8, 0, NUMBER_OF_KEYS),
		children: make([]Node, 0, NUMBER_OF_KEYS+1),
	}
}

func (n *InternalNode) isEmpty() bool {
	return len(n.keys) == 0
}

func (n *InternalNode) findLeaf(key uint8) *LeafNode {
	if n.isEmpty() {
		return nil
	}
	c := n
	for c.Type() == INTERNAL_NODE {
		i := 0
		for i < len(c.keys) && key > c.keys[i] {
			i++
		}
		if i == len(c.keys) {
			c = NewInternalNode()
		}
	}
	return nil
}

func (n *InternalNode) isFull() bool {
	return len(n.keys) == NUMBER_OF_KEYS
}

func (n *InternalNode) Split() (uint8, *InternalNode) {
	if !n.isFull() {
		panic("Split: should not split an empty or non-full node")
	}

	index := uint8(len(n.keys) / 2)
	key := n.keys[index]
	newNode := NewInternalNode()
	newNode.keys = append(newNode.keys, n.keys[index:]...)
	newNode.children = append(newNode.children, nil)
	newNode.children = append(newNode.children, n.children[index+1:]...)
	n.keys = n.keys[:index]
	n.children = n.children[:index+1]
	return key, newNode
}

func (n *InternalNode) findIndex(key uint8) int {
	i := 0
	for i < len(n.keys) && key > n.keys[i] {
		i++
	}
	return i
}

func (n *InternalNode) Insert(key uint8, child Node) (uint8, *InternalNode) {
	if !n.isFull() {
		i := n.findIndex(key)
		if i == len(n.keys) {
			n.keys = append(n.keys, key)
			if len(n.children) == 0 {
				n.children = append(n.children, nil)
			}
			n.children = append(n.children, child)
		} else if n.keys[i] == key {
			panic("Insert: key already exists")
		} else {
			n.keys = append(n.keys, 0)
			copy(n.keys[i+1:], n.keys[i:])
			n.keys[i] = key
			n.children = append(n.children, nil)
			copy(n.children[i+1:], n.children[i:])
			n.children[i] = child
		}
		return NUMBER_OF_KEYS + 1, nil
	}
	m, newNode := n.Split()
	if key < m {
		n.Insert(key, child)
		return m, newNode
	} else {
		newNode.Insert(key, child)
		return m, newNode
	}
}
