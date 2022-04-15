package btree

type Entry struct {
	key   uint8
	value string
}

func NewEntry(key uint8, value string) *Entry {
	return &Entry{
		key:   key,
		value: value,
	}
}
