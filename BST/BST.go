package BST

import "fmt"

type node[T any] struct {
	value T
	left  *node[T]
	right *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		value: value,
	}
}

type compareFn[T any] func(a, b T) int // return -1 if a < b , 0 if a == b and 1 if a > b

type BST[T any] struct {
	root      *node[T]
	compareFn compareFn[T]
	length    int
}

func New[T any](compareFn compareFn[T]) *BST[T] {
	return &BST[T]{
		compareFn: compareFn,
	}
}

func (bst *BST[T]) Insert(value T) {
	node := newNode[T](value)
	bst.length++
	if bst.root == nil {
		bst.root = node
		return
	}
	bst.recursiveInsert(bst.root, node)
}

func (bst *BST[T]) recursiveInsert(curr *node[T], node *node[T]) {
	if bst.compareFn(node.value, curr.value) <= 0 {
		if curr.left == nil {
			curr.left = node
			return
		}
		bst.recursiveInsert(curr.left, node)
	} else {
		if curr.right == nil {
			curr.right = node
			return
		}
		bst.recursiveInsert(curr.right, node)
	}
}

func (bst *BST[T]) Search(value T) bool {
	return bst.recursiveSearch(bst.root, value)
}

func (bst *BST[T]) recursiveSearch(curr *node[T], value T) bool {
	if curr == nil {
		return false
	}

	comparation := bst.compareFn(value, curr.value)
	if comparation < 0 {
		return bst.recursiveSearch(curr.left, value)
	} else if comparation > 0 {
		return bst.recursiveSearch(curr.right, value)
	}

	return true
}

func (bst *BST[T]) Delete(value T) bool {
	if bst.root == nil {
		return false
	}

	out := bst.recursiveDelete(bst.root, value, nil, 0)
	if !out {
		return out
	}

	bst.length--
	return out
}

func (bst *BST[T]) recursiveDelete(curr *node[T], value T, parent *node[T], lastComparisson int) bool {
	if curr == nil {
		return false
	}

	comparation := bst.compareFn(value, curr.value)
	if comparation < 0 {
		return bst.recursiveDelete(curr.left, value, curr, comparation)
	} else if comparation > 0 {
		return bst.recursiveDelete(curr.right, value, curr, comparation)
	}

	if curr.right == nil {
		if parent == nil {
			bst.root = curr.left
		} else if lastComparisson < 0 {
			parent.left = curr.left
		} else {
			parent.right = curr.left
		}
		return true
	}

	minor, minorParent := curr.right.findMinorAndDetach(curr)
	minor.left = curr.left
	curr.left = nil
	minorParent.left = minor.right
	minor.right = curr.right
	if parent == nil {
		bst.root = minor
	} else if lastComparisson < 0 {
		parent.left = minor
	} else {
		parent.right = minor
	}
	return true
}

// assume node is not nil
func (node *node[T]) findMinorAndDetach(parent *node[T]) (minor *node[T], minorParent *node[T]) {
	if node.left == nil {
		return node, parent
	}

	return node.left.findMinorAndDetach(node)
}

func (bst *BST[T]) Print() {
	curr := []*node[T]{bst.root}
	for len(curr) > 0 {
		print[T](curr)
		newCurr := make([]*node[T], 0, len(curr)*2)
		for _, node := range curr {
			if node.left != nil {
				newCurr = append(newCurr, node.left)
			}
			if node.right != nil {
				newCurr = append(newCurr, node.right)
			}
		}
		curr = newCurr
	}
}

func print[T any](nodes []*node[T]) {
	out := make([]T, len(nodes))
	for i, node := range nodes {
		out[i] = node.value
	}
	fmt.Println(out)
}
