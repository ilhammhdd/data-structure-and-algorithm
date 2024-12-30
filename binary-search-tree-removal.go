package main

import "cmp"

func Remove[T cmp.Ordered](node *BSTNode[T], val T) *BSTNode[T] {
	return remove(node, node, val)
}

func remove[T cmp.Ordered](root, node *BSTNode[T], val T) *BSTNode[T] {
	if node == nil {
		return nil
	} else if val < node.Val {
		node.left = remove(root, node.left, val)
		return node
	} else if val > node.Val {
		node.right = remove(root, node.right, val)
		return node
	}

	isLeftReplacement := true
	replacementParent, replacement := LeftReplacement(node)
	if replacement == nil {
		isLeftReplacement = false
		replacementParent, replacement = RightReplacement(node)
	}
	if replacement == nil {
		return nil
	}

	isDirectChild := replacementParent == node
	node.Val, replacement.Val = replacement.Val, node.Val

	if isDirectChild {
		if isLeftReplacement {
			replacementParent.left = replacement.left
		} else {
			replacementParent.right = replacement.right
		}
		return node
	}

	if isLeftReplacement {
		replacementParent.right = replacement.left
	} else {
		replacementParent.left = replacement.right
	}
	return node
}

func LeftReplacement[T cmp.Ordered](target *BSTNode[T]) (replacementParent, replacement *BSTNode[T]) {
	return leftReplacement(target, target, target.left)
}

func leftReplacement[T cmp.Ordered](target, replacementParent, replacement *BSTNode[T]) (*BSTNode[T], *BSTNode[T]) {
	if replacement == nil {
		return replacementParent, nil
	}
	if replacement.right != nil {
		return leftReplacement(target, replacement, replacement.right)
	}
	return replacementParent, replacement
}

func RightReplacement[T cmp.Ordered](target *BSTNode[T]) (replacementParent, replacement *BSTNode[T]) {
	return rightReplacement(target, target, target.right)
}

func rightReplacement[T cmp.Ordered](target, replacementParent, replacement *BSTNode[T]) (*BSTNode[T], *BSTNode[T]) {
	if replacement == nil {
		return replacementParent, nil
	}
	if replacement.left != nil {
		return rightReplacement(target, replacement, replacement.left)
	}
	return replacementParent, replacement
}
