package task1

import "sort"

/*
*
21. 合并两个有序链表：将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。可以定义一个函数，接收两个链表的头节点作为参数，在函数内部使用双指针法，通过比较两个链表节点的值，将较小值的节点添加到新链表中，直到其中一个链表为空，然后将另一个链表剩余的节点添加到新链表中。
*/
func MergeTwoLinkTables(nums1, nums2 []int) *NodeList {
	sort.Ints(nums1)
	sort.Ints(nums2)
	nodeList1 := convert(nums1)
	nodeList2 := convert(nums2)
	result := &NodeList{}
	current := result
	for nodeList1 != nil && nodeList2 != nil {
		if nodeList1.Num < nodeList2.Num {
			current.Next = nodeList1
			nodeList1 = nodeList1.Next
		} else {
			current.Next = nodeList2
			nodeList2 = nodeList2.Next
		}
		current = current.Next
	}
	if nodeList1 != nil {
		current.Next = nodeList1
	} else {
		current.Next = nodeList2
	}
	return result.Next
}

func convert(nums []int) *NodeList {
	result := &NodeList{}
	current := result
	for _, num := range nums {
		current.Next = &NodeList{Num: num}
		current = current.Next
	}
	return result.Next
}

type NodeList struct {
	Num  int
	Next *NodeList
}
