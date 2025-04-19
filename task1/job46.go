package task1

/*
*
46. 全排列：给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。可以使用回溯算法，定义一个函数来进行递归操作，在函数中通过交换数组元素的位置来生成不同的排列，使用 for 循环遍历数组，每次选择一个元素作为当前排列的第一个元素，然后递归调用函数处理剩余的元素。
*/
func GetAllArrange(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, len(nums))
	backtrack(&state, &nums, &selected, &res)
	return res
}

/* 回溯算法：全排列 */
func backtrack(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	// 当状态长度等于元素数量时，记录解
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
	}
	// 遍历所有选择
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		// 剪枝：不允许重复选择元素
		if !(*selected)[i] {
			// 尝试：做出选择，更新状态
			(*selected)[i] = true
			*state = append(*state, choice)
			// 进行下一轮选择
			backtrack(state, choices, selected, res)
			// 回退：撤销选择，恢复到之前的状态
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}
