package leetcode

/*
A string S of lowercase letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.

Example 1:
Input: S = "ababcbacadefegdehijhklij"
Output: [9,7,8]

Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.

Note:
S will have length in range [1, 500].
S will consist of lowercase letters ('a' to 'z') only.
*/

func PartitionLabels(S string) []int {
	res := make([]int, 0)
	lastMap := make(map[rune]int, 0)
	for index, char := range S {
		lastMap[char] = index
	}

	var count, last int
	for i, char := range S {
		count++
		if last < lastMap[char] {
			last = lastMap[char]
		}

		if i == last {
			res = append(res, count)
			count = 0
		}
	}
	return res
}
