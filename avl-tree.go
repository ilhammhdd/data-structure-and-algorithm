package main

import (
	"cmp"
	"fmt"
	"strings"
)

type AVLNode[T cmp.Ordered] struct {
	Val    T
	left   *AVLNode[T]
	right  *AVLNode[T]
	height int32
}

func InsertAVLNode[T cmp.Ordered](node *AVLNode[T], val T) *AVLNode[T] {
	if node == nil {
		return &AVLNode[T]{Val: val, height: 1}
	}

	switch {
	case val < node.Val:
		node.left = InsertAVLNode(node.left, val)
	case val > node.Val:
		node.right = InsertAVLNode(node.right, val)
	case val == node.Val:
		return node
	}

	updateHeight(node)
	balancing(node, val)
	// balanceFactor := calculateBalanceFactor(node)

	// if balanceFactor > 1 {
	// 	if val < node.left.Val {
	// 		return rotateRight(node)
	// 	} else if val > node.left.Val {
	// 		node.left = rotateLeft(node.left)
	// 		return rotateRight(node)
	// 	}
	// } else if balanceFactor < -1 {
	// 	if val > node.right.Val {
	// 		return rotateLeft(node)
	// 	} else if val < node.right.Val {
	// 		node.right = rotateRight(node.right)
	// 		return rotateLeft(node)
	// 	}
	// }

	return node
}

func DeleteAVLNode[T cmp.Ordered](root *AVLNode[T], val T) *AVLNode[T] {
	root = deleteAVLNode(root, val)
	if root == nil {
		return nil
	}
	updateHeight(root)
	return balancing(root, val)
}

func deleteAVLNode[T cmp.Ordered](node *AVLNode[T], val T) *AVLNode[T] {
	if node == nil {
		return nil
	}
	switch {
	case val < node.Val:
		leftResult := deleteAVLNode(node.left, val)
		node.left = leftResult
	case val > node.Val:
		rightResult := deleteAVLNode(node.right, val)
		node.right = rightResult
	}

	updateHeight(node)
	node = balancing(node, val)

	if val != node.Val {
		return node
	}

	replacement, parent := findReplacementWithParent(node)
	if replacement == nil {
		return nil
	}
	if replacement.left != nil && node.left != nil && replacement.left.Val != node.left.Val ||
		replacement.left == nil && node.left != nil && replacement.Val != node.left.Val {
		replacement.left = node.left
	}
	if replacement.right != nil && node.right != nil && replacement.right.Val != node.right.Val ||
		replacement.right == nil && node.right != nil && replacement.Val != node.right.Val {
		replacement.right = node.right
	}

	if parent != nil {
		if replacement.Val < parent.Val {
			parent.left = nil
		} else if replacement.Val > parent.Val {
			parent.right = nil
		}
		updateHeight(parent)
	}

	updateHeight(replacement)

	return replacement
}

func balancing[T cmp.Ordered](node *AVLNode[T], val T) (newRoot *AVLNode[T]) {
	balanceFactor := calculateBalanceFactor(node)

	if balanceFactor > 1 {
		if val < node.left.Val {
			return rotateRight(node)
		} else if val > node.left.Val {
			node.left = rotateLeft(node.left)
			return rotateRight(node)
		}
	} else if balanceFactor < -1 {
		if val > node.right.Val {
			return rotateLeft(node)
		} else if val < node.right.Val {
			node.right = rotateRight(node.right)
			return rotateLeft(node)
		}
	}
	return node
}

func updateHeight[T cmp.Ordered](node *AVLNode[T]) {
	if node == nil {
		return
	}
	node.height = 1 + max(getHeight(node.left), getHeight(node.right))
}

func findReplacement[T cmp.Ordered](node *AVLNode[T]) *AVLNode[T] {
	if node == nil {
		return nil
	}

	leftInOrder := inOrder(node.left)
	rightInOrder := inOrder(node.right)

	if len(leftInOrder) > 0 {
		return leftInOrder[len(leftInOrder)-1]
	} else if len(rightInOrder) > 0 {
		return rightInOrder[0]
	}
	return nil
}

