package task1

/*
*
198. 打家劫舍：你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。给定一个代表每个房屋存放金额的非负整数数组，计算你不触动警报装置的情况下，一夜之内能够偷窃到的最高金额。这道题可以使用动态规划的思想，通过 for 循环遍历数组，利用 if 条件判断来决定是否选择当前房屋进行抢劫，状态转移方程为 dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])。
*/
func FindRoomAtMostCash(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return getMaxNum(nums[0], nums[1])
	} else {
		preNum := nums[0]
		maxNum := getMaxNum(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			tmpNum := getMaxNum(maxNum, preNum+nums[i])
			preNum = maxNum
			maxNum = tmpNum
		}
		return maxNum
	}
}

func getMaxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
