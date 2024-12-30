package main

import (
	"cmp"
)

type BSTNode[T cmp.Ordered] struct {
	Val   T
	left  *BSTNode[T]
	right *BSTNode[T]
}

func Insert[T cmp.Ordered](node *BSTNode[T], val T) *BSTNode[T] {
	if node == nil {
		return &BSTNode[T]{Val: val}
	} else if val < node.Val {
		node.left = Insert(node.left, val)
	} else if val > node.Val {
		node.right = Insert(node.right, val)
	}
	return node
}

func Contains[T cmp.Ordered](node *BSTNode[T], val T) bool {
	if node == nil {
		return false
	} else if val < node.Val {
		return Contains(node.left, val)
	} else if val > node.Val {
		return Contains(node.right, val)
	}
	return true
}

func InOrder[T cmp.Ordered](node *BSTNode[T]) []T {
	if node == nil {
		return []T{}
	}
	result := InOrder(node.left)
	result = append(result, node.Val)
	result = append(result, InOrder(node.right)...)
	return result
}

func PreOrder[T cmp.Ordered](node *BSTNode[T]) []T {
	if node == nil {
		return []T{}
	}
	result := []T{node.Val}
	result = append(result, PreOrder(node.left)...)
	result = append(result, PreOrder(node.right)...)
	return result
}

func PostOrder[T cmp.Ordered](node *BSTNode[T]) []T {
	if node == nil {
		return []T{}
	}
	result := PostOrder(node.left)
	result = append(result, PostOrder(node.right)...)
	result = append(result, node.Val)
	return result
}
