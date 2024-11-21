package main

import (
	"cmp"
)

type MinHeap[T cmp.Ordered] []T

func (heap *MinHeap[T]) Swap(keyA, keyB int) {
	if heap == nil || keyA >= len(*heap) || keyB >= len(*heap) {
		return
	}
	(*heap)[keyA], (*heap)[keyB] = (*heap)[keyB], (*heap)[keyA]
}

func (heap *MinHeap[T]) ParentKey(key int) int {
	return (key - 1) / 2
}

func (heap *MinHeap[T]) LeftChildKey(key int) int {
	return 2*key + 1
}

func (heap *MinHeap[T]) RightChildKey(key int) int {
	return 2*key + 2
}

func (heap *MinHeap[T]) Insert(val T) {
	if heap == nil {
		return
	}
	temp := append(*heap, val)
	key := len(temp) - 1
	parentKey := heap.ParentKey(key)
	for key > 0 && temp[key] < temp[parentKey] {
		temp.Swap(key, parentKey)
		key = parentKey
		parentKey = heap.ParentKey(key)
	}
	*heap = temp
}

func (heap *MinHeap[T]) Delete(key int) {
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

func (heap *MinHeap[T]) GetMin() T {
	var zero T
	if heap == nil {
		return zero
	}
	return (*heap)[0]
}

func (heap *MinHeap[T]) ExtractMin() T {
	var zero T
	if heap == nil || len(*heap) < 1 {
		return zero
	}

	top := heap.GetMin()
	(*heap)[0] = (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]
	heap.Heapify(0)

	return top
}

func (heap *MinHeap[T]) Heapify(key int) {
	if heap == nil {
		return
	}

	leftKey := heap.LeftChildKey(key)
	rightKey := heap.RightChildKey(key)

	for (leftKey < len(*heap) && (*heap)[key] > (*heap)[leftKey]) ||
		(rightKey < len(*heap) && (*heap)[key] > (*heap)[rightKey]) {
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
