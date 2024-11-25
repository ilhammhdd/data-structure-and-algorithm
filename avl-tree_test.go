package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: need extensive test cases for insertion

func TestFindReplacement_Predecessor(t *testing.T) {
	//
	//   8
	// 5   10
	//
	root := InsertAVLNode(nil, 8)
	eight := root
	root = InsertAVLNode(root, 5)
	InsertAVLNode(root, 10)

	const expected = 5
	actual := findReplacement(eight)

	assert.NotNil(t, actual)
	assert.Equal(t, expected, actual.Val)
}

func TestFindReplacement_Successor(t *testing.T) {
	//
	//   8
	//     10
	//
	root := InsertAVLNode(nil, 8)
	eight := root
	InsertAVLNode(root, 10)

	const expected = 10
	actual := findReplacement(eight)

	assert.NotNil(t, actual)
	assert.Equal(t, expected, actual.Val)
}

func TestFindReplacement_NoChildren(t *testing.T) {
	//
	//   8
	//
	root := InsertAVLNode(nil, 8)

	actual := findReplacement(root)

	assert.Nil(t, actual)
}

func TestFindReplacement_NotFound(t *testing.T) {
	//
	//   8
	//
	root := InsertAVLNode(nil, 8)

	actual := findReplacement(root)

	assert.Nil(t, actual)
}

func TestFindReplacement_InOrderPredecessor(t *testing.T) {
	//
	//   8
	// 5   10
	//  6
	//
	eight := &AVLNode[int]{Val: 8}
	five := &AVLNode[int]{Val: 5}
	ten := &AVLNode[int]{Val: 10}
	six := &AVLNode[int]{Val: 6}
	five.right = six
	eight.left, eight.right = five, ten

	const expected = 6
	actual := findReplacement(eight)

	assert.Equal(t, expected, actual.Val)
}

// NOTE: the tree in not balanced, the purpose of this case is just to demonstrate if the successor as replacement
func TestFindReplacement_InOrderSuccessor(t *testing.T) {
	//   8
	//     10
	//    9
	eight := &AVLNode[int]{Val: 8}
	ten := &AVLNode[int]{Val: 10}
	nine := &AVLNode[int]{Val: 9}
	ten.left = nine
	eight.right = ten

	const expected = 9
	actual := findReplacement(eight)

	assert.Equal(t, expected, actual.Val)
}

func TestDeleteAVLNode_RootWithChildren(t *testing.T) {
	//   8
	// 5   10
	//  6
	root := InsertAVLNode(nil, 8)
	root = InsertAVLNode(root, 5)
	root = InsertAVLNode(root, 10)
	root = InsertAVLNode(root, 6)

	//   6
	// 5   10
	expected := "6:2|5:1,10:1|n,n,n,n"

	root = DeleteAVLNode(root, 8)
	assert.NotNil(t, root)
	assert.Equal(t, expected, FormatBreadthFirst(root))
}

func TestDeleteAVLNode_RootLeftReplacement(t *testing.T) {
	//   8
	// 5   10
	root := InsertAVLNode(nil, 8)
	root = InsertAVLNode(root, 5)
	root = InsertAVLNode(root, 10)

	//   5
	//     10
	expected := "5:2|n,10:1|n,n"

	root = DeleteAVLNode(root, 8)
	assert.NotNil(t, root)
	assert.Equal(t, expected, FormatBreadthFirst(root))
}

func TestDeleteAVLNode_RootRightReplacement(t *testing.T) {
	//   8
	//     10
	root := InsertAVLNode(nil, 8)
	root = InsertAVLNode(root, 10)

	//   10
	expected := "10:1|n,n"

	root = DeleteAVLNode(root, 8)
	assert.NotNil(t, root)
	assert.Equal(t, expected, FormatBreadthFirst(root))
}

func TestDeleteAVLNode_RootWithoutChildren(t *testing.T) {
	//   8
	root := InsertAVLNode(nil, 8)

	expected := ""

	root = DeleteAVLNode(root, 8)
	assert.Nil(t, root)
	assert.Equal(t, expected, FormatBreadthFirst(root))
}

func TestDeleteAVLNode_NonRootLeftNoReplacement(t *testing.T) {
	//   8
	// 5   10
	root := InsertAVLNode(nil, 8)
	root = InsertAVLNode(root, 5)
	root = InsertAVLNode(root, 10)

	//   8
	//     10
	expected := "8:2|n,10:1|n,n"

	root = DeleteAVLNode(root, 5)
	assert.NotNil(t, root)
	assert.Equal(t, expected, FormatBreadthFirst(root))
}

func TestDeleteAVLNode_NonRootRightNoReplacement(t *testing.T) {
	//   8
	// 5   10
	root := InsertAVLNode(nil, 8)
	root = InsertAVLNode(root, 5)
	root = InsertAVLNode(root, 10)

	//   8
	// 5
	expected := "8:2|5:1,n|n,n"

	root = DeleteAVLNode(root, 10)
	assert.NotNil(t, root)
	assert.Equal(t, expected, FormatBreadthFirst(root))
}

func TestDeleteAVLNode_NonRootDeepLeft(t *testing.T) {
	//       8
	//   5      10
	// 4   7   9
	//    6
	root := InsertAVLNode(nil, 8)
	root = InsertAVLNode(root, 5)
	root = InsertAVLNode(root, 10)
	root = InsertAVLNode(root, 4)
	root = InsertAVLNode(root, 7)
	root = InsertAVLNode(root, 9)
	root = InsertAVLNode(root, 6)

	//       8
	//   6      10
	// 4   7   9
	expected := "8:3|6:2,10:2|4:1,7:1,9:1,n|n,n,n,n,n,n"

	root = DeleteAVLNode(root, 5)
	assert.NotNil(t, root)
	assert.Equal(t, expected, FormatBreadthFirst(root))
}

// TODO: need test case for left left, right right, left right and right left balancing after deletion
