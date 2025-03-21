package singlyLinkedList

func MergeSort[T comparable](head *ListNode[T], cmp func(a, b *ListNode[T]) int) *ListNode[T] {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}
	var prev, slow, fast *ListNode[T] = nil, head, head
	for fast != nil && fast.Next != nil {
		prev, slow, fast = slow, slow.Next, fast.Next.Next
	}
	prev.Next = nil
	return merge(MergeSort(head, cmp), MergeSort(slow, cmp), cmp)
}

func merge[T comparable](left, right *ListNode[T], cmp func(a, b *ListNode[T]) int) *ListNode[T] {
	head := &ListNode[T]{}
	curr := head
	for left != nil && right != nil {
		if result := cmp(left, right); result <= 0 {
			curr.Next, left = left, left.Next
		} else {
			curr.Next, right = right, right.Next
		}
		curr = curr.Next
	}
	if left != nil {
		curr.Next = left
	}
	if right != nil {
		curr.Next = right
	}
	return head.Next
}
