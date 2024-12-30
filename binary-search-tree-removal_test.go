package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftReplacement_DirectLeaf(t *testing.T) {
	//         8
	//    4
	// 2     5
	givenToInsert := []int{8, 4, 2, 5}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//    8
	// 2
	//   5
	actualReplacementParent, actualReplacement := LeftReplacement(root.left)
	assert.Equal(t, 4, actualReplacementParent.Val)
	assert.Equal(t, 2, actualReplacement.Val)
}

func TestLeftReplacement_DirectWithLeftChild(t *testing.T) {
	//	        8
	//	   4
	//	2     5
	// 1
	givenToInsert := []int{8, 4, 2, 5, 1}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	   2
	//	1     5
	actualReplacementParent, actualReplacement := LeftReplacement(root.left)
	assert.Equal(t, 4, actualReplacementParent.Val)
	assert.Equal(t, 2, actualReplacement.Val)
}

func TestLeftReplacement_RightLeaf(t *testing.T) {
	//	        8
	//	    4
	//   1
	// 0   2
	givenToInsert := []int{8, 4, 1, 0, 2}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	    2
	//   1
	// 0
	actualReplacementParent, actualReplacement := LeftReplacement(root.left)
	assert.Equal(t, 1, actualReplacementParent.Val)
	assert.Equal(t, 2, actualReplacement.Val)
}

func TestLeftReplacement_RightLeafWithLeftChild(t *testing.T) {
	//	        8
	//	    4
	//   1
	// 0   3
	//    2
	givenToInsert := []int{8, 4, 1, 0, 3, 2}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	    3
	//   1
	// 0   2
	actualReplacementParent, actualReplacement := LeftReplacement(root.left)
	assert.Equal(t, 1, actualReplacementParent.Val)
	assert.Equal(t, 3, actualReplacement.Val)
}

func TestRightReplacement_DirectLeaf(t *testing.T) {
	// 2
	//   5
	//    8
	givenToInsert := []int{2, 5, 8}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//   8
	actualReplacementParent, actualReplacement := RightReplacement(root.right)
	assert.Equal(t, 5, actualReplacementParent.Val)
	assert.Equal(t, 8, actualReplacement.Val)
}

func TestRightReplacement_DirectWithRightChild(t *testing.T) {
	// 2
	//    5
	//      8
	//       9
	givenToInsert := []int{2, 5, 8, 9}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//   8
	//    9
	actualReplacementParent, actualReplacement := RightReplacement(root.right)
	assert.Equal(t, 5, actualReplacementParent.Val)
	assert.Equal(t, 8, actualReplacement.Val)
}

func TestRightReplacement_Leaf(t *testing.T) {
	// 2
	//    5
	//      8
	//     7
	givenToInsert := []int{2, 5, 8, 7}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//    7
	//      8
	actualReplacementParent, actualReplacement := RightReplacement(root.right)
	assert.Equal(t, 8, actualReplacementParent.Val)
	assert.Equal(t, 7, actualReplacement.Val)
}

func TestRightReplacement_WithRightChild(t *testing.T) {
	// 2
	//    5
	//       9
	//     7
	//      8
	givenToInsert := []int{2, 5, 9, 7, 8}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//    7
	//       9
	//     8
	actualReplacementParent, actualReplacement := RightReplacement(root.right)
	assert.Equal(t, 9, actualReplacementParent.Val)
	assert.Equal(t, 7, actualReplacement.Val)
}

func TestRemove_NoExistsLess(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	root = Remove(root, 0)
	expectedInOrder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	assert.Equal(t, expectedInOrder, InOrder(root))
}

func TestRemove_NoExistsGreater(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	root = Remove(root, 16)
	expectedInOrder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	assert.Equal(t, expectedInOrder, InOrder(root))
}

func TestRemove_Root(t *testing.T) {
	// 8
	givenToInsert := []int{8}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//
	root = Remove(root, 8)
	expectedInOrder := []int{}
	assert.Equal(t, expectedInOrder, InOrder(root))
}

func TestRemove_Root_LeftReplacementDirect(t *testing.T) {
	//  8
	// 7
	givenToInsert := []int{8, 7}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 7
	root = Remove(root, 8)
	expectedInOrder := []int{7}
	assert.Equal(t, expectedInOrder, InOrder(root))
}

func TestRemove_Root_RightReplacementDirect(t *testing.T) {
	// 8
	//  9
	givenToInsert := []int{8, 9}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 7
	root = Remove(root, 8)
	expectedInOrder := []int{9}
	assert.Equal(t, expectedInOrder, InOrder(root))
}

