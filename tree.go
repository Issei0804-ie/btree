package my_btree

const (
	maxDegree = 3
)

type BTree struct {
	top    *node
	length int
}

type node struct {
	Items    []Item
	Children []*node
}

type Item interface {
	// EqualOrMore if thisItem >= tar, return true. else false.
	EqualOrMore(tar Item)
	// Less if thisItem < tar, return true. else false.
	Less(tar Item) bool
}

func (b BTree) find(item Item) (ok bool, idx int) {

}