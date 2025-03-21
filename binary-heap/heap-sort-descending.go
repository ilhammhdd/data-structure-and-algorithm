package binaryHeap

import (
	"cmp"
)

func HeapSortDescending[T cmp.Ordered](source []T) {
	buildMinHeap(source)
	sortDescending(source)
}

func buildMinHeap[T cmp.Ordered](source []T) {
	if len(source) <= 1 {
		return
	}
	key := 0
	parentKey := 0
	heapLen := 0
	for heapLen < len(source) {
		key = heapLen
		parentKey = (key - 1) / 2
		for key > 0 && parentKey >= 0 && source[key] < source[parentKey] {
			source[key], source[parentKey] = source[parentKey], source[key]
			key = parentKey
			parentKey = (key - 1) / 2
		}
		heapLen++
	}
}

func sortDescending[T cmp.Ordered](source []T) {
	if len(source) <= 1 {
		return
	}

	sortedHead := len(source) - 1
	key := 0
	leftKey := 0
	rightKey := 0
	smallestChildKey := 0
	for sortedHead > 0 {
		source[sortedHead], source[0] = source[0], source[sortedHead]
		key = 0

		for key < sortedHead {
			leftKey = key*2 + 1
			rightKey = key*2 + 2
			smallestChildKey = leftKey
			if rightKey < sortedHead && source[rightKey] < source[smallestChildKey] {
				smallestChildKey = rightKey
			}
			if smallestChildKey < sortedHead && source[smallestChildKey] < source[key] {
				source[smallestChildKey], source[key] = source[key], source[smallestChildKey]
			}
			key = smallestChildKey
		}

		sortedHead--
	}
}
