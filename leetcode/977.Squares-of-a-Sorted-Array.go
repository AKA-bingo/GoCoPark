package leetcode

/*
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.

Example 1:
Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]

Example 2:
Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Note:
1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.
*/

func SortedSquares(A []int) []int {
	res := make([]int, len(A))
	left, right := 0, len(A)-1
	for i := len(A) - 1; i >= 0; i-- {
		if abs(A[left]) > abs(A[right]) {
			res[i] = A[left] * A[left]
			left++
		} else if abs(A[left]) == abs(A[right]) && left != right {
			res[i], res[i-1] = A[left]*A[left], A[right]*A[right]
			left++
			right--
			i--
		} else {
			res[i] = A[right] * A[right]
			right--
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
