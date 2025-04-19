package task1

import "fmt"

/*
*
344. 反转字符串：编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。可以使用 for 循环和两个指针，一个指向字符串的开头，一个指向字符串的结尾，然后交换两个指针所指向的字符，直到两个指针相遇。
*/
func ReversalStr(strs []string) []string {
	if len(strs) == 0 {
		fmt.Println("数组为空")
		return make([]string, 0)
	}
	start := 0
	end := len(strs) - 1
	for ; start < end; start++ {
		temp := strs[start]
		strs[start] = strs[end]
		strs[end] = temp
		end--
	}
	return strs
}
