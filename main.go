package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	// 遍历链表，两个指针
	//声明三个节点，分别代表前中后的节点
	a, b, c := head, head.Next, ListNode{0, nil}
	for b.Next != nil { // 遍历所有的节点
		c.Val = glc(a.Val, b.Val)
		// 插入
		// a.Next = &c
		// c.Next = b
		// a = b
		// b = b.Next
		fmt.Println(c)
	}
	return head
}

// 求最大公约数
func glc(m, n int) int {
	if m < n {
		m, n = n, m
	}
	for n != 0 {
		m, n = n, m%n
	}
	return m
}
