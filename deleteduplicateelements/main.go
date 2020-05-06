package main

import (
	"fmt"
)

type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

func removeNode(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	if l.Val == l.Next.Val {
		l = removeNode(l.Next)
	} else {
		l.Next = removeNode(l.Next)
	}
	return l
}

func showList(l *ListNode) []int {
	list := []int{}
	if l == nil {
		fmt.Println("l", l.Val)
		list = append(list, l.Val)
	}
	for {
		list = append(list, l.Val)
		if l.Next != nil {
			l = l.Next
		} else {
			break
		}
	}
	return list
}
func main() {
	listnodes := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 1, Next: &ListNode{
				Val: 2, Next: &ListNode{
					Val: 3, Next: &ListNode{
						Val: 3, Next: nil,
					},
				},
			},
		},
	}
	fmt.Println("after:", showList(listnodes))
	nodes := removeNode(listnodes)
	fmt.Println("before:", showList(nodes))
}
