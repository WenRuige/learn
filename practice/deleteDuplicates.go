package main

/*

定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:
输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3


*/
type ListNode4 struct {
	Val  int
	Next *ListNode4
}

func deleteDuplicates(head *ListNode4) *ListNode4 {

}

func main() {
	l := new(ListNode4)
	l.Val = 1
	l.Next = new(ListNode4)
	l.Next.Val = 1
	l.Next.Next = new(ListNode4)
	l.Next.Next.Val = 1
	l.Next.Next.Next = new(ListNode4)
	l.Next.Next.Next.Val = 2
	l.Next.Next.Next = new(ListNode4)
	l.Next.Next.Next.Val = 2
}
