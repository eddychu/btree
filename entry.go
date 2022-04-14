package btree

type Entry struct {
	Key   string
	Value string
}

func NewEntry(key, value string) *Entry {
	return &Entry{
		Key:   key,
		Value: value,
	}
}
