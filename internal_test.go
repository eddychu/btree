package btree

import (
	"fmt"
	"testing"
)

func TestInternalSplit(t *testing.T) {
	internal := NewInternalNode()
	internal.Insert(1, NewLeafNode())
	internal.Insert(2, NewLeafNode())

	m, newNode := internal.Insert(3, NewLeafNode())
	if m != 2 {
		t.Log(m)
		t.Error("InternalNode should split at 2")
	}
	fmt.Println(newNode)
	fmt.Println(internal)
}
