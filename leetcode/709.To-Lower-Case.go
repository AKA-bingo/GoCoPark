package leetcode

/*
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

Example 1:
Input: "Hello"
Output: "hello"

Example 2:
Input: "here"
Output: "here"

Example 3:
Input: "LOVELY"
Output: "lovely"
*/

func ToLowerCase(str string) string {
	resStr := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			resStr[i] = str[i] + 'a' - 'A'
		} else {
			resStr[i] = str[i]
		}
	}
	return string(resStr)
}
