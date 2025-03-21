package binaryHeap

import (
	"cmp"
)

func HeapSortAscending[T cmp.Ordered](source []T) {
	buildMaxHeap(source)
	sortAscending(source)
}

func buildMaxHeap[T cmp.Ordered](source []T) {
	if len(source) <= 1 {
		return
	}
	key := 0
	parentKey := 0
	heapLen := 0
	for heapLen < len(source) {
		key = heapLen
		parentKey = (key - 1) / 2
		for key > 0 && parentKey >= 0 && source[key] > source[parentKey] {
			source[key], source[parentKey] = source[parentKey], source[key]
			key = parentKey
			parentKey = (key - 1) / 2
		}
		heapLen++
	}
}

func sortAscending[T cmp.Ordered](source []T) {
	if len(source) <= 1 {
		return
	}

	sortedHead := len(source) - 1
	key := 0
	leftKey := 0
	rightKey := 0
	biggestChildKey := 0
	for sortedHead > 0 {
		source[sortedHead], source[0] = source[0], source[sortedHead]
		key = 0

		for key < sortedHead {
			leftKey = key*2 + 1
			rightKey = key*2 + 2
			biggestChildKey = leftKey
			if rightKey < sortedHead && source[rightKey] > source[biggestChildKey] {
				biggestChildKey = rightKey
			}
			if biggestChildKey < sortedHead && source[biggestChildKey] > source[key] {
				source[biggestChildKey], source[key] = source[key], source[biggestChildKey]
			}
			key = biggestChildKey
		}

		sortedHead--
	}
}
