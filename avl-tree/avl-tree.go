package avlTree

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
	node = balancing(node)

	return node
}

func DeleteAVLNode[T cmp.Ordered](node *AVLNode[T], val T) *AVLNode[T] {
	if node == nil {
		return nil
	}
	switch {
	case val < node.Val:
		leftResult := DeleteAVLNode(node.left, val)
		node.left = leftResult
	case val > node.Val:
		rightResult := DeleteAVLNode(node.right, val)
		node.right = rightResult
	}

	if val != node.Val {
		updateHeight(node)
		node = balancing(node)
		return node
	}

	replacement, parent := findReplacementWithParent(node)
	if replacement == nil {
		return nil
	}
	setLeftIfNotEqual(replacement, node)
	setRightIfNotEqual(replacement, node)

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

func balancing[T cmp.Ordered](node *AVLNode[T]) (newRoot *AVLNode[T]) {
	if node == nil {
		return node
	}
	balanceFactor := calculateBalanceFactor(node)

	if balanceFactor > 1 {
		if calculateBalanceFactor(node.left) >= 0 {
			return rotateRight(node)
		} else {
			node.left = rotateLeft(node.left)
			return rotateRight(node)
		}
	} else if balanceFactor < -1 {
		if calculateBalanceFactor(node.right) <= 0 {
			return rotateLeft(node)
		} else {
			node.right = rotateRight(node.right)
			return rotateLeft(node)
		}
	}
	return node
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

func setLeftIfNotEqual[T cmp.Ordered](target, source *AVLNode[T]) {
	if target.left != nil && source.left != nil && target.left.Val != source.left.Val ||
		target.left == nil && source.left != nil && target.Val != source.left.Val {
		target.left = source.left
	}
}

func setRightIfNotEqual[T cmp.Ordered](target, source *AVLNode[T]) {
	if target.right != nil && source.right != nil && target.right.Val != source.right.Val ||
		target.right == nil && source.right != nil && target.Val != source.right.Val {
		target.right = source.right
	}
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

func updateHeight[T cmp.Ordered](node *AVLNode[T]) {
	if node == nil {
		return
	}
	node.height = 1 + max(getHeight(node.left), getHeight(node.right))
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

func SerializeBreadthFirst[T cmp.Ordered](root *AVLNode[T]) string {
	if root == nil {
		return ""
	}
	sb := strings.Builder{}
	que := []*AVLNode[T]{root}
	levelLen := 0
	for len(que) > 0 {
		levelLen = len(que)
		sb.WriteString("|")
		for idx := range levelLen {
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
