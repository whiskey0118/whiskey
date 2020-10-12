package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return 0
	}
	if head.Next == nil {
		return head.Val
	}
	dummy := &ListNode{Next: head}
	p, q := dummy, dummy
	for i := 0; i < k+1; i++ {
		q = q.Next
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}
	return p.Next.Val
}

//自己编写的
func MykthToLast(head *ListNode, k int) int {
	a := 0
	b := 0
	if head == nil {
		return 0
	}
	if head.Next == nil {
		return head.Val
	}
	tmp := &ListNode{Next: head}
	tmp1 := &ListNode{Next: head}
	b = k + 1
	for head.Next == nil {
		tmp = tmp.Next
		b++
		tmp1 = tmp1.Next
		a++
	}
	return tmp1.Val
}
