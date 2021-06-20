package two_sum

// O(n^2)
func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return res
}

// O(n)
func optimizedTwoSum(nums []int, target int) []int {
	res := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		find := target - nums[i]
		v, found := res[find]
		if found {
			return []int{i, v}
		}
		res[nums[i]] = i
	}
	return []int{}
}
