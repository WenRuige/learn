package main

type ListNode5 struct {
	val  int
	next *ListNode5
}

func reverse(head *ListNode5) *ListNode5 {

	if head == nil {
		return nil
	}
	dummy := new(ListNode5)
	dummy.next = head

	prev, pcur := new(ListNode5), new(ListNode5)
	prev = dummy.next

	pcur = prev.next

	for pcur != nil {
		prev.next = pcur.next
		pcur.next = dummy.next
		dummy.next = pcur
		pcur = prev.next
	}

	return dummy.next

}
