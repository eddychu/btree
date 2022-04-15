package btree

import "testing"

func TestInsert(t *testing.T) {
	btree := NewBtree()
	btree.Insert(1, "a")

	if btree.root.Type() != LEAF_NODE {
		t.Error("Root should be a leaf node")
	}
	btree.Insert(2, "b")
	btree.Insert(3, "c")
	if btree.root.Type() != INTERNAL_NODE {
		t.Error("Root should be an internal node")
	}

	btree.Insert(4, "d")
	btree.Insert(5, "e")
	// btree.Insert(6, "f")
	// btree.Insert(7, "g")

	t.Log(btree.root)
	t.Log(btree.root.(*InternalNode).children[0])
	t.Log(btree.root.(*InternalNode).children[1])
}
