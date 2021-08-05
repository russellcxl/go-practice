package main

import "fmt"

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{9, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{9, nil}}}}

	//l3 := &ListNode{0, nil}
	//l4 := &ListNode{0, nil}

	finalNode := addTwoNumbers(l1, l2)
	fmt.Println(finalNode.Val, finalNode.Next.Val, finalNode.Next.Next.Val, finalNode.Next.Next.Next.Val, finalNode.Next.Next.Next.Next.Val)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := l1
	n2 := l2
	parentNode := new(ListNode)
	var childNode *ListNode

	var needPlusOne bool

	for n1 != nil || n2 != nil || needPlusOne {

		var value1, value2 int

		if n1 != nil {
			value1 = n1.Val
			// if next node exists, set node 1 = next
			if n1.Next != nil {
				n1 = n1.Next
			} else {
				n1 = nil
			}
		}

		if n2 != nil {
			value2 = n2.Val
			if n2.Next != nil {
				n2 = n2.Next
			} else {
				n2 = nil
			}
		}

		// need to handle case for if values 1 and 2 are 0, add one more node with Val=1

		sum := value1 + value2

		if needPlusOne {
			sum += 1
		}

		if sum >= 10 {
			sum = sum % 10
			needPlusOne = true
		} else {
			needPlusOne = false
		}

		fmt.Println(value1, value2, sum)

		// if childNode is nil, fill up parentNode first
		if childNode == nil {
			fmt.Println("storing in parent node:", sum)
			parentNode.Val = sum
			if n1 != nil || n2 != nil || needPlusOne {
				parentNode.Next = new(ListNode)
				childNode = parentNode.Next
			}
		} else {
			fmt.Println("storing in child node:", sum)
			childNode.Val = sum
			if n1 != nil || n2 != nil || needPlusOne {
				childNode.Next = new(ListNode)
				childNode = childNode.Next
			}
		}

	}

	return parentNode
}
