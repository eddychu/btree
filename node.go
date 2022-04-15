package btree

type NodeType int

const NUMBER_OF_KEYS = 2

const (
	INTERNAL_NODE NodeType = iota
	LEAF_NODE
)

type Node interface {
	Type() NodeType
}
