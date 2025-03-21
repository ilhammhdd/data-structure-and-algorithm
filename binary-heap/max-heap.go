package binaryHeap

import (
	"cmp"
)

type MaxHeap[T cmp.Ordered] []T

func (heap *MaxHeap[T]) Swap(keyA, keyB int) {
	if heap == nil || keyA >= len(*heap) || keyB >= len(*heap) {
		return
	}
	(*heap)[keyA], (*heap)[keyB] = (*heap)[keyB], (*heap)[keyA]
}

func (heap *MaxHeap[T]) ParentKey(key int) int {
	return (key - 1) / 2
}

func (heap *MaxHeap[T]) LeftChildKey(key int) int {
	return 2*key + 1
}

func (heap *MaxHeap[T]) RightChildKey(key int) int {
	return 2*key + 2
}

func (heap *MaxHeap[T]) Insert(val T) {
	if heap == nil {
		return
	}
	temp := append(*heap, val)
	key := len(temp) - 1
	parentKey := heap.ParentKey(key)
	for key > 0 && temp[key] > temp[parentKey] {
		temp.Swap(key, parentKey)
		key = parentKey
		parentKey = heap.ParentKey(key)
	}
	*heap = temp
}

func (heap *MaxHeap[T]) Delete(key int) {
	if heap == nil {
		return
	}

	outerLeafKey := len(*heap) - 1
	heap.Swap(key, outerLeafKey)
	if len(*heap)-1 >= 0 {
		*heap = (*heap)[:len(*heap)-1]
	}
	heap.Heapify(key)
}

func (heap *MaxHeap[T]) GetMax() T {
	var zero T
	if heap == nil {
		return zero
	}
	return (*heap)[0]
}

func (heap *MaxHeap[T]) ExtractMax() T {
	var zero T
	if heap == nil || len(*heap) < 1 {
		return zero
	}

	top := heap.GetMax()
	(*heap)[0] = (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]
	heap.Heapify(0)

	return top
}

func (heap *MaxHeap[T]) Heapify(key int) {
	if heap == nil {
		return
	}

	leftKey := heap.LeftChildKey(key)
	rightKey := heap.RightChildKey(key)

	for (leftKey < len(*heap) && (*heap)[key] < (*heap)[leftKey]) ||
		(rightKey < len(*heap) && (*heap)[key] < (*heap)[rightKey]) {
		if leftKey < len(*heap) {
			heap.Swap(key, leftKey)
			key = leftKey
			leftKey = heap.LeftChildKey(key)
			rightKey = heap.RightChildKey(key)
		} else if rightKey < len(*heap) {
			heap.Swap(key, rightKey)
			key = rightKey
			leftKey = heap.LeftChildKey(key)
			rightKey = heap.RightChildKey(key)
		}
	}
}
