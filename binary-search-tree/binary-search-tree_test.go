package binarySearchTree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTInsert(t *testing.T) {
	//    4
	//  2   5
	// 1 3
	givenToInsert := []int{4, 2, 5, 1, 3}
	expected := []int{1, 2, 3, 4, 5}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	actual := InOrder(root)
	assert.Equal(t, expected, actual)
}

func TestBSTInsert_SkewedLeft(t *testing.T) {
	//     5
	//    4
	//   3
	//  2
	// 1
	givenToInsert := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	actual := InOrder(root)
	assert.Equal(t, expected, actual)
}

func TestBSTInsert_SkewedRight(t *testing.T) {
	// 1
	//  2
	//   3
	//    4
	//     5
	givenToInsert := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	actual := InOrder(root)
	assert.Equal(t, expected, actual)
}

func TestBSTInsert_Balanced(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	actual := InOrder(root)
	assert.Equal(t, expected, actual)
}

func TestBSTContains_ExistsRoot(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	givenContains := 8
	assert.True(t, Contains(root, givenContains))
}

func TestBSTContains_ExistsSubrootLeft(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	givenContains := 6
	assert.True(t, Contains(root, givenContains))
}

func TestBSTContains_ExistsSubrootRight(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	givenContains := 10
	assert.True(t, Contains(root, givenContains))
}

func TestBSTContains_ExistsLeafLeft(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	givenContains := 1
	assert.True(t, Contains(root, givenContains))
}

func TestBSTContains_ExistsLeafRight(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	givenContains := 15
	assert.True(t, Contains(root, givenContains))
}

func TestBSTContains_NoExistsLess(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	givenContains := 0
	assert.False(t, Contains(root, givenContains))
}

func TestBSTContains_NoExistsGreater(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	givenContains := 16
	assert.False(t, Contains(root, givenContains))
}
