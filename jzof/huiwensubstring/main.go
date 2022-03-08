package main

import (
	"fmt"
)

//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
// 示例 3：
//
//
//输入：s = "a"
//输出："a"
//
//
// 示例 4：
//
//
//输入：s = "ac"
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成
//
// Related Topics 字符串 动态规划 👍 4435 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	ans := ""
	l, r := 0, 0
	for k, v := range s {
		fmt.Printf("i=%v,j=%v\n", k, string(v))
		fmt.Println(string(s[k]))
		if s[k] == s[k+1] {
			for n := 1; k-n < 0 || n+k+1 > len(s)-1; n++ {
				if s[k-n] == s[k+n+1] {
				}
			}
		}
	}
	return ""
}

func main() {
	s := "abccb"
	fmt.Println(longestPalindrome(s))
}

//leetcode submit region end(Prohibit modification and deletion)
