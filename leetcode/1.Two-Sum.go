package leetcode

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func TwoSumV1(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == (target - num) {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for index, num := range nums {
		m := target - num
		if _, ok := numMap[m]; ok {
			return []int{numMap[m], index}
		} else {
			numMap[num] = index
		}
	}
	return nil
}