func TestRemove_RootLeftLeafReplacement(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5 7   9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//          7
	//     4            12
	//  2     6     10      14
	// 1 3   5     9 11   13  15
	expectedInOrder := []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15}
	expectedPreOrder := []int{7, 4, 2, 1, 3, 6, 5, 12, 10, 9, 11, 14, 13, 15}
	expectedPostOrder := []int{1, 3, 2, 5, 6, 4, 9, 11, 10, 13, 15, 14, 12, 7}
	root = Remove(root, 8)
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RootLeftReplacement(t *testing.T) {
	//          8
	//     4            12
	//  2     6     10      14
	// 1 3   5     9 11   13  15
	givenToInsert := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//          6
	//     4            12
	//  2     5     10      14
	// 1 3         9 11   13  15
	expectedInorder := []int{1, 2, 3, 4, 5, 6, 9, 10, 11, 12, 13, 14, 15}
	expectedPreOrder := []int{6, 4, 2, 1, 3, 5, 12, 10, 9, 11, 14, 13, 15}
	expectedPostOrder := []int{1, 3, 2, 5, 4, 9, 11, 10, 13, 15, 14, 12, 6}
	root = Remove(root, 8)
	assert.Equal(t, expectedInorder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RootRightLeafReplacement(t *testing.T) {
	//   8
	//      12
	//  10      14
	// 9 11   13  15
	givenToInsert := []int{8, 12, 10, 14, 9, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//   9
	//      12
	//  10      14
	//   11   13  15
	expectedInOrder := []int{9, 10, 11, 12, 13, 14, 15}
	expectedPreOrder := []int{9, 12, 10, 11, 14, 13, 15}
	expectedPostOrder := []int{11, 10, 13, 15, 14, 12, 9}
	root = Remove(root, 8)
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RootRightReplacement(t *testing.T) {
	//   8
	//      12
	//  10      14
	//   11   13  15
	givenToInsert := []int{8, 12, 10, 14, 11, 13, 15}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//   10
	//      12
	//  11      14
	//        13  15
	expectedInOrder := []int{10, 11, 12, 13, 14, 15}
	expectedPreOrder := []int{10, 12, 11, 14, 13, 15}
	expectedPostOrder := []int{11, 13, 15, 14, 12, 10}
	root = Remove(root, 8)
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_LeftReplacement_DirectLeaf(t *testing.T) {
	//         8
	//    4
	// 2     5
	givenToInsert := []int{8, 4, 2, 5}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//    8
	// 2
	//   5
	expectedInOrder := []int{2, 5, 8}
	expectedPreOrder := []int{8, 2, 5}
	expectedPostOrder := []int{5, 2, 8}
	root = Remove(root, 4)
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_LeftReplacement_DirectWithLeftChild(t *testing.T) {
	//	        8
	//	   4
	//	2     5
	// 1
	givenToInsert := []int{8, 4, 2, 5, 1}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	   2
	//	1     5
	root = Remove(root, 4)
	expectedInOrder := []int{1, 2, 5, 8}
	expectedPreOrder := []int{8, 2, 1, 5}
	expectedPostOrder := []int{1, 5, 2, 8}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_LeftReplacement_RightLeaf(t *testing.T) {
	//	        8
	//	    4
	//   1
	// 0   2
	givenToInsert := []int{8, 4, 1, 0, 2}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	    2
	//   1
	// 0
	root = Remove(root, 4)
	expectedInOrder := []int{0, 1, 2, 8}
	expectedPreOrder := []int{8, 2, 1, 0}
	expectedPostOrder := []int{0, 1, 2, 8}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_LeftReplacement_RightLeafWithLeftChild(t *testing.T) {
	//	        8
	//	    4
	//   1
	// 0   3
	//    2
	givenToInsert := []int{8, 4, 1, 0, 3, 2}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	//	        8
	//	    3
	//   1
	// 0   2
	root = Remove(root, 4)
	expectedInOrder := []int{0, 1, 2, 3, 8}
	expectedPreOrder := []int{8, 3, 1, 0, 2}
	expectedPostOrder := []int{0, 2, 1, 3, 8}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RightReplacement_DirectLeaf(t *testing.T) {
	// 2
	//   5
	//    8
	givenToInsert := []int{2, 5, 8}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//   8
	root = Remove(root, 5)
	expectedInOrder := []int{2, 8}
	expectedPreOrder := []int{2, 8}
	expectedPostOrder := []int{8, 2}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RightReplacement_DirectWithRightChild(t *testing.T) {
	// 2
	//    5
	//      8
	//       9
	givenToInsert := []int{2, 5, 8, 9}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//   8
	//    9
	root = Remove(root, 5)
	expectedInOrder := []int{2, 8, 9}
	expectedPreOrder := []int{2, 8, 9}
	expectedPostOrder := []int{9, 8, 2}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RightReplacement_Leaf(t *testing.T) {
	// 2
	//    5
	//      8
	//     7
	givenToInsert := []int{2, 5, 8, 7}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//    7
	//      8
	root = Remove(root, 5)
	expectedInOrder := []int{2, 7, 8}
	expectedPreOrder := []int{2, 7, 8}
	expectedPostOrder := []int{8, 7, 2}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}

func TestRemove_RightReplacement_WithRightChild(t *testing.T) {
	// 2
	//    5
	//       9
	//     7
	//      8
	givenToInsert := []int{2, 5, 9, 7, 8}
	var root *BSTNode[int]
	for _, toInsert := range givenToInsert {
		root = Insert(root, toInsert)
	}
	// 2
	//    7
	//       9
	//     8
	root = Remove(root, 5)
	expectedInOrder := []int{2, 7, 8, 9}
	expectedPreOrder := []int{2, 7, 9, 8}
	expectedPostOrder := []int{8, 9, 7, 2}
	assert.Equal(t, expectedInOrder, InOrder(root))
	assert.Equal(t, expectedPreOrder, PreOrder(root))
	assert.Equal(t, expectedPostOrder, PostOrder(root))
}