func findReplacementWithParent[T cmp.Ordered](node *AVLNode[T]) (replacement, parent *AVLNode[T]) {
	if node == nil {
		return nil, nil
	}

	leftInOrder := inOrder(node.left)
	rightInOrder := inOrder(node.right)

	if len(leftInOrder) > 1 {
		return leftInOrder[len(leftInOrder)-1], leftInOrder[len(leftInOrder)-2]
	} else if len(rightInOrder) > 1 {
		return rightInOrder[0], rightInOrder[1]
	} else if len(leftInOrder) > 0 {
		return leftInOrder[len(leftInOrder)-1], nil
	} else if len(rightInOrder) > 0 {
		return rightInOrder[0], nil
	}
	return nil, nil
}

func inOrder[T cmp.Ordered](node *AVLNode[T]) []*AVLNode[T] {
	if node == nil {
		return []*AVLNode[T]{}
	}
	result := inOrder(node.left)
	result = append(result, node)
	result = append(result, inOrder(node.right)...)
	return result
}

func rotateRight[T cmp.Ordered](node *AVLNode[T]) *AVLNode[T] {
	newRoot := node.left

	newRoot.right, node.left = node, newRoot.right

	updateHeight(node)
	updateHeight(newRoot)

	return newRoot
}

func rotateLeft[T cmp.Ordered](node *AVLNode[T]) *AVLNode[T] {
	newRoot := node.right

	newRoot.left, node.right = node, newRoot.left

	updateHeight(node)
	updateHeight(newRoot)

	return newRoot
}

func calculateBalanceFactor[T cmp.Ordered](node *AVLNode[T]) int32 {
	if node == nil {
		return 0
	}
	return getHeight(node.left) - getHeight(node.right)
}

func getHeight[T cmp.Ordered](node *AVLNode[T]) int32 {
	if node == nil {
		return 0
	}
	return node.height
}

func BinarySearchAVLTree[T cmp.Ordered](node *AVLNode[T], val T) *AVLNode[T] {
	if node == nil {
		return nil
	}
	switch {
	case val < node.Val:
		return BinarySearchAVLTree(node.left, val)
	case val > node.Val:
		return BinarySearchAVLTree(node.right, val)
	case val == node.Val:
		return node
	default:
		return nil
	}
}

func PrintBreadthFirst[T cmp.Ordered](root *AVLNode[T]) {
	if root == nil {
		return
	}
	que := []*AVLNode[T]{root}
	levelLen := 0
	for len(que) > 0 {
		levelLen = len(que)
		for idx := 0; idx < levelLen; idx++ {
			if que[idx] != nil {
				fmt.Printf(",%v:%d", que[idx].Val, que[idx].height)
				que = append(que, que[idx].left)
				que = append(que, que[idx].right)
			} else {
				fmt.Printf(",n")
			}
		}
		fmt.Println()
		que = que[levelLen:]
	}
	fmt.Println()
}

func FormatBreadthFirst[T cmp.Ordered](root *AVLNode[T]) string {
	if root == nil {
		return ""
	}
	sb := strings.Builder{}
	que := []*AVLNode[T]{root}
	levelLen := 0
	for len(que) > 0 {
		levelLen = len(que)
		sb.WriteString("|")
		for idx := 0; idx < levelLen; idx++ {
			if idx > 0 {
				sb.WriteByte(',')
			}
			if que[idx] != nil {
				sb.WriteString(fmt.Sprintf("%v:%d", que[idx].Val, que[idx].height))
				que = append(que, que[idx].left)
				que = append(que, que[idx].right)
			} else {
				sb.WriteByte('n')
			}
		}
		que = que[levelLen:]
	}

	return sb.String()[1:]
}
