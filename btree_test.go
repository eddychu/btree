package btree

import (
	"fmt"
	"testing"
)

func TestNewRoot(t *testing.T) {
	root := NewInternalNode(nil)
	if !root.IsRoot() {
		t.Error("root should not be root")
	}
	if root.Parent() != nil {
		t.Error("root should not have a parent")
	}
	if root.Type() != INTERNAL_NODE {
		t.Error("root should be a internal node")
	}

	if len(root.keys) != 0 {
		t.Log(len(root.keys))
		t.Error("root by default should not have keys")
	}

	if len(root.nodes) != 0 {
		t.Log(len(root.nodes))
		t.Error("root by default should not have nodes")
	}
}

func TestInsert(t *testing.T) {
	root := NewInternalNode(nil)
	root.Insert(NewEntry("a", "a"))
	// root.Insert(NewEntry("b", "b"))
	// root.Insert(NewEntry("c", "c"))
	// root.Insert(NewEntry("d", "d"))
	// root.Insert(NewEntry("e", "e"))
	// root.Insert(NewEntry("f", "f"))
	// root.Insert(NewEntry("g", "g"))
	// root.Insert(NewEntry("h", "h"))
	// root.Insert(NewEntry("i", "i"))
	// root.Insert(NewEntry("j", "j"))
	// root.Insert(NewEntry("k", "k"))
	// root.Insert(NewEntry("l", "l"))
	// root.Insert(NewEntry("m", "m"))
	// root.Insert(NewEntry("n", "n"))
	// root.Insert(NewEntry("o", "o"))
	// root.Insert(NewEntry("p", "p"))
	// root.Insert(NewEntry("q", "q"))
	// root.Insert(NewEntry("r", "r"))
	// root.Insert(NewEntry("s", "s"))
	// root.Insert(NewEntry("t", "t"))
	// root.Insert(NewEntry("u", "u"))
	// root.Insert(NewEntry("v", "v"))
	// root.Insert(NewEntry("w", "w"))
	// root.Insert(NewEntry("x", "x"))
	// root.Insert(NewEntry("y", "y"))
	// root.Insert(NewEntry("z", "z"))
	if len(root.keys) == 0 {
		t.Error("root should have keys")
	}
	fmt.Printf("%+v\n", root)
}
