package task1

/*
*
69. x 的平方根：实现 int sqrt(int x) 函数。计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。可以使用二分查找法来解决，定义左右边界 left 和 right，然后通过 while 循环不断更新中间值 mid，直到找到满足条件的平方根或者确定不存在精确的平方根。
*/
func GetSqrt(num int) int {
	if num < 0 {
		return 0
	}
	start := 1
	end := num
	result := 0
	for start <= end {
		middle := start + (end-start)/2
		middleValue := middle
		if middleValue <= num/middleValue && middleValue*middleValue <= num {
			result = middle
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return result
}
