package sort

import (
	"fmt"
)

func HeapSort(list []int64) []int64 {
	for i := 0; i < len(list)-2; i++ {
		nums := list[i:]
		for node := len(nums)/2 - 1; node >= 0; node-- {
			index := node
			for index <= len(nums)/2-1 {
				child := index*2 + 1
				if child+1 < len(nums) && nums[child+1] < nums[child] {
					child++
				}
				if nums[child] > nums[index] {
					break
				} else {
					nums[index], nums[child] = nums[child], nums[index]
				}
				index = child
			}
		}
		fmt.Println(list)
	}
	return list
}
