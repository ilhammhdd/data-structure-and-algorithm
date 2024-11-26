package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAVLNode_BalanceLeftLeft(t *testing.T) {
	//     10
	//   3
	// 1
	givenVals := []int{10, 3, 1}
	var givenRoot *AVLNode[int]
	for idx := range givenVals {
		givenRoot = InsertAVLNode(givenRoot, givenVals[idx])
	}

	//   3
	// 1  10
	const expected = "3:2|1:1,10:1|n,n,n,n"

	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestInsertAVLNode_BalanceLeftRight(t *testing.T) {
	//   10
	// 3
	//  4
	givenVals := []int{10, 3, 4}
	var givenRoot *AVLNode[int]
	for idx := range givenVals {
		givenRoot = InsertAVLNode(givenRoot, givenVals[idx])
	}

	//   4
	// 3  10
	const expected = "4:2|3:1,10:1|n,n,n,n"

	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestInsertAVLNode_BalanceRightRight(t *testing.T) {
	// 3
	//  4
	//   5
	givenVals := []int{3, 4, 5}
	var givenRoot *AVLNode[int]
	for idx := range givenVals {
		givenRoot = InsertAVLNode(givenRoot, givenVals[idx])
	}

	//  4
	// 3 5
	const expected = "4:2|3:1,5:1|n,n,n,n"

	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestInsertAVLNode_BalanceRightLeft(t *testing.T) {
	// 3
	//   5
	//  4
	givenVals := []int{3, 5, 4}
	var givenRoot *AVLNode[int]
	for idx := range givenVals {
		givenRoot = InsertAVLNode(givenRoot, givenVals[idx])
	}

	//  4
	// 3 5
	const expected = "4:2|3:1,5:1|n,n,n,n"

	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestFindReplacement_Predecessor(t *testing.T) {
	//   8
	// 5   10
	givenRoot := InsertAVLNode(nil, 8)
	eight := givenRoot
	givenRoot = InsertAVLNode(givenRoot, 5)
	InsertAVLNode(givenRoot, 10)

	const expected = 5
	actual, actualParent := findReplacementWithParent(eight)

	assert.NotNil(t, actual)
	assert.Nil(t, actualParent)
	assert.Equal(t, expected, actual.Val)
}

func TestFindReplacement_Successor(t *testing.T) {
	//   8
	//     10
	givenRoot := InsertAVLNode(nil, 8)
	eight := givenRoot
	InsertAVLNode(givenRoot, 10)

	const expected = 10
	actual, actualParent := findReplacementWithParent(eight)

	assert.NotNil(t, actual)
	assert.Nil(t, actualParent)
	assert.Equal(t, expected, actual.Val)
}

func TestFindReplacement_NoChildren(t *testing.T) {
	//   8
	givenRoot := InsertAVLNode(nil, 8)

	actual, actualParent := findReplacementWithParent(givenRoot)

	assert.Nil(t, actual)
	assert.Nil(t, actualParent)
}

func TestFindReplacement_InOrderPredecessor(t *testing.T) {
	//   8
	// 5   10
	//  6
	givenRoot := &AVLNode[int]{Val: 8}
	five := &AVLNode[int]{Val: 5}
	ten := &AVLNode[int]{Val: 10}
	six := &AVLNode[int]{Val: 6}
	five.right = six
	givenRoot.left, givenRoot.right = five, ten

	const expected, expectedParent = 6, 5
	actual, actualParent := findReplacementWithParent(givenRoot)

	assert.NotNil(t, actual)
	assert.NotNil(t, actualParent)
	assert.Equal(t, expected, actual.Val)
	assert.Equal(t, expectedParent, actualParent.Val)
}

// NOTE: the tree in not balanced, the purpose of this case is
// just to demonstrate using successor as replacement
func TestFindReplacement_InOrderSuccessor(t *testing.T) {
	//   8
	//     10
	//    9
	givenRoot := &AVLNode[int]{Val: 8}
	ten := &AVLNode[int]{Val: 10}
	nine := &AVLNode[int]{Val: 9}
	ten.left = nine
	givenRoot.right = ten

	const expected, expectedParent = 9, 10
	actual, actualParent := findReplacementWithParent(givenRoot)

	assert.NotNil(t, actual)
	assert.NotNil(t, actualParent)
	assert.Equal(t, expected, actual.Val)
	assert.Equal(t, expectedParent, actualParent.Val)
}

func TestDeleteAVLNode_RootWithChildren(t *testing.T) {
	//   8
	// 5   10
	//  6
	givenRoot := InsertAVLNode(nil, 8)
	givenRoot = InsertAVLNode(givenRoot, 5)
	givenRoot = InsertAVLNode(givenRoot, 10)
	givenRoot = InsertAVLNode(givenRoot, 6)

	//   6
	// 5   10
	expected := "6:2|5:1,10:1|n,n,n,n"

	givenRoot = DeleteAVLNode(givenRoot, 8)
	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestDeleteAVLNode_RootLeftReplacement(t *testing.T) {
	//   8
	// 5   10
	givenRoot := InsertAVLNode(nil, 8)
	givenRoot = InsertAVLNode(givenRoot, 5)
	givenRoot = InsertAVLNode(givenRoot, 10)

	//   5
	//     10
	expected := "5:2|n,10:1|n,n"

	givenRoot = DeleteAVLNode(givenRoot, 8)
	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestDeleteAVLNode_RootRightReplacement(t *testing.T) {
	//   8
	//     10
	givenRoot := InsertAVLNode(nil, 8)
	givenRoot = InsertAVLNode(givenRoot, 10)

	//   10
	expected := "10:1|n,n"

	givenRoot = DeleteAVLNode(givenRoot, 8)
	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestDeleteAVLNode_RootWithoutChildren(t *testing.T) {
	//   8
	givenRoot := InsertAVLNode(nil, 8)

	expected := ""

	givenRoot = DeleteAVLNode(givenRoot, 8)
	assert.Nil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestDeleteAVLNode_NonRootLeftNoReplacement(t *testing.T) {
	//   8
	// 5   10
	givenRoot := InsertAVLNode(nil, 8)
	givenRoot = InsertAVLNode(givenRoot, 5)
	givenRoot = InsertAVLNode(givenRoot, 10)

	//   8
	//     10
	expected := "8:2|n,10:1|n,n"

	givenRoot = DeleteAVLNode(givenRoot, 5)
	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestDeleteAVLNode_NonRootRightNoReplacement(t *testing.T) {
	//   8
	// 5   10
	givenRoot := InsertAVLNode(nil, 8)
	givenRoot = InsertAVLNode(givenRoot, 5)
	givenRoot = InsertAVLNode(givenRoot, 10)

	//   8
	// 5
	expected := "8:2|5:1,n|n,n"

	givenRoot = DeleteAVLNode(givenRoot, 10)
	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestDeleteAVLNode_NonRootDeepLeft(t *testing.T) {
	//       8
	//   5      10
	// 4   7   9
	//    6
	givenRoot := InsertAVLNode(nil, 8)
	givenRoot = InsertAVLNode(givenRoot, 5)
	givenRoot = InsertAVLNode(givenRoot, 10)
	givenRoot = InsertAVLNode(givenRoot, 4)
	givenRoot = InsertAVLNode(givenRoot, 7)
	givenRoot = InsertAVLNode(givenRoot, 9)
	givenRoot = InsertAVLNode(givenRoot, 6)

	//       8
	//   6      10
	// 4   7   9
	expected := "8:3|6:2,10:2|4:1,7:1,9:1,n|n,n,n,n,n,n"

	givenRoot = DeleteAVLNode(givenRoot, 5)
	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestDeleteAVLNode_NonRootLeftLeftBalancing(t *testing.T) {
	//        8
	//    6      10
	//  4   7   9
	// 3
	givenRoot := InsertAVLNode(nil, 8)
	givenRoot = InsertAVLNode(givenRoot, 6)
	givenRoot = InsertAVLNode(givenRoot, 10)
	givenRoot = InsertAVLNode(givenRoot, 4)
	givenRoot = InsertAVLNode(givenRoot, 7)
	givenRoot = InsertAVLNode(givenRoot, 9)
	givenRoot = InsertAVLNode(givenRoot, 3)

	//        8
	//    4      10
	//  3   6   9
	expected := "8:3|4:2,10:2|3:1,6:1,9:1,n|n,n,n,n,n,n"

	givenRoot = DeleteAVLNode(givenRoot, 7)

	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestDeleteAVLNode_NonRootLeftRightBalancing(t *testing.T) {
	//        8
	//    6      10
	//  4   7   9
	//   5
	givenRoot := InsertAVLNode(nil, 8)
	givenRoot = InsertAVLNode(givenRoot, 6)
	givenRoot = InsertAVLNode(givenRoot, 10)
	givenRoot = InsertAVLNode(givenRoot, 4)
	givenRoot = InsertAVLNode(givenRoot, 7)
	givenRoot = InsertAVLNode(givenRoot, 9)
	givenRoot = InsertAVLNode(givenRoot, 5)

	//        8
	//    5      10
	//  4   6   9
	expected := "8:3|5:2,10:2|4:1,6:1,9:1,n|n,n,n,n,n,n"

	givenRoot = DeleteAVLNode(givenRoot, 7)

	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestDeleteAVLNode_NonRootRightRightBalancing(t *testing.T) {
	//        8
	//    6       10
	//  4       9    12
	//                 14
	givenRoot := InsertAVLNode(nil, 8)
	givenRoot = InsertAVLNode(givenRoot, 6)
	givenRoot = InsertAVLNode(givenRoot, 10)
	givenRoot = InsertAVLNode(givenRoot, 4)
	givenRoot = InsertAVLNode(givenRoot, 9)
	givenRoot = InsertAVLNode(givenRoot, 12)
	givenRoot = InsertAVLNode(givenRoot, 14)

	//        8
	//    6       12
	//  4      10    14
	expected := "8:3|6:2,12:2|4:1,n,10:1,14:1|n,n,n,n,n,n"

	givenRoot = DeleteAVLNode(givenRoot, 9)

	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}

func TestDeleteAVLNode_NonRootRightLeftBalancing(t *testing.T) {
	//         8
	//    6        10
	//  4       9     12
	//              11
	givenRoot := InsertAVLNode(nil, 8)
	givenRoot = InsertAVLNode(givenRoot, 6)
	givenRoot = InsertAVLNode(givenRoot, 10)
	givenRoot = InsertAVLNode(givenRoot, 4)
	givenRoot = InsertAVLNode(givenRoot, 9)
	givenRoot = InsertAVLNode(givenRoot, 12)
	givenRoot = InsertAVLNode(givenRoot, 11)

	//        8
	//    6       11
	//  4      10    12
	expected := "8:3|6:2,11:2|4:1,n,10:1,12:1|n,n,n,n,n,n"

	givenRoot = DeleteAVLNode(givenRoot, 9)

	assert.NotNil(t, givenRoot)
	assert.Equal(t, expected, SerializeBreadthFirst(givenRoot))
}
