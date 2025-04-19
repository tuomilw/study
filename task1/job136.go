package task1

/*
*
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func FindOneTimesElement(nums []int) int {
	mapObj := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if _, ok := mapObj[num]; ok {
			mapObj[num] = mapObj[num] + 1
		} else {
			mapObj[num] = 1
		}
	}
	for key, value := range mapObj {
		if value == 1 {
			return key
		}
	}
	return -1
}
