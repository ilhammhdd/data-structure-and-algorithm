package singlyLinkedList

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	givenN := 100_000
	expected := make([]int32, givenN)
	var actualList *ListNode[int32]
	for i := range givenN {
		num := rand.Int32()
		expected[i] = num
		actualList = Prepend(actualList, &ListNode[int32]{Val: num})
	}
	slices.Sort(expected)
	actualList = MergeSort(actualList, func(a, b *ListNode[int32]) int {
		if a.Val < b.Val {
			return -1
		} else if a.Val > b.Val {
			return 1
		}
		return 0
	})
	actualSlice := ToSlice(actualList)
	assert.Equal(t, expected, actualSlice)
}
