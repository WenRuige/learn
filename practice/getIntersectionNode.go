package main

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

func getIntersectionNode(headA, headB *ListNode3) *ListNode3 {

	// 存入栈
	stackA := []int{}
	stackB := []int{}

	for headA != nil {
		stackA = append(stackA, headA.Val)
	}
	for headB != nil {
		stackB = append(stackB, headB.Val)
	}

	return nil
}

func main() {
	h1 := new(ListNode3)
	h1.Val = 4
	h1.Next = new(ListNode3)
	h1.Next.Val = 1
	h1.Next.Next = new(ListNode3)
	h1.Next.Next.Val = 8
	h1.Next.Next.Next = new(ListNode3)
	h1.Next.Next.Next.Val = 4
	h1.Next.Next.Next.Next = new(ListNode3)
	h1.Next.Next.Next.Next.Val = 5

	h2 := new(ListNode3)
	h2.Val = 5
	h2.Next = new(ListNode3)
	h2.Next.Val = 0
	h2.Next.Next = new(ListNode3)
	h2.Next.Next.Val = 1
	h2.Next.Next.Next = new(ListNode3)
	h2.Next.Next.Next.Val = 8
	h2.Next.Next.Next.Next = new(ListNode3)
	h2.Next.Next.Next.Next.Val = 4
	h2.Next.Next.Next.Next.Next = new(ListNode3)
	h2.Next.Next.Next.Next.Next.Val = 5

	getIntersectionNode(h1, h2)
}
