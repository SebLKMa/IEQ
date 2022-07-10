package skiptree

import "fmt"

// ItemSkipTree the binary search tree of Items to embed functions to this type
type ItemSkipTree struct {
	root *RangeNode
}

// Insert inserts the Item t in the tree, starting from root traversing to the left
// or right side of tree. The `lower` range value is the key.
func (st *ItemSkipTree) Insert(lower float64, upper float64, value float64) {
	n := &RangeNode{lower, lower, upper, value, nil, nil}
	if st.root == nil {
		st.root = n
	} else {
		insert(st.root, n)
	}
}

// Remove removes the Item with key `key` from the tree
func (st *ItemSkipTree) Remove(key float64) {
	remove(st.root, key)
}

// Search returns the Value most closely associated with Relative Range
func (st *ItemSkipTree) Search(rr float64) (value float64, found bool) {
	//fmt.Printf("Search Value: %g\n", rr)
	return search(st.root, rr)
}

// String prints a visual representation of the tree
func (st *ItemSkipTree) String() {
	fmt.Println("------------------------------------------------")
	st.stringify(st.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func (st *ItemSkipTree) stringify(n *RangeNode, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "   "
		}
		format += "---[ "
		level++
		st.stringify(n.left, level)
		fmt.Println(format + stringify(n))
		st.stringify(n.right, level)
	}
}
