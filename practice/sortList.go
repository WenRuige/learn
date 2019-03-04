package main

import (
	"fmt"
	"sort"
)

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func sortList(head *ListNode2) *ListNode2 {

	intVal := []int{}

	for head != nil {
		intVal = append(intVal, head.Val)
		head = head.Next
	}

	sort.Ints(intVal)
	newHead := &ListNode2{Val: 0}
	newnewHead := newHead

	for i := 0; i < len(intVal); i++ {
		newHead.Next = &ListNode2{Val: intVal[i]}
		newHead = newHead.Next
	}
	return newnewHead.Next
}

func main() {

	head := new(ListNode2)
	head.Val = 3
	head.Next = new(ListNode2)
	head.Next.Val = 4
	head.Next.Next = new(ListNode2)
	head.Next.Next.Val = 1
	result := sortList(head)

	finalResult := []int{}
	for result != nil {
		finalResult = append(finalResult, result.Val)
		result = result.Next
	}
	fmt.Printf("finalResult=%+v", finalResult)
}
