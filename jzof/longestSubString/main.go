package main

import (
	"fmt"
)

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
// 示例 4:
//
//
//输入: s = ""
//输出: 0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
//
// Related Topics 哈希表 字符串 滑动窗口
// 👍 6432 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	s_length := len(s)
	if s_length == 0 || s_length == 1 {
		return s_length
	}

	s_hash := make(map[string]int)
	left := 0
	right := 0
	l := 1
	finlong := 0
	for k, v := range s {
		right = k
		s_key := string(v)
		//fmt.Printf("now key = %v, left = %v, right =%v, hash = %v, l = %v, finallong = %v \n", s_key, left, right, s_hash, l, finlong)
		fmt.Printf("now key = %v, left = %v, right =%v, hash = %v, l = %v, \n", s_key, left, right, s_hash, l)
		if _, ok := s_hash[s_key]; ok && s_hash[s_key] >= left {
			left = s_hash[s_key] + 1
		} else {
			fmt.Printf("next...left = %v , right =%v \n", left, right)
		}
		s_hash[s_key] = k
		l = right - left + 1
		if finlong <= l {
			finlong = l
		}
	}
	return finlong
}

func main() {
	res := lengthOfLongestSubstring("aadaebcabbaa")
	fmt.Println("res", res)
}

//leetcode submit region end(Prohibit modification and deletion)
