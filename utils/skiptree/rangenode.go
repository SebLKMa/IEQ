package skiptree

import "fmt"

// RangeNode defines a node in the tree
type RangeNode struct {
	Key   float64
	Lower float64
	Upper float64
	Value float64
	left  *RangeNode //left side of tree
	right *RangeNode //right side of tree
}

// internal function to recursively find the correct place for a node in a tree
func insert(node *RangeNode, newNode *RangeNode) {
	if newNode.Key < node.Key {
		if node.left == nil {
			node.left = newNode
		} else {
			insert(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insert(node.right, newNode)
		}
	}
}

// internal recursive function to remove an item by its key
func remove(node *RangeNode, key float64) *RangeNode {
	if node == nil {
		return nil
	}

	if key < node.Key {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.Key {
		node.right = remove(node.right, key)
		return node
	}
	// Not Found, key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftmostrightside := node.right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	node.Key, node.Value = leftmostrightside.Key, leftmostrightside.Value
	node.right = remove(node.right, node.Key)
	return node
}

// internal recursive function to search an item in the tree
func search(n *RangeNode, rr float64) (value float64, found bool) {
	if n == nil {
		return 0.0, false
	}

	//if node.lower <= key <= node.upper:
	//return node.value

	//fmt.Printf("Lower:%f RR:%f Upper:%f Value:%f\n", n.Lower, rr, n.Upper, n.Value)
	if n.Lower <= rr && rr <= n.Upper {
		return n.Value, true
	}

	if rr < n.Key {
		return search(n.left, rr)
	}
	if rr > n.Key {
		return search(n.right, rr)
	}
	return 0.0, false
}

// returns the contents of a RangeNode
func stringify(n *RangeNode) string {
	return fmt.Sprintf("Key:%g {Lower:%g Upper:%g} Value:%g", n.Key, n.Lower, n.Upper, n.Value)
}
