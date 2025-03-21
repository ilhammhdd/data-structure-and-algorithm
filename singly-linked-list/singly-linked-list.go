package singlyLinkedList

type ListNode[T comparable] struct {
	Val  T
	Next *ListNode[T]
}

func Append[T comparable](head, node *ListNode[T]) *ListNode[T] {
	if head == nil {
		return node
	}
	curr := head
	for curr != nil && curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
	return head
}

func Prepend[T comparable](head, node *ListNode[T]) *ListNode[T] {
	if head == nil {
		return node
	}
	node.Next = head
	head = node
	return head
}

func Remove[T comparable](head *ListNode[T], val T) *ListNode[T] {
	var prev, curr *ListNode[T] = nil, head
	for curr != nil && curr.Val != val {
		prev, curr = curr, curr.Next
	}
	return removeNode(head, prev, curr)
}

func RemoveNode[T comparable](head, node *ListNode[T]) *ListNode[T] {
	var prev, curr *ListNode[T] = nil, head
	for curr != nil && curr != node {
		prev, curr = curr, curr.Next
	}
	return removeNode(head, prev, curr)
}

func removeNode[T comparable](head, prev, node *ListNode[T]) *ListNode[T] {
	if prev == nil && node == head {
		head.Next, head = nil, head.Next
		return head
	}
	prev.Next = node.Next
	node.Next = nil
	return head
}

func Find[T comparable](head *ListNode[T], val T) *ListNode[T] {
	curr := head
	for curr != nil && curr.Val != val {
		curr = curr.Next
	}
	return curr
}

func ToSlice[T comparable](head *ListNode[T]) []T {
	var result []T
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}
