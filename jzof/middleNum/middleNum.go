package main

import (
	"fmt"
)

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
// 示例 3：
//
//
//输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
//
//
// 示例 4：
//
//
//输入：nums1 = [], nums2 = [1]
//输出：1.00000
//
//
// 示例 5：
//
//
//输入：nums1 = [2], nums2 = []
//输出：2.00000
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
//
// Related Topics 数组 二分查找 分治 👍 4723 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	var i, j int
	var ans float64
	sumNums := make([]int, m+n)
	for k := 0; k < len(sumNums); k++ {
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				sumNums[k] = nums1[i]
				i++
			} else {
				sumNums[k] = nums2[j]
				j++
			}
		} else if i < m {
			sumNums[k] = nums1[i]
			i++
		} else {
			sumNums[k] = nums2[j]
			j++
		}
	}
	lenSum := len(sumNums)
	if lenSum%2 == 0 {
		ans = float64(sumNums[lenSum/2]+sumNums[lenSum/2-1]) / 2
	} else {
		ans = float64(sumNums[lenSum/2])
	}
	fmt.Println(sumNums)
	return ans
}

func main() {
	num1 := []int{0, 0}
	num2 := []int{0, 0}
	ans := findMedianSortedArrays(num1, num2)
	fmt.Println(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
