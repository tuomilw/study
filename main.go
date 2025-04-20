package main

import (
	"awesomeProject/task1"
	"awesomeProject/task2"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	/**
	Job136
	*/
	nums1 := []int{6, 2, 3, 5, 5, 3, 2}
	findValue1 := task1.FindOneTimesElement(nums1)
	fmt.Println("FindOneTimesElement:", findValue1)
	/**
	Job198
	*/
	nums2 := []int{5, 8, 10, 22, 68, 21, 18, 31}
	findValue2 := task1.FindRoomAtMostCash(nums2)
	fmt.Println("FindRoomAtMostCash:", findValue2)
	/**
	Job21
	*/
	nums3 := []int{2, 5, 3, 8, 6}
	nums4 := []int{1, 4, 9, 7}
	resultNums := task1.MergeTwoLinkTables(nums3, nums4)
	for resultNums != nil {
		fmt.Print(resultNums.Num)
		if resultNums.Next != nil {
			fmt.Print("->")
		}
		resultNums = resultNums.Next
	}
	fmt.Println("")
	/**
	Job46
	*/
	nums5 := []int{1, 2, 3}
	resultAllArrange := task1.GetAllArrange(nums5)
	fmt.Println("GetAllArrange:", resultAllArrange)
	/**
	Job344
	*/
	strs := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("ReversalStr:", task1.ReversalStr(strs))
	/**
	Job69
	*/
	fmt.Println("GetSqrt:", task1.GetSqrt(82))
	/**
	Job26
	*/
	fmt.Println("RemoveDuplicates:", task1.RemoveDuplicates([]int{1, 2, 2, 4, 3, 6, 8, 6, 2, 9}))
	/**
	Job56
	*/
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("Intervals:", task1.MergeIntervals(intervals))

	pointnum := 10
	task2.PointerTask1(&pointnum)
	fmt.Println(pointnum)

	pointnums := []int{1, 2, 3, 4, 5}
	task2.PointerTask2(&pointnums)
	fmt.Println(pointnums)

	task2.RunGoroutineTask1()
}
