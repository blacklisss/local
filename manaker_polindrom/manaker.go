package main

import "fmt"

func longestPalindromicSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := max(len1, len2)

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right

	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func main() {
	s := "babad"
	fmt.Println(longestPalindromicSubstring(s)) // Expected: "bab" or "aba"

	s2 := "bb"
	fmt.Println(longestPalindromicSubstring(s2)) //

}
