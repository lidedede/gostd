package main

import (
	"fmt"
	"math"
)

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
// Related Topics 递归 链表 数学
// 👍 7070 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(i int) *ListNode {
	head := new(ListNode)
	head.Val = i
	fmt.Println("new val = ", head.Val)
	return head

}
func (head *ListNode) Append(i int) {
	curNode := new(ListNode)
	curNode.Val = i
	if head.Next == nil {
		head.Next = curNode
		curNode.Val = i
	} else {
		nextPoint := head.Next
		for nextPoint.Next != nil {
			nextPoint = nextPoint.Next
		}
		nextPoint.Next = curNode
	}
}

// append 从尾部插入节点
func appendVal(head *ListNode, i int) {
	curNode := new(ListNode)
	curNode.Val = i
	if head.Next == nil {
		head.Next = curNode
		curNode.Val = i
	} else {
		nextPoint := head.Next
		for nextPoint.Next != nil {
			nextPoint = nextPoint.Next
		}
		nextPoint.Next = curNode
	}
}

//遍历链表
func (head *ListNode) Traverse() {
	fmt.Println("--------start----------")
	point := head.Next
	for nil != point {
		fmt.Printf("----------------------one'value=%v, adress=%v \n", point.Val, point)
		point = point.Next
	}
	fmt.Println("--------done----------")
}

func getNum(l *ListNode) (number int) {
	s := []int{}
	s = append(s, l.Val)
	point := l.Next
	for point != nil {
		s = append(s, point.Val)
		point = point.Next
	}
	for i := 0; i < len(s); i++ {
		number += s[i] * int(math.Pow10(i))
	}
	return number
}

func getLen(i int) int {
	n := 0
	for ; i > 0; n++ {
		i /= 10
	}
	return n
}

func getSumSlice(num, len int) (uniqSlice []int) {
	var sumSlice []int
	for i := len - 1; i >= 0; i-- {
		val := num / int(math.Pow10(i))
		num = num - (val * int(math.Pow10(i)))
		fmt.Println("val", i, "=", val)
		sumSlice = append(sumSlice, val)
	}
	for i := len - 1; i >= 0; i-- {
		uniqSlice = append(uniqSlice, sumSlice[i])
	}
	fmt.Println("sumSlice ==", sumSlice)
	fmt.Println("uniqSlice ==", uniqSlice)
	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Num := getNum(l1)
	l2Num := getNum(l2)
	sum := l1Num + l2Num
	sumLen := getLen(sum)

	res := ListNode{}
	fmt.Println("长度是 ", sumLen)
	uniqSumSlice := getSumSlice(sum, sumLen)
	fmt.Println("uniqSumSlice ==", uniqSumSlice)
	for k, v := range uniqSumSlice {
		if k == 0 {
			res.Val = v
		} else {
			appendVal(&res, v)
		}
	}
	return &res
}
func main() {
	l1 := New(1)
	l1.Append(0)
	for i :=0;i <15;i++{
		l1.Append(0)
	}
	l1.Append(1)
	l1.Traverse()

	l2 := New(5)
	l2.Append(6)
	l2.Append(4)
	l2.Traverse()

	res := addTwoNumbers(l1, l2)
	fmt.Println("结果等于", res.Val, res)
	res.Traverse()

	var f float64
	f = 1000000000000000000000000000000000001
	fmt.Println(f)
}

//leetcode submit region end(Prohibit modification and deletion)
